package nhtsa

import (
	"database/sql"
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
}

func Parse(db IDB) {
	f, err := os.Open("nhtsa\\example.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
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
		storeDefect(db, defect)
	}
}

func storeDefect(db IDB, defect *Defect) {

	major_category_id, err := db.GetDefectCategory(defect.MajorCategory)
	if err != nil {
		major_category_id, err = db.SaveDefectCategory(defect.MajorCategory)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	var minor_category_id int32 = major_category_id

	if defect.MinorCategory != defect.MajorCategory {
		minor_category_id, err = db.GetDefectCategory(defect.MinorCategory)
		if err != nil {
			minor_category_id, err = db.SaveDefectCategory(defect.MinorCategory)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}

	var category_id int32 = minor_category_id

	if defect.Category != defect.MinorCategory {
		category_id, err = db.GetDefectCategory(defect.Category)
		if err != nil {
			category_id, err = db.SaveDefectCategory(defect.Category)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}

	brand_id, err := db.GetBrand(defect.BrandName)
	if err != nil {
		brand_id, err = db.SaveBrand(defect.BrandName)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	db.SaveDefect(&DB.Defect{
		BrandID: brand_id,
	})
}

type Defect struct {
	BrandName string
	ModelName string
	ModelYear sql.NullInt32
	// YYYYMMDD
	Date          string
	MajorCategory string
	MinorCategory string
	Category      string
	Miles         sql.NullInt32
	Freq          sql.NullInt32
	Desc          string
}

func (c *Defect) Init(v []string) {
	c.BrandName = strings.ToLower(v[0])
	c.ModelName = v[1]
	year, err := strconv.ParseInt(v[2], 10, 32)
	c.ModelYear = sql.NullInt32{Int32: int32(year), Valid: err == nil}
	c.Date = v[3]
	categories := strings.Split(v[4], ":")
	if l := len(categories); l < 3 {
		add := []string{categories[0], categories[0]}
		categories = append(add, categories...)
	}
	c.MajorCategory = categories[0]
	c.MinorCategory = categories[1]
	c.Category = categories[2]

	miles, err := strconv.ParseUint(v[5], 10, 32)
	c.Miles = sql.NullInt32{Int32: int32(miles), Valid: err == nil}

	freq, err := strconv.ParseUint(v[6], 10, 32)
	c.Freq = sql.NullInt32{Int32: int32(freq), Valid: err == nil}
	c.Desc = v[7]
}
