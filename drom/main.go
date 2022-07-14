package drom

import (
	"fmt"
	"log"
	DB "main/db"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slices"
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

	GetGenID(model_id int32, name string) (int32, error)
	PostGen(data *DB.GenerationData) (int32, error)
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

	// var wg sync.WaitGroup
	done := func(brand_name string) {
		println("Done " + brand_name)
		// wg.Done()
	}
	parsed_brands_names, err := db.GetLastBrands()
	if err == nil && len(parsed_brands_names) > 1 {
		// не уверены что последний бренд спарсили до конца, поэтому выкидываем его
		parsed_brands_names = parsed_brands_names[1:]
	}
	brand_urls := getLinks(document.Selection.Find(`[data-ftid="component_cars-list"]`))
	total_brands := len(brand_urls)
	parsed_brands_count := 0

	for _, brand_url := range brand_urls {
		// wg.Add(1)
		// go parseBrand(db, getReq(), brand_url, &done)
		path := strings.Split(brand_url, "/")
		brand_name := path[len(path)-2]
		if slices.Contains(parsed_brands_names, brand_name) {
			// !FIXME
			// continue
		}
		parseBrand(db, getReq(), brand_url, &done)

		parsed_brands_count++
		fmt.Println(parsed_brands_count,
			"of", total_brands, " brands was parsed: ",
			math.Round(float64(parsed_brands_count*100/total_brands)), "%")
	}
	// wg.Wait()
}

func parseBrand(db IDB, req IReq, brand_url string, done *func(string)) {
	brand_doc, err := req.Get(brand_url)
	if err != nil {
		log.Fatal(err)
		(*done)(brand_url)
		return
	}

	brand_name := brand_doc.Selection.Find(`div[data-ftid="header_breadcrumb"] div`).Last().Find("span").Text()
	brand_name = regexp.MustCompile(`[\w|\d|-|\s]+`).FindString(strings.ToLower(brand_name))
	brand_name = strings.TrimSpace(brand_name)

	brand_id, err := db.GetBrandByName(brand_name)
	if err != nil {
		brand_id, err = db.PostBrand(brand_name)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}

	for _, model_url := range getLinks(brand_doc.Selection.Find(`[data-ftid="component_cars-list"]`)) {
		model_doc, err := req.Get(model_url)
		if err != nil {
			log.Println(err)
			continue
		}

		model_name := model_doc.Selection.Find(`a[data-ftid="component_brand-model_title"]`).Text()
		model_name = strings.ReplaceAll(strings.ToLower(model_name), brand_name, "")
		model_name = regexp.MustCompile(`[\w|\d|-|\s]+`).FindString(model_name)
		model_name = strings.TrimSpace(model_name)

		model_id, err := db.GetModelID(brand_id, model_name)
		if err != nil {
			model_id, err = db.PostModel(&DB.ModelData{Name: model_name, BrandID: brand_id})
			if err != nil {
				log.Fatalln(err)
				return
			}
		}

		for _, gen_url := range getLinks(model_doc.Selection.Find(`[data-ga-stats-name="generations_outlet_item"]`).First().Parent()) {
			gen_url = model_url + gen_url
			gen_doc, err := req.Get(gen_url)
			if err != nil {
				log.Println(err)
				continue
			}

			header := strings.ToLower(gen_doc.Selection.Find("h1.b-title_type_h1").Text())
			gen_start := parseYear(strings.Split(header, "-")[0])
			gen_end := parseYear(strings.Split(header, "-")[1])
			gen_name := strings.Split(header, "(")[0]
			gen_name = strings.Join(regexp.MustCompile(`[\w|\d|-|\s]+`).FindAllString(gen_name, -1), "")
			gen_name = strings.ReplaceAll(gen_name, fmt.Sprintf("%v", gen_start), "")
			gen_name = strings.ReplaceAll(gen_name, fmt.Sprintf("%v", gen_end), "")
			gen_name = regexp.MustCompile(`[\s]+`).ReplaceAllString(gen_name, " ")
			gen_name = strings.TrimSpace(gen_name)
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
			println(gen_id)
			for _, version_url := range getLinks(gen_doc.Selection.Find("table.b-table")) {
				version_url = BASE_URL + version_url
				version_doc, err := req.Get(version_url)
				if err != nil {
					log.Println(err)
					continue
				}

				var data map[string]string = make(map[string]string)
				version_doc.Find("table.b-table tr.b-table__row").Each(func(i int, s *goquery.Selection) {
					k := s.Find("td").First().Text()
					if k == `Название комплектации` {
						fmt.Println(k)
					}
					v, err := s.Find("td").Last().Html()
					if err != nil {
						return
					}
					data[k] = v
				})

			}
			// engine_name := strings.ReplaceAll(strings.Join(path[4:], ""), "Engine", "")

			// // img_url, _ := model_doc.Selection.Find("body > div.content-main > div > div > div:nth-child(2) > div.content-left > div:nth-child(3) > div.eng-right-block > img").Attr("src")
			// // img, _ := req.GetImg(url + img_url)

			// specs := model_doc.Selection.Find("#Specs").Parent()
			// fuel_type := strings.ToLower(specs.Find("div:nth-child(5) > div.spec-table-value").Text())
			// if strings.Contains(fuel_type, "gasoline") {
			// 	fuel_type = "petrol"
			// }
			// cylinders, _ := strconv.Atoi(regexp.MustCompile(`[0-9]{1,2}`).FindString(specs.Find("div:nth-child(8) > div.spec-table-value").Text()))
			// valves_per_cylinder, _ := strconv.Atoi(regexp.MustCompile(`[0-9]{1}`).FindString(specs.Find("div:nth-child(9) > div.spec-table-value").Text()))
			// displacement_str := regexp.MustCompile(`[\d|,]+`).FindString(specs.Find("div:nth-child(13) > div.spec-table-value").Text())
			// displacement_str = strings.ReplaceAll(displacement_str, ",", "")
			// displacement, _ := strconv.Atoi(displacement_str)
			// power := specs.Find("div:nth-child(16) > div.spec-table-value").Text()
			// power = regexp.MustCompile(`[0-9]{2,3}`).FindString(power)
			// power_hp, _ := strconv.Atoi(power)
			// torque_str := specs.Find("div:nth-child(17) > div.spec-table-value").Text()
			// torque_str = regexp.MustCompile(`[0-9]{2,3}`).FindString(torque_str)
			// torque, _ := strconv.Atoi(torque_str)

			// engine_id, err := db.PostEngine(&DB.EngineData{
			// 	Name:         engine_name,
			// 	Displacement: displacement,
			// 	Valves:       valves_per_cylinder * cylinders,
			// 	Cylinders:    cylinders,
			// 	Fuel_type:    fuel_type,
			// 	Power_hp:     power_hp,
			// 	Torque:       torque,
			// 	Img:          img,
			// })
			// if err != nil {
			// 	log.Println(err)
			// }
			// log.Println(engine_id)
			// (*done)(brand_url)
		}
	}
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
