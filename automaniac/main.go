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
	GetBrandByName(string) (int32, error)

	GetEngineID(string) (int32, error)
	GetEngineByParams(displacement int, valves int, power_hp int, torque int) (int32, error)
	PostEngine(*DB.EngineData) (int32, error)

	PostTrans(*DB.TransData) (int32, error)
	GetTransID(int32, string, int32) (int32, error)

	SaveModel(*DB.ModelData) (int32, error)
	GetModelID(brand_id int32, model_name string) (int32, error)

	GetGenerationByStartYear(model_id int32, start int32) (int32, error)
	GetGenID(model_id int32, name string) (int32, error)
	PostGeneration(data *DB.GenerationData) (int32, error)

	GetVersionID(name string, generation_id int32) (int32, error)
	PostVersion(*DB.VersionData) (int32, error)
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
	brand_id, err := db.GetBrandByName(brand_name)
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
		if len(model) < 2 {
			continue
		}
		model_name := strings.ToLower(model[1])

		model_id, err := db.GetModelID(brand_id, model_name)
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
		gen_id, err := db.GetGenID(model_id, gen_name)
		if err != nil {
			gen_id, err = db.GetGenerationByStartYear(model_id, int32(gen_star))
			if err != nil {
				img, _ := req.GetImg(url + model_doc.Find("#main-model-image").AttrOr("src", ""))
				gen_id, err = db.PostGeneration(&DB.GenerationData{
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
	version_name := clean(path[len(path)-1])

	version_doc.Selection.Find("#predlog-auto > div.podaci-box-wrap > div").Each(
		func(i int, info *goquery.Selection) {

			if !strings.Contains(info.Find("div.podaci-naslov").Text(), "Engine") {
				return
			}
			engine_name := clean(info.Find("div.podaci-box-b > p > a").Text())
			engine_id, err := db.GetEngineID(engine_name)
			if err != nil {
				displacement, _ := strconv.Atoi(clean(info.Find("div:nth-child(5) > div.d2 > strong").Text()))
				cfg := clean(info.Find("div:nth-child(7) > div.d2 > strong").Text())
				params := regexp.MustCompile(`\d+`).FindAllString(cfg, -1)
				var valves, cylinders int
				if len(params) == 2 {
					valves, _ = strconv.Atoi(params[0])
					cylinders, _ = strconv.Atoi(params[1])
				}
				power_hp, _ := strconv.Atoi(clean(info.Find("div:nth-child(10) > div.d2 > strong").Text()))
				torque, _ := strconv.Atoi(clean(info.Find("div:nth-child(11) > div.d2 > strong").Text()))
				engine_id, err = db.GetEngineByParams(displacement, valves, power_hp, torque)
				if err != nil {
					engine_id, err = db.PostEngine((&DB.EngineData{
						Name:         engine_name,
						Displacement: displacement,
						Cylinders:    cylinders,
						Valves:       valves,
						Fuel_type:    strings.ToLower(clean(info.Find("div:nth-child(9) > div.d2 > strong").Text())),
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
				trans_data := &DB.TransData{
					BrandID:      brand_id,
					Name:         trans_type,
					Gears:        gears,
					Consumtion:   float32(math.Round(cons*100) / 100),
					Acceleration: float32(math.Round(acc*100) / 100),
				}

				trans_id, err := db.GetTransID(brand_id, trans_type, int32(gears))
				if err != nil {
					trans_id, err = db.PostTrans(trans_data)
					if err != nil {
						log.Println(err)
						return
					}
				}

				version_id, err := db.GetVersionID(version_name, gen_id)
				if err != nil {
					version_id, err = db.PostVersion(&DB.VersionData{
						Name:     version_name,
						GenID:    gen_id,
						EngineID: engine_id,
						TransID:  trans_id,
					})

					if err != nil {
						log.Println(err)
						return
					}
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
