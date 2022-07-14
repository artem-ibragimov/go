package carproblemzoo

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
	SaveDefect(d *DB.Defect) (int32, error)

	GetBrandByName(string) (int32, error)
	PostBrand(brand string) (int32, error)

	GetModelID(brand_id int32, model_name string) (int32, error)
	PostModel(model *DB.ModelData) (int32, error)

	PostVersion(*DB.VersionData) (int32, error)
}

type IReq interface {
	Get(url string) (*goquery.Document, error)
}

const url = "https://www.carproblemzoo.com/"

func Parse(db IDB, getReq func() IReq) {
	// document, err := req.Get(url)
	document, err := getHTML(main_html) //TODO req.Get
	if err != nil {
		log.Println(err)
		return
	}
	req := getReq()

	for _, brand_url := range getLinks(document.Selection.Find("body > div.container > div.row > div.col-md-8 > div:nth-child(6) > div.panel-body")) {
		fmt.Println("Brand: ", brand_url)
		// brand_doc, err := req.Get(url + brand_url)
		brand_doc, err := getHTML(brand_html) //TODO req.Get
		if err != nil {
			log.Println(err)
			continue
		}
		brand_name := strings.ToLower(
			strings.Split(
				brand_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > h1").Text(),
				" - ")[0],
		)
		brand_id, err := db.GetBrandByName(brand_name)
		if err != nil {
			brand_id, err = db.PostBrand(brand_name)
			if err != nil {
				log.Println(err)
				return
			}
		}
		parseBrand(db, req, brand_name, brand_id, brand_doc)
	}
}

func parseBrand(db IDB, req IReq, brand_name string, brand_id int32, brand_doc *goquery.Document) {
	model_links := getLinks(brand_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > div:nth-child(6) > div.panel-body > table"))
	for _, model_url := range model_links {
		fmt.Println("model_url: ", model_url)
		// model_doc, err := req.Get(url + model_url)
		model_doc, err := getHTML(model_html) //TODO req.Get
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

		for _, year_url := range getLinks(model_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > div:nth-child(5) > div.panel-body")) {

			println(year_url, model_id)
			// year_doc, err := req.Get(url + year_url)
			year_doc, err := getHTML(year_html) //TODO req.Get
			if err != nil {
				log.Println(err)
				continue
			}

			model_year, _ := strconv.Atoi(strings.ReplaceAll(year_url, "/", ""))
			log.Println(model_year)
			// year, err := strconv.Atoi(
			// 	strings.TrimSpace(
			// 		strings.ReplaceAll(
			// 			strings.ReplaceAll(
			// 				strings.ToLower(
			// 					strings.Split(
			// 						year_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > h1").Text(),
			// 						" - ")[0],
			// 				),
			// 				brand_name, ""),
			// 			model_name, ""),
			// 	),
			// )

			for _, minor_cat_url := range getLinks(year_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > div:nth-child(6) > div.panel-body")) {
				println(minor_cat_url)
				// minor_cat_doc, err := req.Get(url + minor_cat_url)
				minor_cat_doc, err := getHTML(minor_cat_html) //TODO req.Get
				if err != nil {
					log.Println(err)
					continue
				}
				for _, cat_url := range getLinks(minor_cat_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > div.panel.panel-info > div.panel-body > table")) {
					println(cat_url)
					// cat_doc, err := req.Get(url + cat_url)
					cat_doc, err := getHTML(cat_html) //TODO req.Get
					if err != nil {
						log.Println(err)
						continue
					}

					path := removeEmptyStrings(strings.Split(
						strings.ToLower(
							cat_doc.Selection.Find("body > div.container > div.row > div.col-md-8 > nav > ol").Text(),
						),
						"\n"))
					if len(path) < 6 {
						fmt.Println("Path is too short", path)
						continue
					}
					brand_name := path[1]
					model_name := path[2]
					car_year, err := strconv.Atoi(path[3])
					if err != nil {
						println(err)
						continue
					}
					defect_min_cat := path[4]
					defect_cat := path[5]
					println(brand_name, model_name, car_year, defect_min_cat, defect_cat)

					// model_years, err := db.GetModelIDYears(brand_id, model_name)
					if err != nil {
						println(err)
						continue
					}
					// var model_year int
					// for _, y := range model_years {
					// 	if car_year > y {
					// 		model_year = y
					// 	}
					// }
					// println(model_year)
					// println(model_years)
					cat_doc.Selection.Find("#div_pslist > div.problem-item").Each(func(i int, item *goquery.Selection) {
						defect_date := strings.Split(
							strings.ReplaceAll(
								item.Find("div.pull-right.faildate-float").Text(),
								"\n", ""),
							"/")
						defect_year, err := strconv.Atoi(defect_date[len(defect_date)-1])
						if err != nil {
							return
						}
						defect_desc := strings.ReplaceAll(item.Find("p.ptext_list").Text(), "\n", "")
						println(defect_year, defect_desc)
						db.SaveDefect(&DB.Defect{
							BrandID: brand_id,
						})
					})
				}
			}
		}
	}
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

func getHTML(html string) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(html))
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		str = strings.TrimSpace(str)
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
