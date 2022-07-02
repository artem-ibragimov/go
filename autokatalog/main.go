package autokatalog

import (
	"errors"
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
)

type IDB interface {
	GetBrandByName(string) (int32, error)
	GetLastBrands() ([]string, error)
	SaveBrand(string) (int32, error)

	GetEngineID(name string) (int32, error)
	SaveEngine(*DB.EngineData) (int32, error)

	GetTransmissionID(brand_id int32, name string, gears int32) (int32, error)
	SaveTransmission(*DB.TransmissionData) (int32, error)

	GetModelID(brand_id int32, model_name string) (int32, error)
	GetLastModelNamesByBrand(brand_id int32) ([]string, error)
	SaveModel(*DB.ModelData) (int32, error)

	GetGenerationID(model_id int32, name string) (int32, error)
	SaveGeneration(data *DB.GenerationData) (int32, error)

	GetVersionID(name string, generation_id int32) (int32, error)
	SaveVersion(*DB.VersionData) (int32, error)
}

type IReq interface {
	Get(url string) (*goquery.Document, error)
	GetImg(url string) (string, error)
}

const url = "https://www.adac.de/rund-ums-fahrzeug/autokatalog/marken-modelle/"

func Parse(db IDB, getReq func() IReq) {
	req := getReq()
	document, err := req.Get(url + "?filter=ONLY_RECENT&sort=NAME_ASC")
	// document, err := getHTML(main_html) //TODO req.Get
	if err != nil {
		log.Println(err)
		return
	}
	state, err := extractState(document)
	if err != nil {
		log.Println(err)
		return
	}
	brand_names := extractTags("slug", state)
	parsed_brands_names, err := db.GetLastBrands()
	if err == nil && len(parsed_brands_names) > 1 {
		// не уверены что последний бренд спарсили до конца, поэтому выкидываем его
		parsed_brands_names = parsed_brands_names[1:]
	}
	total_brands := len(brand_names)
	parsed_brands_count := 0
	var wg sync.WaitGroup
	done := func() {
		wg.Done()
		parsed_brands_count = parsed_brands_count + 1
		fmt.Println(parsed_brands_count,
			"of", total_brands, " brands was parsed: ",
			math.Round(float64(parsed_brands_count*100/total_brands)), "%")
	}
	for _, brand_name := range brand_names {
		if slices.Contains(parsed_brands_names, brand_name) {
			continue
		}
		wg.Add(1)
		go parseBrand(db, getReq(), brand_name, &done)
	}
	wg.Wait()
}

