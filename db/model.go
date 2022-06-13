package db

func (database *DB) GetModel(brand_id int32, model_name string) (int32, error) {
	return database.Exec(`SELECT id FROM model WHERE name = $1 AND brand_id = $2`, model_name, brand_id)
}

func (database *DB) GetModelNamesByBrand(brand_id int32) ([]string, error) {
	return database.ExecRows(`SELECT name FROM model WHERE brand_id = $1 ORDER BY id DESC`, brand_id)
}

func (database *DB) SaveModel(model *ModelData) (int32, error) {
	return database.Exec(`INSERT INTO model (
		name, brand_id
		) VALUES ( $1, $2 ) ON CONFLICT DO NOTHING RETURNING id`,
		model.Name,
		model.BrandID,
	)
}

type ModelData struct {
	Name    string
	BrandID int32
}
