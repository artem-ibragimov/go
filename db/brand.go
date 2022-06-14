package db

func (database *DB) GetBrand(brand string) (int32, error) {
	return database.Exec(`SELECT id FROM brand WHERE name= $1 `, brand)
}

func (database *DB) GetLastBrands() ([]string, error) {
	return database.ExecRows(`SELECT name FROM brand ORDER BY id DESC`)
}

func (database *DB) GetBrands() ([]string, error) {
	return database.ExecRows(`SELECT name FROM brand ORDER BY id INC`)
}

func (database *DB) SaveBrand(brand string) (int32, error) {
	return database.Exec(`INSERT INTO brand (name) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`, brand)
}
