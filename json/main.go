package json

import (
	"fmt"
	DB "main/db"
	"os"
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
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	// var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	// json.Unmarshal(byteValue, &users)
}
