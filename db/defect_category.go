package db

func (database *DB) GetDefectCategory(category string) (int32, error) {
	return database.Exec(
		`SELECT id FROM defect_category WHERE name = $1`,
		category)
}

func (database *DB) PostDefectCategory(category string) (int32, error) {
	return database.Exec(
		`INSERT INTO defect_category ( name ) VALUES ( $1 ) ON CONFLICT DO NOTHING RETURNING id`,
		category)
}
