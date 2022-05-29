package automaniac

import (
	"fmt"
	"log"
	DB "main/db"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type IDB interface {
	SaveBrand(string) (int32, error)
	SaveEngine(*DB.EngineData) (int32, error)
	SaveTransmission(*DB.TransmissionData) (int32, error)
	SaveModel(*DB.ModelData) (int32, error)
}

func Parse(db IDB) {
	// const url = "https://www.automaniac.org/specs"

	// document, err := getDocument(url)
	document, err := getHTML(main_html) //TODO getDocument
	if err != nil {
		fmt.Println(err)
		return
	}
	for brand_name, brand_url := range getLinks(document.Selection.Find("#main-wrapper #modeli div.box-wrap")) {
		brand_id, err := db.SaveBrand(brand_name)
		if err != nil {
			log.Fatal(err)
			continue
		}
		fmt.Println("Brand: ", brand_name, brand_url)

		brand_doc, err := getHTML(brand_html) //TODO getDocument
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, model_url := range getLinks(brand_doc.Selection.Find("#modeli-auto-lista > div:nth-child(4)")) {
			fmt.Println("model_url: ", model_url)
			model_doc, err := getHTML(model_html) //TODO getDocument
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, version_url := range getLinks(model_doc.Selection.Find("#modeli-auto-detaljnije #model-auto-wrap")) {
				fmt.Println("version_url: ", version_url)
				version_doc, err := getHTML(version_html) //TODO getDocument
				if err != nil {
					fmt.Println(err)
					continue
				}

				path := strings.Split(version_doc.Find("#breadcrumb-wrap > div.breadcrumb-nav").Text(), "/")
				version := clean(path[len(path)-1])
				model := strings.Split(clean(path[len(path)-2]), " ")
				model_name := model[1]
				model_year := model[0]

				version_doc.Selection.Find("#predlog-auto > div.podaci-box-wrap > div").Each(
					func(i int, info *goquery.Selection) {

						if !strings.Contains(info.Find("div.podaci-naslov").Text(), "Engine") {
							return
						}

						engine_id, err := db.SaveEngine((&DB.EngineData{
							Name:         clean(info.Find("div.podaci-box-b > p > a").Text()),
							Displacement: clean(info.Find("div:nth-child(5) > div.d2 > strong").Text()),
							Config:       clean(info.Find("div:nth-child(6) > div.d2 > strong").Text()),
							Valves:       clean(info.Find("div:nth-child(7) > div.d2 > strong").Text()),
							Aspiration:   clean(info.Find("div:nth-child(8) > div.d2 > strong").Text()),
							Fuel_type:    clean(info.Find("div:nth-child(9) > div.d2 > strong").Text()),
							Power_hp:     clean(info.Find("div:nth-child(10) > div.d2 > strong").Text()),
							Torque:       clean(info.Find("div:nth-child(11) > div.d2 > strong").Text()),
						}))

						if err != nil {
							log.Fatal(err)
						}
						fmt.Println("Engine saved with id=", engine_id)

						version_doc.Selection.Find("#predlog-auto > div.podaci-box-wrap > div").Each(func(_ int, info *goquery.Selection) {
							if !strings.Contains(info.Find("div.podaci-naslov").Text(), "gearbox performance") {
								return
							}
							trans_id, err := db.SaveTransmission((&DB.TransmissionData{
								Name:         clean(info.Find("div.podaci-box-b").Text()),
								Consumtion:   clean(info.Find("div.podaci-box-c > div > span").Text()),
								Acceleration: clean(info.Find("div:nth-child(6) > div.d2 > strong").Text()),
							}))

							if err != nil {
								log.Fatal(err)
							}
							fmt.Println("Transmission saved with id=", trans_id)

							db.SaveModel(&DB.ModelData{
								Name:     model_name,
								Year:     model_year,
								Version:  version,
								BrandID:  brand_id,
								EngineID: engine_id,
								TransID:  trans_id,
							})
						})
					})
			}
		}
	}
}

func clean(s string) string {
	space := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(space.ReplaceAllString(s, " "))
}

func getLinks(selection *goquery.Selection) map[string]string {
	links := make(map[string]string)
	selection.Each(func(_ int, ul *goquery.Selection) {
		ul.Find("a").Each(func(_ int, ahref *goquery.Selection) {
			href, h_exists := ahref.Attr("href")
			title, t_exists := ahref.Attr("title")
			text := clean(ahref.Text())
			if h_exists && t_exists {
				links[title] = href
				return
			}
			links[text] = href
		})
	})
	return links
}

// func getDocument(url string) (*goquery.Document, error) {
// 	// tr := &http.Transport{

// 	// 	MaxIdleConns:       10,
// 	// 	IdleConnTimeout:    30 * time.Second,
// 	// 	DisableCompression: true,
// 	// }

// 	// client := &http.Client{Transport: tr, Timeout: 30 * time.Second}
// 	// req, err := http.NewRequest("GET", url, nil)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:97.0) Gecko/20100101 Firefox/97.0")

// 	// response, err := client.Do(req)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// body, _ := io.ReadAll(response.Body)
// 	// fmt.Println(string(body))
// 	// defer response.Body.Close()
// 	// return goquery.NewDocumentFromReader(response.Body)
// 	return goquery.NewDocumentFromReader(strings.NewReader(main_html))
// }
func getHTML(html string) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(html))
}
