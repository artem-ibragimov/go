package nhtsa

import (
	"encoding/csv"
	"io"
	"log"
	DB "main/db"
	"os"
	"strconv"
	"strings"
)

type IDB interface {
	GetDefectCategory(category string) (int32, error)
	PostDefectCategory(category string) (int32, error)
	PostDefect(d *DB.Defect) (int32, error)

	GetBrandByName(brand string) (int32, error)
	PostBrand(brand string) (int32, error)

	GetModelID(brand_id int32, model_name string) (int32, error)
	PostModel(model *DB.ModelData) (int32, error)

	GetCountry(country string) (int32, error)
	SaveCountry(country string) (int32, error)
}

func Parse(db IDB) {
	f, err := os.Open("nhtsa\\nhtsa.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'

	country_id, err := db.GetCountry("🇺🇸")
	if err != nil {
		country_id, err = db.SaveCountry("🇺🇸")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		defect := new(Defect)
		defect.Init(rec)
		if defect.ModelYear < 2100 &&
			defect.ModelYear >= 2000 &&
			defect.Year >= 2000 &&
			defect.Year < 2100 &&
			defect.Category != "" {
			_, err = storeDefect(db, country_id, defect)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func storeDefect(db IDB, country_id int32, defect *Defect) (int32, error) {

	major_category_id, err := db.GetDefectCategory(defect.MajorCategory)
	if err != nil {
		major_category_id, err = db.PostDefectCategory(defect.MajorCategory)
		if err != nil {
			return 0, err
		}
	}

	var minor_category_id = major_category_id

	if defect.MinorCategory != defect.MajorCategory {
		minor_category_id, err = db.GetDefectCategory(defect.MinorCategory)
		if err != nil {
			minor_category_id, err = db.PostDefectCategory(defect.MinorCategory)
			if err != nil {
				return 0, err
			}
		}
	}

	var category_id = minor_category_id

	if defect.Category != defect.MinorCategory {
		category_id, err = db.GetDefectCategory(defect.Category)
		if err != nil {
			category_id, err = db.PostDefectCategory(defect.Category)
			if err != nil {
				return 0, err
			}
		}
	}

	brand_id, err := db.GetBrandByName(defect.BrandName)
	if err != nil {
		brand_id, err = db.PostBrand(defect.BrandName)
		log.Println("Saved", defect.BrandName)
		if err != nil {
			return 0, err
		}
	}

	model_id, err := db.GetModelID(brand_id, defect.ModelName)
	if err != nil {
		model := &DB.ModelData{
			Name:    defect.ModelName,
			BrandID: brand_id,
		}
		model_id, err = db.PostModel(model)
		if err != nil {
			return 0, err
		}
	}

	_, err = db.PostDefect(&DB.Defect{
		BrandID:         brand_id,
		ModelID:         model_id,
		Age:             defect.Year,
		MajorCategoryID: major_category_id,
		MinorCategoryID: minor_category_id,
		CategoryID:      category_id,
		Cost:            "0",
		Rating:          0,
		Mileage:         defect.Miles,
		Freq:            defect.Freq,
		Desc:            defect.Desc,
		Country_ID:      country_id,
	})
	return 0, err
}

type Defect struct {
	BrandName     string
	ModelName     string
	ModelYear     int32
	Year          int
	MajorCategory string
	MinorCategory string
	Category      string
	Miles         int
	Freq          int
	Desc          string
}

func (c *Defect) Init(v []string) {
	c.BrandName = strings.ToLower(v[0])
	c.ModelName = strings.ToLower(v[1])
	year, _ := strconv.ParseInt(v[2], 10, 32)
	c.ModelYear = int32(year)
	if len(v[3]) > 3 {
		c.Year, _ = strconv.Atoi(v[3][:4])
	}
	categories := removeEmptyStrings(strings.Split(strings.ToLower(v[4]), ":"))
	if l := len(categories); l < 3 {
		if l == 0 {
			return
		}
		add := []string{categories[0], categories[0]}
		categories = append(add, categories...)
	}
	c.MajorCategory = categories[0]
	c.MinorCategory = categories[1]
	c.Category = categories[2]

	mileage, _ := strconv.ParseUint(v[5], 10, 32)
	c.Miles = int(mileage)

	freq, _ := strconv.ParseUint(v[6], 10, 32)
	c.Freq = int(freq)
	c.Desc = strings.ToLower(v[7])
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
