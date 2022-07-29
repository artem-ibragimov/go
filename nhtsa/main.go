package nhtsa

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	DB "main/db"
	"os"
	"strconv"
	"strings"
)

type IDB interface {
	GetDefectCategory(category string) (int32, error)
	FindDefectCategory(category string) string
	PostDefectCategory(category string) (int32, error)
	PostDefect(d *DB.Defect) (int32, error)

	GetBrandByName(brand string) (int32, error)
	PostBrand(brand string) (int32, error)

	GetModelID(brand_id int32, model_name string) (int32, error)
	PostModel(model *DB.ModelData) (int32, error)

	GetCountry(country string) (int32, error)
	SaveCountry(country string) (int32, error)

	GetGenByStartYear(model_id int32, start int) (int32, error)
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
				fmt.Println(err)
			}
		}
	}
}

func storeDefect(db IDB, country_id int32, defect *Defect) (int32, error) {
	category := db.FindDefectCategory(defect.Category)
	major_category_id, err := db.GetDefectCategory(category)
	if err != nil {
		major_category_id, err = db.PostDefectCategory(category)
		if err != nil {
			return 0, err
		}
	}

	brand_id, err := db.GetBrandByName(defect.BrandName)
	if err != nil {
		return 0, err
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
	gen_id, _ := db.GetGenByStartYear(model_id, defect.ModelYear)
	_, err = db.PostDefect(&DB.Defect{
		BrandID:    brand_id,
		ModelID:    model_id,
		GenID:      gen_id,
		Age:        defect.Year - int(defect.ModelYear),
		Year:       defect.Year,
		CategoryID: major_category_id,
		Cost:       0,
		Rating:     0,
		Mileage:    defect.Miles,
		Freq:       defect.Freq,
		Desc:       defect.Desc,
		CountryID:  country_id,
	})
	return 0, err
}

type Defect struct {
	BrandName string
	ModelName string
	ModelYear int
	Year      int
	Category  string
	Miles     int
	Freq      int
	Desc      string
}

func (c *Defect) Init(v []string) {
	c.BrandName = strings.ToLower(v[0])
	c.ModelName = strings.ToLower(v[1])
	year, _ := strconv.ParseInt(v[2], 10, 32)
	c.ModelYear = int(year)
	if len(v[3]) > 3 {
		c.Year, _ = strconv.Atoi(v[3][:4])
	}
	categories := removeEmptyStrings(strings.Split(strings.ToLower(v[4]), ":"))
	if len(categories) == 0 {
		return
	}
	c.Category = categories[0]

	mileage, _ := strconv.ParseUint(v[5], 10, 32)
	c.Miles = (int(mileage)/25000 + 1) * 25000

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
