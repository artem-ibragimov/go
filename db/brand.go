package db

func (database *DB) GetBrand(brand string) (int32, error) {
	return database.Exec(`SELECT id FROM brand WHERE name= $1 `, brand)
}

func (database *DB) GetLastBrands() ([]string, error) {
	return database.ExecRows(`SELECT name FROM brand ORDER BY id DESC`)
}

func (database *DB) GetBrands() (map[string]string, error) {
	return database.ExecMap(`SELECT id, name FROM brand`)
}

func (database *DB) SaveBrand(brand string) (int32, error) {
	return database.Exec(`INSERT INTO brand (name) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`, brand)
}
