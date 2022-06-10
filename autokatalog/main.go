package autokatalog

import (
	"fmt"
	"log"
	DB "main/db"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type IDB interface {
	SaveBrand(string) (int32, error)
	GetBrand(string) (int32, error)

	GetEngine(name string) (int32, error)
	SaveEngine(*DB.EngineData) (int32, error)

	GetTransmission(brand_id int32, name string, gears int32) (int32, error)
	SaveTransmission(*DB.TransmissionData) (int32, error)

	GetModel(brand_id int32, model_name string) (int32, error)
	SaveModel(*DB.ModelData) (int32, error)

	GetGeneration(model_id int32, name string) (int32, error)
	SaveGeneration(data *DB.GenerationData) (int32, error)

	GetVersion(name string, generation_id int32) (int32, error)
	SaveVersion(*DB.VersionData) (int32, error)
}

type IReq interface {
	Get(url string) (*goquery.Document, error)
	GetImg(url string) (string, error)
}

const url = "https://www.adac.de/rund-ums-fahrzeug/autokatalog/marken-modelle/"

func Parse(db IDB, req IReq) {
	// document, err := req.Get(url + "?filter=ONLY_RECENT&sort=NAME_ASC")
	document, err := getHTML(main_html) //TODO req.Get
	if err != nil {
		log.Println(err)
		return
	}
	brand_names := extractTags("slug", extractState(document))

	for _, brand_name := range brand_names {

		brand_url := url + brand_name
		fmt.Println("Brand: ", brand_url)

		// brand_doc, err := req.Get(url + brand_url)
		brand_doc, err := getHTML(brand_html) //TODO req.Get
		if err != nil {
			log.Println(err)
			continue
		}

		brand_id, err := db.GetBrand(brand_name)
		if err != nil {
			brand_id, err = db.SaveBrand(brand_name)
			if err != nil {
				log.Println(err)
				return
			}
		}

		model_names := extractTags("slug", extractState(brand_doc))

		for _, model_name := range model_names {

			model_url := brand_url + "/" + model_name
			fmt.Println("Model: ", model_url)
			model_doc, err := getHTML(model_html) //TODO req.Get
			if err != nil {
				log.Println(err)
				continue
			}

			model_id, err := db.GetModel(brand_id, model_name)
			if err != nil {
				model_id, err = db.SaveModel(&DB.ModelData{Name: model_name, BrandID: brand_id})
				if err != nil {
					log.Println(err)
					return
				}
			}

			generation_names := extractTags("slug", extractState(model_doc))
			generation_names = generation_names[:len(generation_names)-2]

			for i, gen_name := range generation_names {

				gen_url := model_url + "/" + gen_name
				fmt.Println("Version: ", gen_url)
				gen_doc, err := getHTML(generation_html) //TODO req.Get
				if err != nil {
					log.Println(err)
					continue
				}

				state := extractState(gen_doc)
				println(strings.Contains(state, "manufacturedFrom"))
				// TODO не достает года
				gen_start := extractTags("manufacturedFrom", state)[i]
				gen_end := extractTags(`manufacturedUntil`, state)[i]
				gen_img_url := extractTag(`caption":"`+gen_name+`","defaultImageUrl`, state)
				gen_img, _ := req.GetImg(gen_img_url)
				gen_id, _ := db.GetGeneration(model_id, gen_name)
				if err != nil {
					gen_id, err = db.SaveGeneration(&DB.GenerationData{
						Name:    gen_name,
						ModelID: model_id,
						Start:   gen_start,
						End:     gen_end,
						Img:     gen_img,
					})
					if err != nil {
						log.Println(err)
						return
					}
				}

				versions_ids := extractTags("id", extractState(gen_doc))
				// versions_names := extractTag("name", extractState(gen_doc))
				for _, id := range versions_ids {
					if len(id) == 6 {

						version_url := gen_url + "/" + id + "/#technische-daten"
						fmt.Println("version_url: ", version_url)
						version_doc, err := getHTML(version_html) //TODO req.Get
						if err != nil {
							log.Println(err)
							continue
						}
						state := extractState(version_doc)
						version_name := extractTag("car", state)

						engine_name := extractTag(`name":"Motorcode","value`, state)
						engine_id, err := db.GetEngine(engine_name)

						if err != nil {
							displacement, _ := strconv.Atoi(extractTag(`name":"Hubraum (Verbrennungsmotor)","value`, state))
							cylinders, _ := strconv.Atoi(extractTag(`name":"Anzahl Zylinder (Verbrennungsmotor)","value`, state))
							valves, _ := strconv.Atoi(extractTag(`name":"Anzahl Ventile (Verbrennungsmotor)","value`, state))
							power_hp, _ := strconv.Atoi(extractTag(`Leistung maximal in PS (Systemleistung)","value`, state))
							torque, _ := strconv.Atoi(strings.ReplaceAll(extractTag(`name":"Drehmoment (Systemleistung)","value`, state), " Nm", ""))

							engine := DB.EngineData{
								Name:         engine_name,
								Displacement: displacement,
								Cylinders:    cylinders,
								Valves:       valves,
								Fuel_type:    strings.ReplaceAll(extractTag(`name":"Kraftstoffart","value`, state), "Super", "Benzin"),
								Power_hp:     power_hp,
								Torque:       torque,
							}
							engine_id, err = db.SaveEngine(&engine)
							if err != nil {
								log.Println(err)
								continue
							}
						}

						trans_name := extractTag(`name":"Getriebeart","value`, state)
						trans_name = strings.ReplaceAll(trans_name, "-Getriebe", "")
						trans_name = strings.ReplaceAll(trans_name, "Automatikgetriebe", "auto")
						trans_name = strings.ReplaceAll(trans_name, "Schaltgetriebe", "manual")
						gears, _ := strconv.Atoi(extractTag(`name":"Anzahl Gänge","value`, state))

						trans_id, err := db.GetTransmission(brand_id, trans_name, int32(gears))
						if err != nil {
							consumtion, _ := strconv.ParseFloat(strings.ReplaceAll(extractTag(`name":"Verbrauch Gesamt (NEFZ)","value`, state), " l/100 km", ""), 32)
							acceleration, _ := strconv.ParseFloat(strings.ReplaceAll(extractTag(`name":"Beschleunigung 0-100km/h","value`, state), " s", ""), 32)
							transmission := &DB.TransmissionData{
								BrandID:      brand_id,
								Name:         trans_name,
								Gears:        gears,
								Consumtion:   float32(consumtion),
								Acceleration: float32(acceleration),
							}
							trans_id, err = db.SaveTransmission(transmission)
							if err != nil {
								log.Println(err)
								continue
							}
						}

						version_id, err := db.GetVersion(version_name, gen_id)
						if err != nil {
							version_id, err = db.SaveVersion(&DB.VersionData{
								Name:         gen_name,
								GenerationID: gen_id,
								EngineID:     engine_id,
								TransID:      trans_id,
							})
							if err != nil {
								log.Println(err)
								return
							}
						}

						fmt.Println(version_id)
					}
				}
			}
		}
	}
}

func getHTML(html string) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(html))
}

func extractTags(tag string, s string) []string {
	slugs := regexp.MustCompile(`"`+tag+`":"[\w|\-|\s|\(|\)|\.|\/]+"`).FindAllString(s, -1)
	extract_slug := regexp.MustCompile(`"` + tag + `":"([\w|\-|\s|\(|\)|\.|\/]+)"`)
	for i, str := range slugs {
		slugs[i] = extract_slug.ReplaceAllString(str, "$1")
	}
	return slugs
}
func extractTag(tag string, s string) string {
	slugs := regexp.MustCompile(`"` + tag + `":"[\w|\-|\s|\(|\)|\.|\/]+"`).FindString(s)
	extract_slug := regexp.MustCompile(`"` + tag + `":"([\w|\-|\s|\(|\)|\.|\/]+)"`)
	return extract_slug.ReplaceAllString(slugs, "$1")
}

func extractState(doc *goquery.Document) string {
	html := doc.Selection.Find("body > script:nth-child(2)").Text()
	return html[strings.Index(html, "window.__APOLLO_STATE__=")+len("window.__APOLLO_STATE__="):]
}
