package db

func (database *DB) GetModel(brand_id int32, model_name string) (int32, error) {
	return database.Exec(`SELECT id FROM model WHERE name = $1 AND brand_id = $2`, model_name, brand_id)
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
