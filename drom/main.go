package drom

import (
	"fmt"
	"log"
	DB "main/db"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slices"
	"golang.org/x/text/encoding/charmap"
)

type IDB interface {
	GetBrandByName(string) (int32, error)
	GetLastBrands() ([]string, error)
	PostBrand(string) (int32, error)

	GetModelID(brand_id int32, model_name string) (int32, error)
	GetLastModelNamesByBrand(brand_id int32) ([]string, error)
	PostModel(*DB.ModelData) (int32, error)

	GetEngineID(name string) (int32, error)
	PostEngine(*DB.EngineData) (int32, error)

	GetTransID(brand_id int32, name string, gears int32) (int32, error)
	PostTrans(*DB.TransData) (int32, error)

	GetGenID(model_id int32, name string) (int32, error)
	GetLastGenNamesByModel(model_id int32) ([]string, error)
	PostGen(data *DB.GenerationData) (int32, error)

	GetVersionID(name string, gen_id int32) (int32, error)
	PostVersion(*DB.VersionData) (int32, error)
}

type IReq interface {
	Get(url string) (*goquery.Document, error)
	GetImg(url string) (string, error)
}

const BASE_URL = "https://www.drom.ru/"
const CATALOG_URL = "https://www.drom.ru/catalog/"

func Parse(db IDB, getReq func() IReq) {
	req := getReq()
	document, err := req.Get(CATALOG_URL)
	if err != nil {
		log.Fatal(err)
		return
	}

	parsed_brands_names, err := db.GetLastBrands()
	if err == nil && len(parsed_brands_names) > 0 {
		// не уверены что последний бренд спарсили до конца, поэтому выкидываем его
		parsed_brands_names = parsed_brands_names[1:]
	}
	brand_urls := getLinks(document.Find(`[data-ftid="component_cars-list"]`))
	total_brands := len(brand_urls)
	parsed_brands_count := 0

	var wg sync.WaitGroup
	done := func(brand_name string) {
		parsed_brands_count++
		println("Done "+brand_name, parsed_brands_count,
			"of", total_brands, " brands was parsed: ",
			math.Round(float64(parsed_brands_count*100/total_brands)), "%")
		wg.Done()
	}
	for _, brand_url := range brand_urls {
		path := strings.Split(brand_url, "/")
		brand_name := strings.TrimSpace(path[len(path)-2])
		if brand_name == "" {
			continue
		}
		if slices.Contains(parsed_brands_names, brand_name) {
			parsed_brands_count++
			continue
		}
		wg.Add(1)
		parseBrand(db, getReq(), brand_url, brand_name, &done)
	}
	wg.Wait()
}

