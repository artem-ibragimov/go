package db

import "database/sql"

func (database *DB) GetDefectCategory(category string) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}
	err = tx.QueryRow(`SELECT id FROM defect_category WHERE name = $1`, category).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id.Int32, nil
}

func (database *DB) SaveDefectCategory(category string) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}

	query, err := tx.Prepare(`INSERT INTO defect_category ( name ) VALUES ( $1 ) ON CONFLICT DO NOTHING RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(category).Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id.Int32, nil
}
