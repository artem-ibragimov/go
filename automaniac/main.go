package automaniac

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type IDB interface {
	SaveBrand(string)
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
		fmt.Println("Brand: ", brand_name, brand_url)
		// db.SaveBrand(brand_name)
		brand_doc, err := getHTML(brand_html) //TODO getDocument
		if err != nil {
			fmt.Println(err)
			continue
		}

		for model_name, model_url := range getLinks(brand_doc.Selection.Find("#modeli-auto-lista > div:nth-child(4)")) {
			fmt.Println("Model: ", model_name, model_url)
			model_doc, err := getHTML(model_html) //TODO getDocument
			if err != nil {
				fmt.Println(err)
				continue
			}

			for version_name, version_url := range getLinks(model_doc.Selection.Find("#modeli-auto-detaljnije #model-auto-wrap")) {
				fmt.Println("Version: ", version_name, version_url)
				version_doc, err := getHTML(variants_html) //TODO getDocument
				if err != nil {
					fmt.Println(err)
					continue
				}

				for version_name, version_url := range getLinks(version_doc.Selection.Find("#modeli-auto-detaljnije #model-auto-wrap")) {
					fmt.Println("Version: ", version_name, version_url)
					version_doc, err := getHTML(variants_html) //TODO getDocument
					if err != nil {
						fmt.Println(err)
						continue
					}
					println(version_doc)
				}
			}

			// 	for _, problem_category_url := range getLinks(age_doc.Selection.Find("#container #graph ul")) {
			// 		fmt.Println("Category: ", problem_category_url)
			// 		category_doc, err := getHTML(category_html) //TODO getDocument
			// 		if err != nil {
			// 			fmt.Println(err)
			// 			continue
			// 		}

			// 		for _, problem_url := range getLinks(category_doc.Selection.Find("#category #graph ul")) {
			// 			fmt.Println("Problem: ", problem_url)
			// 			problem_doc, err := getHTML(problem_html) //TODO getDocument
			// 			if err != nil {
			// 				fmt.Println(err)
			// 				continue
			// 			}
			// 			problem, err := parseProblem(problem_doc.Selection.Find("#container"))
			// 			if err == nil {
			// 				fmt.Println("Parsed ", problem.name)
			// 				// save(&problem)
			// 			}
			// 		}
			// 	}
			// }
		}
	}
}

func getLinks(selection *goquery.Selection) map[string]string {
	links := make(map[string]string)
	selection.Each(func(_ int, ul *goquery.Selection) {
		ul.Find("a").Each(func(_ int, ahref *goquery.Selection) {
			href, h_exists := ahref.Attr("href")
			title, t_exists := ahref.Attr("title")
			text := strings.TrimSpace(ahref.Text())
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