func parseBrand(db IDB, req IReq, brand_url string, brand_name string, done *func(string)) {
	brand_doc, err := req.Get(brand_url)
	if err != nil {
		log.Fatal(err)
		(*done)(brand_name)
		return
	}

	brand_id, err := db.GetBrandByName(brand_name)
	if err != nil {
		brand_id, err = db.PostBrand(brand_name)
		if err != nil {
			log.Fatal(err)
			(*done)(brand_name)
			return
		}
	}

	parsed_model_names, err := db.GetLastModelNamesByBrand(brand_id)
	if err == nil && len(parsed_model_names) > 1 {
		// не уверены что последнюю модель спарсили до конца, поэтому выкидываем его
		parsed_model_names = parsed_model_names[1:]
	}

	model_urls := getLinks(brand_doc.Find(`div[data-ftid="component_cars-list"]`))
	if len(model_urls) == 0 {
		fmt.Println(brand_doc.Html())
	}
	for _, model_url := range model_urls {
		model_doc, err := req.Get(model_url)
		if err != nil {
			log.Println(err)
			continue
		}

		p := strings.Split(model_url, "/")
		model_name := strings.Split(strings.ReplaceAll(p[len(p)-2], "_", " "), " ")[0]
		model_name = strings.ReplaceAll(model_name, brand_name, "")
		model_name = strings.TrimSpace(model_name)

		if model_name == "" || slices.Contains(parsed_model_names, model_name) {
			continue
		}

		model_id, err := db.GetModelID(brand_id, model_name)
		if err != nil {
			model_id, err = db.PostModel(&DB.ModelData{Name: model_name, BrandID: brand_id})
			if err != nil {
				log.Fatalln(err)
				continue
			}
		}

		parsed_gen_names, err := db.GetLastGenNamesByModel(model_id)
		if err == nil && len(parsed_gen_names) > 0 {
			// не уверены что последний поколение спарсили до конца, поэтому выкидываем его
			parsed_gen_names = parsed_gen_names[1:]
		}

		gen_urls := getLinks(model_doc.Find(`[data-ga-stats-name="generations_outlet_item"]`).First().Parent())
		for _, gen_url := range gen_urls {
			gen_url = model_url + gen_url
			gen_doc, err := req.Get(gen_url)
			if err != nil {
				log.Println(err)
				continue
			}

			header := gen_doc.Find("h1.b-title_type_h1").Text()
			gen_start := parseYear(strings.Split(header, "-")[0])
			if gen_start < 2000 {
				continue
			}
			gen_end := parseYear(strings.Split(header, "-")[1])
			header = strings.ToLower(
				strings.ReplaceAll(decodeWindows1251(header), " - технические характеристики и комплектации", ""),
			)
			gen_name := strings.Split(header, "(")[0]
			gen_name = strings.ReplaceAll(gen_name, "поколение", "generation")
			gen_name = strings.ReplaceAll(gen_name, "рестайлинг", "restyle")
			gen_name = strings.ReplaceAll(gen_name, brand_name, "")
			gen_name = strings.ReplaceAll(gen_name, model_name, "")
			gen_name = strings.Join(regexp.MustCompile(`[\w|\d|-|\s]+`).FindAllString(gen_name, -1), "")
			gen_name = regexp.MustCompile(`[\s]+`).ReplaceAllString(gen_name, " ")
			gen_name = strings.TrimSpace(gen_name)

			if slices.Contains(parsed_gen_names, gen_name) {
				continue
			}

			gen_id, err := db.GetGenID(model_id, gen_name)
			if err != nil {
				gen_img_src, ok := gen_doc.Find(`a[data-drom-gallery="generation_images"] img`).First().Attr("src")
				if !ok {
					continue
				}
				gen_img, err := req.GetImg(gen_img_src)
				if err != nil {
					continue
				}
				gen_id, err = db.PostGen(&DB.GenerationData{
					Name:    gen_name,
					ModelID: model_id,
					Start:   gen_start,
					Finish:  gen_end,
					Img:     gen_img,
				})
				if err != nil {
					continue
				}
			}
			version_urls := getLinks(gen_doc.Find("table.b-table_text-left"))
			for _, version_url := range version_urls {
				if version_url == "" || version_url == "#" || strings.Contains(version_url, "engine") {
					continue
				}
				version_url = BASE_URL + version_url
				version_doc, err := req.Get(version_url)
				if err != nil {
					log.Println(err)
					continue
				}

				version := DB.VersionData{GenID: gen_id}
				engine := DB.EngineData{}
				trans := DB.TransData{BrandID: brand_id}
				version_doc.Find("table.b-table tr.b-table__row").Each(func(_ int, s *goquery.Selection) {
					k := decodeWindows1251(s.Find("td").First().Text())
					v := decodeWindows1251(s.Find("td").Last().Text())
					if k == v || v == "" || k == "" {
						return
					}
					if k == `Название комплектации` {
						version.Name = strings.Join(regexp.MustCompile(`[\w|\d|-|\s]+`).FindAllString(v, -1), "")
					}
					if k == `Тип трансмиссии` {
						name := strings.ReplaceAll(v, "МКПП ", "manual-")
						name = strings.ReplaceAll(name, "РКПП ", "amt-")
						name = strings.ReplaceAll(name, "Вариатор", "cvt")
						name = strings.ReplaceAll(name, "АКПП ", "auto-")
						name = strings.ReplaceAll(name, "Редуктор ", "reduction")
						trans.Name = name
						trans.Gears = parseDigit(v)
					}
					if k == `Время разгона 0-100 км/ч, с` {
						trans.Acceleration = parseFloat(v)
					}
					if k == `Расход топлива в городском цикле, л/100 км` || k == `Расход топлива в смешанном цикле, л/100 км` {
						trans.Consumtion = parseFloat(v)
					}
					if k == `Марка двигателя` {
						engine.Name = v
					}
					if k == `Объем двигателя, куб.см` {
						engine.Displacement = parseDigit(v)
					}
					if k == `Используемое топливо` {
						if strings.Contains(v, "Бензин") {
							engine.Fuel_type = "petrol"
						}
						if strings.Contains(v, "Дизель") {
							engine.Fuel_type = "diesel"
						}
						if strings.Contains(v, "Электричество") {
							engine.Fuel_type = "electric"
						}
					}

					if k == `Максимальный крутящий момент, Н*м (кг*м) при об./мин.` {
						engine.Torque = parseDigit(v)
					}
					if k == `Максимальная мощность, л.с. (кВт) при об./мин.` {
						engine.Power_hp = parseDigit(v)
					}
					if k == `Количество клапанов на цилиндр` {
						engine.Valves = parseDigit(v)
					}
					if k == `Тип двигателя` {
						engine.Cylinders = parseDigit(v)
					}
				})

				if version.Name == "" {
					continue
				}
				engine_id, err := db.GetEngineID(engine.Name)
				if err != nil {
					engine_id, err = db.PostEngine(&engine)
					if err != nil {
						log.Fatalln(err)
						continue
					}
				}
				trans_id, err := db.GetTransID(brand_id, trans.Name, int32(trans.Gears))
				if err != nil {
					trans_id, err = db.PostTrans(&trans)
					if err != nil {
						log.Fatalln(err)
						continue
					}
				}

				version.EngineID = engine_id
				version.TransID = trans_id

				_, err = db.GetVersionID(version.Name, gen_id)
				if err != nil {
					id, err := db.PostVersion(&version)
					if err != nil {
						log.Fatalln(err)
						continue
					}
					fmt.Println("#", id, brand_name, "|", model_name, "|", gen_name, "|", version.Name)
				}
			}
		}
	}
	(*done)(brand_name)
}

func getLinks(selection *goquery.Selection) []string {
	keys := make(map[string]bool, 50)
	list := []string{}
	selection.Find("a").Each(func(_ int, ahref *goquery.Selection) {
		href, _ := ahref.Attr("href")
		if !keys[href] {
			keys[href] = true
			list = append(list, href)
		}
	})
	return list
}

func parseYear(s string) int {
	str := regexp.MustCompile(`\d{4}`).FindString(s)
	d, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return d
}

func decodeWindows1251(ba string) string {
	dec := charmap.Windows1251.NewDecoder()
	out, _ := dec.String(clean(ba))
	return out
}

func clean(s string) string {
	space := regexp.MustCompile(`[\s|\n]+`)
	return strings.TrimSpace(space.ReplaceAllString(s, " "))
}

func parseDigit(s string) int {
	str := regexp.MustCompile(`\d+`).FindString(s)
	d, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return d
}
func parseFloat(s string) float32 {
	str := regexp.MustCompile(`[\d|\.]+`).FindString(strings.ReplaceAll(s, ",", "."))
	d, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}
	return float32(d)
}
