package carproblemzoo

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
	PostDefect(d *DB.Defect) (int32, error)

	GetBrandByName(string) (int32, error)
	PostBrand(brand string) (int32, error)

	GetModelID(brand_id int32, model_name string) (int32, error)
	PostModel(model *DB.ModelData) (int32, error)

	GetGenerationByStartYear(model_id int32, start int32) (int32, error)

	PostVersion(*DB.VersionData) (int32, error)

	GetCountry(country string) (int32, error)
	SaveCountry(country string) (int32, error)

	GetDefectCategory(category string) (int32, error)
	PostDefectCategory(category string) (int32, error)
}

type IReq interface {
	Get(url string) (*goquery.Document, error)
}

const url = "https://www.carproblemzoo.com/"

func Parse(db IDB, getReq func() IReq) {
	req := getReq()
	document, err := req.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	country_id, err := db.GetCountry("🇺🇸")
	if err != nil {
		country_id, err = db.SaveCountry("🇺🇸")
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	brand_urls := getLinks(document.Selection.Find("body > div.container > div.row > div.col-md-8 > div:nth-child(6) > div.panel-body"))
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
		brand_doc, err := req.Get(url + brand_url)
		if err != nil {
			log.Println(err)
			continue
		}
		brand_name := strings.ToLower(
			strings.Split(
				brand_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > h1").Text(),
				" - ")[0],
		)
		wg.Add(1)
		parseBrand(db, getReq(), brand_name, brand_doc, country_id, &done)
	}
	wg.Wait()
}

func parseBrand(db IDB, req IReq, brand_name string, brand_doc *goquery.Document, country_id int32, done *func(string)) {
	brand_id, err := db.GetBrandByName(brand_name)
	if err != nil {
		brand_id, err = db.PostBrand(brand_name)
		if err != nil {
			log.Println(err)
			(*done)(brand_name)
			return
		}
	}
	model_links := getLinks(brand_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > div:nth-child(6) > div.panel-body > table"))
	for _, model_url := range model_links {
		model_url = url + model_url
		model_doc, err := req.Get(model_url)
		if err != nil {
			log.Println(err)
			continue
		}

		model_name := strings.ToLower(
			strings.Split(
				model_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > h1").Text(),
				" - ")[0],
		)
		model_name = strings.TrimSpace(strings.ReplaceAll(model_name, brand_name, ""))

		model_id, err := db.GetModelID(brand_id, model_name)
		if err != nil {
			model_id, err = db.PostModel(&DB.ModelData{Name: model_name, BrandID: brand_id})
			if err != nil {
				log.Println(err)
				continue
			}
		}

		model_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > div:nth-child(15) > div.panel-body > table").Find("tr").Each(
			func(i int, s *goquery.Selection) {
				if i == 0 {
					return
				}
				year_url, ok := s.Find("a").Attr("href")
				if !ok {
					return
				}
				model_year := parseDigit(s.Find("a").Text())
				sales := parseDigit(s.Find("td:nth-child(3)").Text())
				println(model_year, model_id, sales)
				gen_id, _ := db.GetGenerationByStartYear(model_id, int32(model_year))

				year_url = model_url + year_url

				year_doc, err := req.Get(year_url)
				if err != nil {
					log.Println(err)
					return
				}

				for _, major_cat_url := range getLinks(year_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > div:nth-child(6) > div.panel-body")) {
					major_cat_url = year_url + major_cat_url
					major_cat_doc, err := req.Get(major_cat_url)
					if err != nil {
						log.Println(err)
						continue
					}

					for _, minor_cat_url := range getLinks(major_cat_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > div.panel.panel-info > div.panel-body > table")) {
						minor_cat_url = strings.ReplaceAll(year_url, "index.php#ppmy", "") + minor_cat_url
						minor_cat_doc, err := req.Get(minor_cat_url)
						if err != nil {
							log.Println(err)
							continue
						}

						major_cat := minor_cat_doc.Find("body > div.container > div.row > div.col-md-8 > nav > ol > li.breadcrumb-item").Last().Text()
						major_cat_id, err := db.GetDefectCategory(major_cat)
						if err != nil {
							major_cat_id, err = db.PostDefectCategory(major_cat)
							if err != nil {
								log.Println(err)
								continue
							}
						}
						minor_cat_id := major_cat_id

						parse_defect := func(i int, s *goquery.Selection) {
							date := s.Find("div.pull-right.faildate-float").Text()
							year := parseYear(date)
							text := clean(s.Find("p.ptext_list").Text())
							var age int
							if year != 0 {
								age = year - model_year
							}
							_, err = db.PostDefect(&DB.Defect{
								BrandID:         brand_id,
								ModelID:         model_id,
								GenID:           gen_id,
								MajorCategoryID: major_cat_id,
								MinorCategoryID: minor_cat_id,
								CategoryID:      minor_cat_id,
								Country_ID:      country_id,
								Desc:            text,
								Age:             age,
							})
							if err != nil {
								log.Println(err)
							}
						}

						cat_urls := getLinks(minor_cat_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > div.panel.panel-info > div.panel-body > table"))
						for _, cat_url := range cat_urls {
							cat_url = year_url + "/" + cat_url
							cat_doc, err := req.Get(cat_url)
							if err != nil {
								log.Println(err)
								continue
							}

							minor_cat := minor_cat_doc.Find("body > div.container > div.row > div.col-md-8 > nav > ol > li.breadcrumb-item").Last().Text()
							minor_cat_id, err = db.GetDefectCategory(minor_cat)
							if err != nil {
								minor_cat_id, err = db.PostDefectCategory(minor_cat)
								if err != nil {
									log.Println(err)
									continue
								}
							}

							cat_doc.Selection.Find("#div_pslist > div.problem-item").Each(parse_defect)
						}
						if len(cat_urls) == 0 {
							minor_cat_doc.Find("#div_pslist > div.problem-item").Each(parse_defect)
						}
					}
				}
			})
	}
	(*done)(brand_name)
}

func clean(s string) string {
	space := regexp.MustCompile(`[\s|\n]+`)
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

func parseDigit(s string) int {
	str := regexp.MustCompile(`\d+`).FindString(strings.ReplaceAll(s, ",", ""))
	d, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return d
}

func parseYear(s string) int {
	str := regexp.MustCompile(`\d{4}`).FindString(s)
	d, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return d
}

// TODO
// sales
