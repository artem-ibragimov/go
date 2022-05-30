package automaniac

import (
	"fmt"
	"log"
	DB "main/db"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type IDB interface {
	SaveBrand(string) (int32, error)
	SaveEngine(*DB.EngineData) (int32, error)
	SaveTransmission(*DB.TransmissionData) (int32, error)
	SaveModel(*DB.ModelData) (int32, error)
}

const url = "https://www.automaniac.org"

func Parse(db IDB) {
	document, err := getDocument(url + "/specs")
	// document, err := getHTML(main_html) //TODO getDocument
	if err != nil {
		log.Fatalln(err)
		return
	}
	for brand_name, brand_url := range getLinks(document.Selection.Find("#main-wrapper #modeli div.box-wrap")) {
		brand_id, err := db.SaveBrand(brand_name)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		// fmt.Println("Brand: ", brand_name, brand_url)

		brand_doc, err := getDocument(url + brand_url)
		// brand_doc, err := getHTML(brand_html) //TODO getDocument
		if err != nil {
			log.Fatalln(err)
			continue
		}
		parseBrand(db, brand_name, brand_id, brand_doc)
	}
}

func parseBrand(db IDB, brand_name string, brand_id int32, brand_doc *goquery.Document) {
	for _, model_url := range getLinks(brand_doc.Selection.Find("#main-wrapper #modeli-auto-lista")) {
		// fmt.Println("model_url: ", model_url)
		model_doc, err := getDocument(url + model_url) //TODO getDocument
		// model_doc, err := getHTML(model_html) //TODO getDocument
		if err != nil {
			log.Fatalln(err)
			continue
		}

		for _, version_url := range getLinks(model_doc.Selection.Find("#modeli-auto-detaljnije #model-auto-wrap")) {
			// fmt.Println("version_url: ", version_url)
			version_doc, err := getDocument(url + version_url)
			// version_doc, err := getHTML(version_html) //TODO getDocument
			if err != nil {
				log.Fatalln(err)
				continue
			}
			parseVersion(db, brand_name, brand_id, version_doc)
		}
	}
}

func parseVersion(db IDB, brand_name string, brand_id int32, version_doc *goquery.Document) {
	path := strings.Split(version_doc.Find("#breadcrumb-wrap > div.breadcrumb-nav").Text(), "/")
	version := clean(path[len(path)-1])
	model := strings.Split(clean(path[len(path)-2]), " ")
	model_name := model[1]
	model_year, _ := strconv.ParseInt(model[0], 10, 32)

	version_doc.Selection.Find("#predlog-auto > div.podaci-box-wrap > div").Each(
		func(i int, info *goquery.Selection) {

			if !strings.Contains(info.Find("div.podaci-naslov").Text(), "Engine") {
				return
			}
			displacement, _ := strconv.ParseInt(clean(info.Find("div:nth-child(5) > div.d2 > strong").Text()), 10, 32)
			aspiration, _ := strconv.ParseFloat(clean(info.Find("div:nth-child(8) > div.d2 > strong").Text()), 32)
			power_hp, _ := strconv.ParseInt(clean(info.Find("div:nth-child(10) > div.d2 > strong").Text()), 10, 32)
			torque, _ := strconv.ParseInt(clean(info.Find("div:nth-child(11) > div.d2 > strong").Text()), 10, 32)
			engine_id, err := db.SaveEngine((&DB.EngineData{
				Name:         clean(info.Find("div.podaci-box-b > p > a").Text()),
				Displacement: int32(displacement),
				Config:       clean(info.Find("div:nth-child(6) > div.d2 > strong").Text()),
				Valves:       clean(info.Find("div:nth-child(7) > div.d2 > strong").Text()),
				Aspiration:   float32(aspiration),
				Fuel_type:    clean(info.Find("div:nth-child(9) > div.d2 > strong").Text()),
				Power_hp:     int32(power_hp),
				Torque:       int32(torque),
			}))

			if err != nil {
				log.Fatalln(err)
			}

			version_doc.Selection.Find("#predlog-auto > div.podaci-box-wrap > div").Each(func(_ int, info *goquery.Selection) {
				if !strings.Contains(info.Find("div.podaci-naslov").Text(), "gearbox performance") {
					return
				}
				cons, _ := strconv.ParseFloat(clean(info.Find("div.podaci-box-c > div > span").Text()), 64)
				acc, _ := strconv.ParseFloat(clean(info.Find("div:nth-child(6) > div.d2 > strong").Text()), 64)
				trans_id, err := db.SaveTransmission((&DB.TransmissionData{
					BrandID:      brand_id,
					Desc:         clean(info.Find("div.podaci-box-b").Text()),
					Consumtion:   float32(math.Round(cons*100) / 100),
					Acceleration: float32(math.Round(acc*100) / 100),
				}))

				if err != nil {
					log.Fatalln(err)
				}

				model_id, err := db.SaveModel(&DB.ModelData{
					Name:     model_name,
					Year:     int32(model_year),
					Version:  version,
					BrandID:  brand_id,
					EngineID: engine_id,
					TransID:  trans_id,
				})

				if err != nil {
					log.Fatalln(err)
				}
				if model_id != 0 {
					fmt.Println(brand_name, model_name, version, model_year)
				}
			})
		})
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
func getUA() string {
	uas := []string{
		"Mozilla/5.0 (Linux; Android 7.0; SM-G930V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_4 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B350 Safari/8536.25",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:57.0) Gecko/20100101 Firefox/57.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36",
		"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 5 Build/LMY48B; wv) AppleWebKit/537.36 (KHTML, like Gecko)  Version/4.0 Chrome/43.0.2357.65 Mobile Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/85 Version/11.1.1 Safari/605.1.15",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/603.1.23 (KHTML, like Gecko) Version/10.0 Mobile/14E5239e Safari/602.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X)  AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1",
	}
	return uas[rand.Intn(len(uas))]
}

func getDocument(url string) (*goquery.Document, error) {
	tr := &http.Transport{

		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr, Timeout: 60 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", getUA())

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// body, _ := io.ReadAll(response.Body)
	// fmt.Println(string(body))
	defer response.Body.Close()
	return goquery.NewDocumentFromReader(response.Body)
}

// func getHTML(html string) (*goquery.Document, error) {
// 	return goquery.NewDocumentFromReader(strings.NewReader(html))
// }
