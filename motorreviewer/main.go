package motorreviewer

import (
	"log"
	DB "main/db"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type IDB interface {
	GetBrand(string) (int32, error)
	GetLastBrands() ([]string, error)
	SaveBrand(string) (int32, error)

	GetEngine(name string) (int32, error)
	SaveEngine(*DB.EngineData) (int32, error)
}

type IReq interface {
	Get(url string) (*goquery.Document, error)
	GetImg(url string) (string, error)
}

const url = "https://www.motorreviewer.com/"

func Parse(db IDB, getReq func() IReq) {
	req := getReq()
	document, err := req.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	var wg sync.WaitGroup
	done := func(brand_name string) {
		println("Done " + brand_name)
		wg.Done()
	}
	for _, brand_url := range getLinks(document.Selection.Find("body > div.content-main > div > div.content-big > div.block > div.content-left > div.en-makes-list")) {
		wg.Add(1)
		go parseBrand(db, getReq(), url+brand_url, &done)
	}
	wg.Wait()
}

func parseBrand(db IDB, req IReq, brand_url string, done *func(string)) {
	engines_doc, err := req.Get(brand_url)
	if err != nil {
		log.Fatal(err)
		(*done)(brand_url)
		return
	}
	for _, engine_url := range getLinks(engines_doc.Selection.Find("body > div.content-main div.engines-box")) {
		engine_url = url + engine_url
		engine_doc, err := req.Get(engine_url)
		if err != nil {
			log.Println(err)
			continue
		}
		path := strings.Split(engine_doc.Selection.Find("body > div.content-main > div > div > div.way").Text(), " ")
		if len(path) < 5 {
			continue
		}
		engine_name := strings.ReplaceAll(strings.Join(path[4:], ""), "Engine", "")

		img_url, _ := engine_doc.Selection.Find("body > div.content-main > div > div > div:nth-child(2) > div.content-left > div:nth-child(3) > div.eng-right-block > img").Attr("src")
		img, _ := req.GetImg(url + img_url)

		specs := engine_doc.Selection.Find("#Specs").Parent()
		fuel_type := strings.ToLower(specs.Find("div:nth-child(5) > div.spec-table-value").Text())
		if strings.Contains(fuel_type, "gasoline") {
			fuel_type = "petrol"
		}
		cylinders, _ := strconv.Atoi(regexp.MustCompile(`[0-9]{1,2}`).FindString(specs.Find("div:nth-child(8) > div.spec-table-value").Text()))
		valves_per_cylinder, _ := strconv.Atoi(regexp.MustCompile(`[0-9]{1}`).FindString(specs.Find("div:nth-child(9) > div.spec-table-value").Text()))
		displacement_str := regexp.MustCompile(`[\d|,]+`).FindString(specs.Find("div:nth-child(13) > div.spec-table-value").Text())
		displacement_str = strings.ReplaceAll(displacement_str, ",", "")
		displacement, _ := strconv.Atoi(displacement_str)
		power := specs.Find("div:nth-child(16) > div.spec-table-value").Text()
		power = regexp.MustCompile(`[0-9]{2,3}`).FindString(power)
		power_hp, _ := strconv.Atoi(power)
		torque_str := specs.Find("div:nth-child(17) > div.spec-table-value").Text()
		torque_str = regexp.MustCompile(`[0-9]{2,3}`).FindString(torque_str)
		torque, _ := strconv.Atoi(torque_str)

		engine_id, err := db.SaveEngine(&DB.EngineData{
			Name:         engine_name,
			Displacement: displacement,
			Valves:       valves_per_cylinder * cylinders,
			Cylinders:    cylinders,
			Fuel_type:    fuel_type,
			Power_hp:     power_hp,
			Torque:       torque,
			Img:          img,
		})
		if err != nil {
			log.Println(err)
		}
		log.Println(engine_id)
		(*done)(brand_url)
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
