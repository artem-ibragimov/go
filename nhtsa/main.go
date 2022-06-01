package nhtsa

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Compilant struct {
	BrandName string
	ModelName string
	ModelYear sql.NullInt32
	Miles     sql.NullInt32
	Freq      sql.NullInt32
	Desc      string
}

func (c *Compilant) Init(v []string) {
	c.BrandName = v[0]
	c.ModelName = v[1]
	year, err := strconv.ParseInt(v[2], 10, 32)
	c.ModelYear = sql.NullInt32{Int32: int32(year), Valid: err == nil}

	miles, err := strconv.ParseUint(v[3], 10, 32)
	c.Miles = sql.NullInt32{Int32: int32(miles), Valid: err == nil}

	freq, err := strconv.ParseUint(v[4], 10, 32)
	c.Freq = sql.NullInt32{Int32: int32(freq), Valid: err == nil}
	c.Desc = v[5]
	// TODO
}

func Parse() {

	// open file
	f, err := os.Open("nhtsa\\example.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
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
		c := new(Compilant)
		c.Init(rec)
		fmt.Printf("%+v\n", c)
	}
}
