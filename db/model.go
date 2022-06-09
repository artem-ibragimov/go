package db

import (
	"sort"
	"strconv"
)

func (database *DB) GetModel(brand_id int32, model_name string, year int32) (int32, error) {
	return database.Exec(`SELECT id FROM model WHERE name = $1 AND brand_id = $2 AND year = $3`, model_name, brand_id, year)
}

func (database *DB) GetModelYears(brand_id int32, model_name string) ([]int, error) {
	years, err := database.ExecRows(`SELECT year FROM model WHERE name = $1 AND brand_id = $2 `, model_name, brand_id)
	if err != nil {
		return []int{}, err
	}
	result := make([]int, len(years))
	for i, y := range years {
		v, _ := strconv.Atoi(y)
		result[i] = int(v)
	}
	sort.Ints(result)
	return result, nil
}

func (database *DB) SaveModel(model *ModelData) (int32, error) {
	return database.Exec(`INSERT INTO model (
		name, brand_id, year
		) VALUES ( $1, $2, $3 ) ON CONFLICT DO NOTHING RETURNING id`,
		model.Name,
		model.BrandID,
		model.Year,
	)
}

type ModelData struct {
	Name    string
	BrandID int32
	Year    int32
}
