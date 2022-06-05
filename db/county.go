package db

func (database *DB) GetCountry(country string) (int32, error) {
	return database.Exec(`SELECT id FROM country WHERE name= $1 `, country)
}

func (database *DB) SaveCountry(country string) (int32, error) {
	return database.Exec(`INSERT INTO country (name) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`, country)
}
