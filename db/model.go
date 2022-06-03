package db

import "database/sql"

func (database *DB) GetModel(brand_id int32, model_name string) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}

	query, err := tx.Prepare(`SELECT id FROM model WHERE name = $1 AND brand_id = $2`)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(model_name, brand_id).Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id.Int32, nil
}

func (database *DB) SaveModel(model *ModelData) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}

	query, err := tx.Prepare(`INSERT INTO model (
		name, brand_id, year
		) VALUES ( $1, $2, $3 ) ON CONFLICT DO NOTHING RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(
		model.Name,
		model.BrandID,
		model.Year,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id.Int32, nil
}

type ModelData struct {
	Name    string
	BrandID int32
	Year    int32
}
