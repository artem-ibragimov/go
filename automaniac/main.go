package automaniac

import (
	"fmt"
	"log"
	DB "main/db"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type IDB interface {
	SaveBrand(string) (int32, error)
	GetBrand(string) (int32, error)

	SaveEngine(*DB.EngineData) (int32, error)
	GetEngine(string) (int32, error)

	SaveTransmission(*DB.TransmissionData) (int32, error)
	GetTransmission(int32, string, int32) (int32, error)

	SaveModel(*DB.ModelData) (int32, error)
	GetModel(int32, string, int32) (int32, error)

	SaveVersion(*DB.VersionData) (int32, error)
	// GetVersion(*DB.VersionData) (int32, error)
}

type IReq interface {
	Get(url string) (*goquery.Document, error)
}

const url = "https://www.automaniac.org"

func Parse(db IDB, req IReq) {
	document, err := req.Get(url + "/specs")
	// document, err := getHTML(main_html) //TODO req.Get
	if err != nil {
		log.Println(err)
		return
	}
	for _, brand_url := range getLinks(document.Selection.Find("#main-wrapper #modeli div.box-wrap")) {
		path := strings.Split(brand_url, "/")
		brand_name := strings.TrimSpace(strings.ReplaceAll(path[len(path)-1], "-", " "))
		brand_id, err := db.GetBrand(brand_name)
		if err != nil {
			brand_id, err = db.SaveBrand(brand_name)
			if err != nil {
				log.Println(err)
				return
			}
		}
		// fmt.Println("Brand: ", brand_name, brand_url)

		brand_doc, err := req.Get(url + brand_url)
		// brand_doc, err := getHTML(brand_html) //TODO req.Get
		if err != nil {
			log.Println(err)
			continue
		}
		parseBrand(db, req, brand_name, brand_id, brand_doc)
	}
}

func parseBrand(db IDB, req IReq, brand_name string, brand_id int32, brand_doc *goquery.Document) {
	model_links := getLinks(brand_doc.Selection.Find("#main-wrapper #modeli-auto-lista"))
	for _, model_url := range model_links {
		// fmt.Println("model_url: ", model_url)
		model_doc, err := req.Get(url + model_url) //TODO req.Get
		// model_doc, err := getHTML(model_html) //TODO req.Get
		if err != nil {
			log.Println(err)
			continue
		}

		for _, version_url := range getLinks(model_doc.Selection.Find("#modeli-auto-detaljnije #model-auto-wrap")) {
			// fmt.Println("version_url: ", version_url)
			version_doc, err := req.Get(url + version_url)
			// version_doc, err := getHTML(version_html) //TODO req.Get
			if err != nil {
				log.Println(err)
				continue
			}

			path := strings.Split(model_doc.Find("#breadcrumb-wrap > div.breadcrumb-nav").Text(), "/")
			model := strings.Split(clean(path[len(path)-1]), " ")
			model_name := strings.ToLower(model[1])
			model_year, _ := strconv.ParseInt(model[0], 10, 32)
			if model_year < 2000 {
				continue
			}
			model_data := &DB.ModelData{
				Name:    model_name,
				BrandID: brand_id,
				Year:    int32(model_year),
			}

			model_id, err := db.GetModel(brand_id, model_name, int32(model_year))
			if err != nil {
				model_id, err = db.SaveModel(model_data)
				if err != nil {
					log.Println(err)
					return
				}
			}

			parseVersion(db, model_id, model_data, version_doc)
		}
		fmt.Println(model_url)
	}
}

func parseVersion(db IDB, model_id int32, model_data *DB.ModelData, version_doc *goquery.Document) {
	path := strings.Split(version_doc.Find("#breadcrumb-wrap > div.breadcrumb-nav").Text(), "/")
	version := clean(path[len(path)-1])

	version_doc.Selection.Find("#predlog-auto > div.podaci-box-wrap > div").Each(
		func(i int, info *goquery.Selection) {

			if !strings.Contains(info.Find("div.podaci-naslov").Text(), "Engine") {
				return
			}
			displacement, _ := strconv.ParseInt(clean(info.Find("div:nth-child(5) > div.d2 > strong").Text()), 10, 32)
			power_hp, _ := strconv.ParseInt(clean(info.Find("div:nth-child(10) > div.d2 > strong").Text()), 10, 32)
			torque, _ := strconv.ParseInt(clean(info.Find("div:nth-child(11) > div.d2 > strong").Text()), 10, 32)
			engine_name := clean(info.Find("div.podaci-box-b > p > a").Text())
			engine_id, err := db.GetEngine(engine_name)
			if err != nil {
				engine_id, err = db.SaveEngine((&DB.EngineData{
					Name:         engine_name,
					Displacement: int32(displacement),
					Config:       clean(info.Find("div:nth-child(6) > div.d2 > strong").Text()),
					Valves:       clean(info.Find("div:nth-child(7) > div.d2 > strong").Text()),
					Aspiration:   clean(info.Find("div:nth-child(8) > div.d2 > strong").Text()),
					Fuel_type:    clean(info.Find("div:nth-child(9) > div.d2 > strong").Text()),
					Power_hp:     int32(power_hp),
					Torque:       int32(torque),
				}))
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

				var gears int = 0
				gear_data := regexp.MustCompile(`(\d+)`).FindAllString(gearbox, 1)
				if len(gear_data) != 0 {
					gears, _ = strconv.Atoi(gear_data[0])
				}
				trans_data := &DB.TransmissionData{
					BrandID:      model_data.BrandID,
					Type:         trans_type,
					Gears:        int32(gears),
					Consumtion:   float32(math.Round(cons*100) / 100),
					Acceleration: float32(math.Round(acc*100) / 100),
				}

				trans_id, err := db.GetTransmission(model_data.BrandID, trans_type, int32(gears))
				if err != nil {
					trans_id, err = db.SaveTransmission(trans_data)
					if err != nil {
						log.Println(err)
						return
					}
				}

				_, err = db.SaveVersion(&DB.VersionData{
					Name:     version,
					ModelID:  model_id,
					BrandID:  model_data.BrandID,
					EngineID: engine_id,
					TransID:  trans_id,
				})

				if err != nil {
					log.Println(err)
				}
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
