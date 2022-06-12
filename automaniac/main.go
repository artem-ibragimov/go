package automaniac

import (
	"log"
	DB "main/db"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type IDB interface {
	SaveBrand(string) (int32, error)
	GetBrand(string) (int32, error)

	GetEngine(string) (int32, error)
	GetEngineByParams(displacement int, valves int, power_hp int, torque int) (int32, error)
	SaveEngine(*DB.EngineData) (int32, error)

	SaveTransmission(*DB.TransmissionData) (int32, error)
	GetTransmission(int32, string, int32) (int32, error)

	SaveModel(*DB.ModelData) (int32, error)
	GetModel(brand_id int32, model_name string) (int32, error)

	GetGenerationByStartYear(model_id int32, start int32) (int32, error)
	GetGeneration(model_id int32, name string) (int32, error)
	SaveGeneration(data *DB.GenerationData) (int32, error)

	SaveVersion(*DB.VersionData) (int32, error)
	// GetVersion(*DB.VersionData) (int32, error)
}

type IReq interface {
	Get(url string) (*goquery.Document, error)
	GetImg(url string) (string, error)
}

const url = "https://www.automaniac.org"

func Parse(db IDB, getReq func() IReq) {
	req := getReq()
	document, err := req.Get(url + "/specs")
	// document, err := getHTML(main_html) //TODO req.Get
	if err != nil {
		log.Println(err)
		return
	}
	var wg sync.WaitGroup
	done := func() {
		wg.Done()
	}
	for _, brand_url := range getLinks(document.Selection.Find("#main-wrapper #modeli div.box-wrap")) {
		wg.Add(1)
		go parseBrandURL(db, getReq(), brand_url, &done)
	}
	wg.Wait()
}

func parseBrandURL(db IDB, req IReq, brand_url string, done *func()) {
	path := strings.Split(brand_url, "/")
	brand_name := strings.TrimSpace(strings.ReplaceAll(path[len(path)-1], "-", " "))
	brand_id, err := db.GetBrand(brand_name)
	if err != nil {
		brand_id, err = db.SaveBrand(brand_name)
		if err != nil {
			log.Println(err)
			(*done)()
			return
		}
	}
	// fmt.Println("Brand: ", brand_name, brand_url)

	brand_doc, err := req.Get(url + brand_url + "/year")
	// brand_doc, err := getHTML(brand_html) //TODO req.Get
	if err != nil {
		log.Println(err)
		(*done)()
		return
	}
	model_links := getLinks(brand_doc.Selection.Find("#main-wrapper #modeli-auto-lista"))
	for _, model_url := range model_links {
		// fmt.Println("model_url: ", model_url)
		model_doc, err := req.Get(url + model_url) //TODO req.Get
		// model_doc, err := getHTML(model_html) //TODO req.Get
		if err != nil {
			log.Println(err)
			continue
		}

		path := strings.Split(model_doc.Find("#breadcrumb-wrap > div.breadcrumb-nav").Text(), "/")
		model := strings.Split(clean(path[len(path)-1]), " ")
		model_name := strings.ToLower(model[1])

		model_id, err := db.GetModel(brand_id, model_name)
		if err != nil {
			model_data := &DB.ModelData{
				Name:    model_name,
				BrandID: brand_id,
			}
			model_id, err = db.SaveModel(model_data)
			if err != nil {
				log.Println(err)
				continue
			}
		}

		header := model_doc.Find("#modeli-auto-detaljnije > div:nth-child(1) > div.predlog-auto-naslov").Text()
		if header == "" {
			continue
		}
		years := regexp.MustCompile(`\d{4}`).FindAllString(header, -1)
		gen_star, _ := strconv.Atoi(years[0])
		if gen_star < 2000 {
			continue
		}
		subheader := strings.ToLower(strings.Split(model_doc.Find("#modeli-auto-detaljnije > div:nth-child(1) > div.predlog-auto-naslov > div").Text(), ",")[0])
		subheader = strings.ReplaceAll(subheader, brand_name, "")
		subheader = strings.ReplaceAll(subheader, model_name, "")
		gen_name := strings.TrimSpace(subheader)
		if gen_name == "" || strings.Contains(gen_name, "segment") {
			gen_name = years[0]
		}
		var gen_end int = 0
		if len(years) > 1 {
			gen_end, _ = strconv.Atoi(years[1])
		}
		gen_id, err := db.GetGeneration(model_id, gen_name)
		if err != nil {
			gen_id, err = db.GetGenerationByStartYear(model_id, int32(gen_star))
			if err != nil {
				img, _ := req.GetImg(url + model_doc.Find("#main-model-image").AttrOr("src", ""))
				gen_id, err = db.SaveGeneration(&DB.GenerationData{
					Name:    gen_name,
					ModelID: model_id,
					Start:   gen_star,
					Finish:  gen_end,
					Img:     img,
				})
				if err != nil {
					log.Println(err)
					continue
				}
			}
		}

		for _, version_url := range getLinks(model_doc.Selection.Find("#modeli-auto-detaljnije #model-auto-wrap")) {
			// fmt.Println("version_url: ", version_url)
			version_doc, err := req.Get(url + version_url)
			// version_doc, err := getHTML(version_html) //TODO req.Get
			if err != nil {
				log.Println(err)
				continue
			}

			parseVersion(db, gen_id, brand_id, version_doc)
		}
	}
	(*done)()
}

func parseVersion(db IDB, gen_id int32, brand_id int32, version_doc *goquery.Document) {
	path := strings.Split(version_doc.Find("#breadcrumb-wrap > div.breadcrumb-nav").Text(), "/")
	version := clean(path[len(path)-1])

	version_doc.Selection.Find("#predlog-auto > div.podaci-box-wrap > div").Each(
		func(i int, info *goquery.Selection) {

			if !strings.Contains(info.Find("div.podaci-naslov").Text(), "Engine") {
				return
			}
			displacement, _ := strconv.Atoi(clean(info.Find("div:nth-child(5) > div.d2 > strong").Text()))
			valves, _ := strconv.Atoi(clean(info.Find("div:nth-child(7) > div.d2 > strong").Text()))
			power_hp, _ := strconv.Atoi(clean(info.Find("div:nth-child(10) > div.d2 > strong").Text()))
			torque, _ := strconv.Atoi(clean(info.Find("div:nth-child(11) > div.d2 > strong").Text()))
			engine_name := clean(info.Find("div.podaci-box-b > p > a").Text())
			engine_id, err := db.GetEngine(engine_name)
			if err != nil {
				engine_id, err = db.GetEngineByParams(displacement, valves, power_hp, torque)
				if err != nil {
					engine_id, err = db.SaveEngine((&DB.EngineData{
						Name:         engine_name,
						Displacement: displacement,
						Valves:       valves,
						Fuel_type:    clean(info.Find("div:nth-child(9) > div.d2 > strong").Text()),
						Power_hp:     power_hp,
						Torque:       torque,
					}))
				}
			}

			if err != nil {
				log.Println(err)
			}

			version_doc.Selection.Find("#predlog-auto > div.podaci-box-wrap > div").Each(func(_ int, info *goquery.Selection) {
				if !strings.Contains(info.Find("div.podaci-naslov").Text(), "gearbox performance") {
					return
				}
				cons, _ := strconv.ParseFloat(clean(info.Find("div.podaci-box-c > div > span").Text()), 64)
				acc, _ := strconv.ParseFloat(clean(info.Find("div:nth-child(6) > div.d2 > strong").Text()), 64)
				gearbox := clean(info.Find("div.podaci-box-b").Text())

				trans_type := regexp.MustCompile(` \d \w+`).ReplaceAllString(gearbox, "")
				trans_type = strings.ReplaceAll(trans_type, "automatic", "auto")
				trans_type = strings.TrimSpace(strings.ReplaceAll(trans_type, "gearbox", ""))

				var gears int = 0
				gear_data := regexp.MustCompile(`(\d+)`).FindAllString(gearbox, 1)
				if len(gear_data) != 0 {
					gears, _ = strconv.Atoi(gear_data[0])
				}
				trans_data := &DB.TransmissionData{
					BrandID:      brand_id,
					Name:         trans_type,
					Gears:        gears,
					Consumtion:   float32(math.Round(cons*100) / 100),
					Acceleration: float32(math.Round(acc*100) / 100),
				}

				trans_id, err := db.GetTransmission(brand_id, trans_type, int32(gears))
				if err != nil {
					trans_id, err = db.SaveTransmission(trans_data)
					if err != nil {
						log.Println(err)
						return
					}
				}

				version_id, err := db.SaveVersion(&DB.VersionData{
					Name:         version,
					GenerationID: gen_id,
					EngineID:     engine_id,
					TransID:      trans_id,
				})

				if err != nil {
					log.Println(err)
					return
				}
				log.Println(version_id)
			})
		})
}

func clean(s string) string {
	space := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(space.ReplaceAllString(s, " "))
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

// func getHTML(html string) (*goquery.Document, error) {
// 	return goquery.NewDocumentFromReader(strings.NewReader(html))
// }
