package db

import "database/sql"

func (database *DB) GetTransmission(trans *TransmissionData) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}

	query, err := tx.Prepare(`SELECT id FROM transmission WHERE name= $1 AND type=$2`)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(trans.BrandID, trans.Type).Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return id.Int32, nil
}

func (database *DB) SaveTransmission(trans *TransmissionData) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}

	query, err := tx.Prepare(`INSERT INTO transmission (
		brand_id, engine_id, type, consumption, acceleration, gears
		) VALUES ( $1, $2, $3, $4, $5, $6 ) ON CONFLICT DO NOTHING RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(
		trans.BrandID,
		trans.EngineID,
		trans.Type,
		trans.Consumtion,
		trans.Acceleration,
		trans.Gears,
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

type TransmissionData struct {
	BrandID      int32
	EngineID     int32
	Type         string
	Gears        int32
	Consumtion   float32
	Acceleration float32
}
