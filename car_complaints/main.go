package car_complaints

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type IDB interface {
	PostBrand(string)
}

type Problem struct {
	brand           string
	model           string
	year            uint
	category        string
	name            string
	count           uint
	cost            uint
	severity_rating uint8
	repair_cost     uint
	average_mileage uint
}

func Parse(db IDB) {
	// const url = "https://www.carcomplaints.com"

	// document, err := getDocument(url)
	document, err := getHTML(main_html) //TODO getDocument
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, brand_url := range getLinks(document.Selection.Find("#container #makes")) {
		fmt.Println("Brand: ", brand_url)
		db.PostBrand(strings.Trim(brand_url, "/"))
		brand_doc, err := getHTML(brand_html) //TODO getDocument
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, model_url := range getLinks(brand_doc.Selection.Find("#make .browseby-content")) {
			fmt.Println("Model: ", model_url)
			model_doc, err := getHTML(model_html) //TODO getDocument
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, age_url := range getLinks(model_doc.Selection.Find("#model .timeline")) {
				fmt.Println("Age: ", age_url)
				age_doc, err := getHTML(problem_age_html) //TODO getDocument
				if err != nil {
					fmt.Println(err)
					continue
				}

				for _, problem_category_url := range getLinks(age_doc.Selection.Find("#container #graph ul")) {
					fmt.Println("Category: ", problem_category_url)
					category_doc, err := getHTML(category_html) //TODO getDocument
					if err != nil {
						fmt.Println(err)
						continue
					}

					for _, problem_url := range getLinks(category_doc.Selection.Find("#category #graph ul")) {
						fmt.Println("Problem: ", problem_url)
						problem_doc, err := getHTML(problem_html) //TODO getDocument
						if err != nil {
							fmt.Println(err)
							continue
						}
						problem, err := parseProblem(problem_doc.Selection.Find("#container"))
						if err == nil {
							fmt.Println("Parsed ", problem.name)
							// save(&problem)
						}
					}
				}
			}
		}
	}
}

func getLinks(selection *goquery.Selection) []string {
	links := make([]string, 0)
	selection.Each(func(_ int, ul *goquery.Selection) {
		ul.Find("a").Each(func(_ int, ahref *goquery.Selection) {
			href, exists := ahref.Attr("href")
			text := ahref.Text()
			if exists && href != "#" && !strings.HasPrefix(text, "NHTSA") {
				links = append(links, href)
			}
		})
	})
	return links
}

func parseProblem(selection *goquery.Selection) (Problem, error) {
	breadcrumb := selection.Find("#breadcrumb span").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})

	name := selection.Find("#vheader h1").First().Text()
	year, _ := strconv.Atoi(breadcrumb[2])

	info := selection.Find("#pinfo dd").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	cost, _ := strconv.Atoi(info[0])
	average_mileage, _ := strconv.Atoi(info[1])
	count, _ := strconv.Atoi(info[2])

	return Problem{
		brand:           breadcrumb[0],
		model:           breadcrumb[1],
		year:            uint(year),
		category:        breadcrumb[3],
		name:            name,
		cost:            uint(cost),
		average_mileage: uint(average_mileage),
		count:           uint(count),
	}, nil
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
