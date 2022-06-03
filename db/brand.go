package db

import "database/sql"

func (database *DB) GetBrand(brand string) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}
	query, err := tx.Prepare(`SELECT id FROM brand WHERE name= $1 `)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(brand).Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id.Int32, nil
}

func (database *DB) SaveBrand(brand string) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}
	query, err := tx.Prepare(`INSERT INTO brand (name) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(brand).Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id.Int32, nil
}