func parseBrand(db IDB, req IReq, brand_name string, done *func()) {
	brand_url := url + brand_name
	// fmt.Println("Brand: ", brand_url)

	brand_doc, err := req.Get(brand_url + "?filter=NONE&sort=NAME_ASC")
	// brand_doc, err := getHTML(brand_html) //TODO req.Get
	if err != nil {
		log.Println(err)
		(*done)()
		return
	}

	brand_id, err := db.GetBrandByName(brand_name)
	if err != nil {
		brand_id, err = db.SaveBrand(brand_name)
		if err != nil {
			log.Println(err)
			(*done)()
			return
		}
	}

	state, err := extractState(brand_doc)
	if err != nil {
		log.Println(err)
		(*done)()
		return
	}
	model_slugs := extractTags("slug", state)
	model_names := extractTags("name", state)
	parsed_model_names, err := db.GetLastModelNamesByBrand(brand_id)
	if err == nil && len(parsed_model_names) > 1 {
		// не уверены что последнюю модель спарсили до конца, поэтому выкидываем его
		parsed_model_names = parsed_model_names[1:]
	}
	var model_slug string
	for i, model_name := range model_names {
		if slices.Contains(parsed_model_names, model_name) {
			continue
		}
		model_slug = model_slugs[i]
		model_url := brand_url + "/" + model_slug
		// fmt.Println("Model: ", model_url)
		model_doc, err := req.Get(model_url + "?filter=NONE&sort=SORTING_ASC")
		// model_doc, err := getHTML(model_html) //TODO req.Get
		if err != nil {
			log.Println(err)
			continue
		}

		model_name = strings.ToLower(model_name)
		if model_name == brand_name {
			continue
		}
		model_id, err := db.GetModelID(brand_id, model_name)
		if err != nil {
			model_id, err = db.SaveModel(&DB.ModelData{Name: model_name, BrandID: brand_id})
			if err != nil {
				log.Println(err)
				continue
			}
		}

		state, err := extractState(model_doc)
		if err != nil {
			log.Println(err)
			continue
		}
		generation_slugs := extractTags("slug", state)
		if len(generation_slugs) == 0 {
			continue
		}
		generation_names := extractTags("name", state)
		gen_starts := extractTags("manufacturedFrom", state)
		gen_ends := extractTags("manufacturedUntil", state)
		gen_img_urls := extractTags(`defaultImageUrl`, state)

		for i := range gen_starts {
			gen_start := parseDigit(strings.TrimSpace(strings.ReplaceAll(gen_starts[i], ",", "")))
			gen_end := parseDigit(strings.TrimSpace(strings.ReplaceAll(gen_ends[i], ",", "")))
			gen_img_url := gen_img_urls[i]
			if gen_start < 2000 {
				continue
			}

			gen_url := model_url + "/" + generation_slugs[i]
			// fmt.Println("Version: ", gen_url)
			gen_doc, err := req.Get(gen_url + "?filter=NONE&sort=SORTING_ASC")
			// gen_doc, err := getHTML(generation_html) //TODO req.Get
			if err != nil {
				log.Println(err)
				continue
			}

			gen_name := strings.ToLower(generation_names[i])
			gen_id, err := db.GetGenerationID(model_id, gen_name)
			if err != nil {
				gen_img, _ := req.GetImg(gen_img_url)
				gen_id, err = db.SaveGeneration(&DB.GenerationData{
					Name:    gen_name,
					ModelID: model_id,
					Start:   gen_start,
					Finish:  gen_end,
					Img:     gen_img,
				})
				if err != nil {
					log.Println(err)
					continue
				}
			}
			state, err := extractState(gen_doc)
			if err != nil {
				log.Println(err)
				continue
			}
			versions_ids := extractTags("id", state)
			// versions_names := extractTag("name", extractState(gen_doc))
			for _, id := range versions_ids {
				if len(id) == 6 {

					version_url := gen_url + "/" + id + "/#technische-daten"
					// fmt.Println("version_url: ", version_url)
					version_doc, err := req.Get(version_url)
					// version_doc, err := getHTML(version_html) //TODO req.Get
					if err != nil {
						log.Println(err)
						continue
					}
					state, err := extractState(version_doc)
					if err != nil {
						log.Println(err)
						continue
					}
					version_name := extractTag("car", state)
					if version_name == "" {
						continue
					}

					engine_name := extractTag(`name":"Motorcode","value`, state)
					if engine_name == "" {
						engine_name = brand_name
					}
					engine_id, err := db.GetEngineID(engine_name)

					if err != nil {
						displacement := parseDigit(extractTag(`name":"Hubraum \(Verbrennungsmotor\)","value`, state))
						cylinders := parseDigit(extractTag(`name":"Anzahl Zylinder \(Verbrennungsmotor\)","value`, state))
						valves := parseDigit(extractTag(`name":"Anzahl Ventile \(Verbrennungsmotor\)","value`, state))
						power_hp := parseDigit(extractTag(`Leistung maximal in PS \(Systemleistung\)","value`, state))
						torque := parseDigit(strings.ReplaceAll(extractTag(`name":"Drehmoment \(Systemleistung\)","value`, state), " Nm", ""))
						fuel_type := strings.ToLower(extractTag(`name":"Kraftstoffart","value`, state))
						if strings.Contains(fuel_type, "benzin") || strings.Contains(fuel_type, "super") {
							fuel_type = "petrol"
						}
						engine := DB.EngineData{
							Name:         engine_name,
							Displacement: displacement,
							Cylinders:    cylinders,
							Valves:       valves,
							Fuel_type:    fuel_type,
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
					trans_name = strings.ReplaceAll(trans_name, "Automatisiertes manual", "amt")
					trans_name = strings.ReplaceAll(trans_name, "Automat. manual (Doppelkupplung)", "amt-2")
					trans_name = strings.ReplaceAll(trans_name, "Reduktionsgetriebe", "reduction")
					gears := parseDigit(extractTag(`name":"Anzahl Gänge","value`, state))

					trans_id, err := db.GetTransmissionID(brand_id, trans_name, int32(gears))
					if err != nil {
						consumtion_s := strings.ReplaceAll(extractTag(`name":"Verbrauch Gesamt \(NEFZ\)","value`, state), " l/100 km", "")
						consumtion, _ := strconv.ParseFloat(strings.ReplaceAll(consumtion_s, ",", "."), 32)
						acceleration_s := strings.ReplaceAll(extractTag(`name":"Beschleunigung 0-100km\/h","value`, state), " s", "")
						acceleration, _ := strconv.ParseFloat(strings.ReplaceAll(acceleration_s, ",", "."), 32)
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

					version_id, err := db.GetVersionID(version_name, gen_id)
					if err != nil {
						version_id, err = db.SaveVersion(&DB.VersionData{
							Name:         version_name,
							GenerationID: gen_id,
							EngineID:     engine_id,
							TransID:      trans_id,
						})
						if err != nil {
							log.Println(err)
							continue
						}
					}

					fmt.Println(version_id)
				}
			}
		}
	}
	(*done)()
}

func getHTML(html string) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(html))
}

func extractTags(tag string, s string) []string {
	slugs := regexp.MustCompile(`"`+tag+`":["]?[\w|\-|\s|\(|\)|\.|\/|:|\,|_]+["]?`).FindAllString(s, -1)
	extract_slug := regexp.MustCompile(`"` + tag + `":["]?([\w|\-|\s|\(|\)|\.|\/|:|\,|_]+)["]?`)
	for i, str := range slugs {
		slugs[i] = extract_slug.ReplaceAllString(str, "$1")
	}
	return slugs
}
func extractTag(tag string, s string) string {
	slugs := regexp.MustCompile(`"` + tag + `":"[\w|\-|\s|\(|\)|\.|\/|:|\,|_]+"`).FindString(s)
	extract_slug := regexp.MustCompile(`"` + tag + `":"([\w|\-|\s|\(|\)|\.|\/|:|\,|_]+)"`)
	return extract_slug.ReplaceAllString(slugs, "$1")
}

func extractState(doc *goquery.Document) (string, error) {
	const script string = "window.__APOLLO_STATE__="
	html := doc.Selection.Find("body > script:nth-child(2)").Text()
	if len(html) == 0 || !strings.Contains(html, script) {
		return "", errors.New("no valid doc:\n" + html)
	}
	return html[strings.Index(html, "window.__APOLLO_STATE__=")+len("window.__APOLLO_STATE__="):], nil
}

func parseDigit(s string) int {
	str := regexp.MustCompile(`\d+`).FindString(s)
	d, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return d
}
