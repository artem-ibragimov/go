package db

func (database *DB) GetModelID(brand_id int32, model_name string) (int32, error) {
	return database.Exec(`SELECT id FROM model WHERE name = $1 AND brand_id = $2`, model_name, brand_id)
}

func (database *DB) GetLastModelNamesByBrand(brand_id int32) ([]string, error) {
	return database.ExecRows(`SELECT name FROM model WHERE brand_id = $1 ORDER BY id DESC`, brand_id)
}
func (database *DB) GetModelsByBrand(brand_id int32) (map[string]string, error) {
	return database.ExecMapRows(`SELECT id, name FROM model WHERE brand_id = $1 `, brand_id)
}
func (database *DB) SearchModels(query string, limit uint) (map[string]string, error) {
	return database.ExecMapRows(`
	SELECT 
		model.id, brand.name || ' ' || model.name 
	FROM 
		model 
	LEFT JOIN 
		brand ON brand.id = model.brand_id 
	WHERE 
		model.name 
	LIKE 
		$1
	LIMIT $2`, query+"%", limit)
}
func (database *DB) PostModel(model *ModelData) (int32, error) {
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
