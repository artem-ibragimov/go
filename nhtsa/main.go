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
	SaveDefectCategory(category string) (int32, error)
	SaveDefect(d *DB.Defect) (int32, error)
	GetBrand(brand string) (int32, error)
	SaveBrand(brand string) (int32, error)
	GetModel(brand_id int32, model_name string) (int32, error)
	SaveModel(model *DB.ModelData) (int32, error)
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

	country_id, err := db.GetCountry("USA")
	if err != nil {
		country_id, err = db.SaveCountry("USA")
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
		if defect.ModelYear < 2100 && defect.ModelYear >= 2000 {
			_, err = storeDefect(db, country_id, defect)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}

func storeDefect(db IDB, country_id int32, defect *Defect) (int32, error) {

	major_category_id, err := db.GetDefectCategory(defect.MajorCategory)
	if err != nil {
		major_category_id, err = db.SaveDefectCategory(defect.MajorCategory)
		if err != nil {
			return 0, err
		}
	}

	var minor_category_id = major_category_id

	if defect.MinorCategory != defect.MajorCategory {
		minor_category_id, err = db.GetDefectCategory(defect.MinorCategory)
		if err != nil {
			minor_category_id, err = db.SaveDefectCategory(defect.MinorCategory)
			if err != nil {
				return 0, err
			}
		}
	}

	var category_id = minor_category_id

	if defect.Category != defect.MinorCategory {
		category_id, err = db.GetDefectCategory(defect.Category)
		if err != nil {
			category_id, err = db.SaveDefectCategory(defect.Category)
			if err != nil {
				return 0, err
			}
		}
	}

	brand_id, err := db.GetBrand(defect.BrandName)
	if err != nil {
		brand_id, err = db.SaveBrand(defect.BrandName)
		log.Println("Saved", defect.BrandName)
		if err != nil {
			return 0, err
		}
	}

	model_id, err := db.GetModel(brand_id, defect.ModelName)
	if err != nil {
		model := &DB.ModelData{
			Name:    defect.ModelName,
			BrandID: brand_id,
			Year:    defect.ModelYear,
		}
		model_id, err = db.SaveModel(model)
		log.Println("Saved", model)
		if err != nil {
			return 0, err
		}
	}

	return db.SaveDefect(&DB.Defect{
		BrandID:         brand_id,
		ModelID:         model_id,
		Year:            defect.Year,
		MajorCategoryID: major_category_id,
		MinorCategoryID: minor_category_id,
		CategoryID:      category_id,
		Cost:            "0",
		Rating:          0,
		Miles:           defect.Miles,
		Freq:            defect.Freq,
		Desc:            defect.Desc,
		Country_ID:      country_id,
	})
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
	categories := strings.Split(strings.ToLower(v[4]), ":")
	if l := len(categories); l < 3 {
		add := []string{categories[0], categories[0]}
		categories = append(add, categories...)
	}
	c.MajorCategory = categories[0]
	c.MinorCategory = categories[1]
	c.Category = categories[2]

	miles, _ := strconv.ParseUint(v[5], 10, 32)
	c.Miles = int(miles)

	freq, _ := strconv.ParseUint(v[6], 10, 32)
	c.Freq = int(freq)
	c.Desc = strings.ToLower(v[7])
}
