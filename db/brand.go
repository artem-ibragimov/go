package db

func (database *DB) GetBrandByName(brand string) (int32, error) {
	return database.Exec(`SELECT id FROM brand WHERE name= $1 `, brand)
}

func (database *DB) GetLastBrands() ([]string, error) {
	return database.ExecRows(`SELECT name FROM brand ORDER BY id DESC`)
}

func (database *DB) GetBrands() (map[string]string, error) {
	return database.ExecMapRows(`SELECT id, name FROM brand`)
}
func (database *DB) SearchBrands(query string, limit uint) (map[string]string, error) {
	return database.ExecMapRows(`SELECT id, name FROM brand WHERE name LIKE $1 LIMIT $2`, query+"%", limit)
}

func (database *DB) PostBrand(brand string) (int32, error) {
	return database.Exec(`INSERT INTO brand (name) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`, brand)
}
