package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func (database *DB) Init() error {
	db, err := sql.Open("postgres", "user=artem dbname=artem password=123 sslmode=disable")
	if err != nil {
		return err
	}
	database.db = db
	return nil
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

	err = query.QueryRow(brand).Scan(&id)
	if err != nil {
		tx.QueryRow(`SELECT id FROM brand WHERE name= $1 `, brand).Scan(&id)
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	defer query.Close()
	return id.Int32, nil
}

func (database *DB) SaveEngine(engine *EngineData) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}

	query, err := tx.Prepare(`INSERT INTO engine (
		name, displacement, config, valves, aspiration, fuel_type, power_hp, torque
		) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8 ) ON CONFLICT DO NOTHING RETURNING id`)
	if err != nil {
		return 0, err
	}

	err = query.QueryRow(
		engine.Name,
		engine.Displacement,
		engine.Config,
		engine.Valves,
		engine.Aspiration,
		engine.Fuel_type,
		engine.Power_hp,
		engine.Torque,
	).Scan(&id)
	if err != nil {
		tx.QueryRow(`SELECT id FROM engine WHERE name= $1 `, engine.Name).Scan(&id)
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	defer query.Close()
	return id.Int32, nil
}
func (database *DB) SaveTransmission(trans *TransmissionData) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}

	query, err := tx.Prepare(`INSERT INTO transmission (
		brand_id, description, consumption, acceleration
		) VALUES ( $1, $2, $3, $4 ) ON CONFLICT DO NOTHING RETURNING id`)
	if err != nil {
		return 0, err
	}

	err = query.QueryRow(
		trans.BrandID,
		trans.Desc,
		trans.Consumtion,
		trans.Acceleration,
	).Scan(&id)
	if err != nil {
		tx.QueryRow(`SELECT id FROM transmission WHERE name= $1 AND description=$2`, trans.BrandID, trans.Desc).Scan(&id)
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	defer query.Close()
	return id.Int32, nil
}
func (database *DB) SaveModel(model *ModelData) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}

	query, err := tx.Prepare(`INSERT INTO model (
		name, version, brand_id, year, transmission_id, engine_id
		) VALUES ( $1, $2, $3, $4, $5, $6 ) ON CONFLICT DO NOTHING RETURNING id`)
	if err != nil {
		return 0, err
	}

	err = query.QueryRow(
		model.Name,
		model.Version,
		model.BrandID,
		model.Year,
		model.TransID,
		model.EngineID,
	).Scan(&id)
	if err != nil {
		tx.QueryRow(`SELECT id FROM model WHERE name = $1 AND version = $2 AND brand_id = $3`,
			model.Name,
			model.Version,
			model.BrandID,
		).Scan(&id)
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	defer query.Close()
	return id.Int32, nil
}

type EngineData struct {
	Name         string
	Displacement int32
	Config       string
	Valves       string
	Aspiration   float32
	Fuel_type    string
	Power_hp     int32
	Torque       int32
}

type TransmissionData struct {
	BrandID      int32
	Desc         string
	Consumtion   float32
	Acceleration float32
}

type ModelData struct {
	Name     string
	Version  string
	Year     int32
	BrandID  int32
	EngineID int32
	TransID  int32
}
