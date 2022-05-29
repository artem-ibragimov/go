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
		return 0, err
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

	query, err := tx.Prepare(`INSERT INTO engines (
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
		return 0, err
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

	query, err := tx.Prepare(`INSERT INTO transmissions (
		name, consumption, acceleration
		) VALUES ( $1, $2, $3 ) ON CONFLICT DO NOTHING RETURNING id`)
	if err != nil {
		return 0, err
	}

	err = query.QueryRow(
		trans.Name,
		trans.Consumtion,
		trans.Acceleration,
	).Scan(&id)
	if err != nil {
		return 0, err
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
		name, version, brand_id, year, transmissoin_id, engine_id
		) VALUES ( $1, $2, $3 ) ON CONFLICT DO NOTHING RETURNING id`)
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
		return 0, err
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
	Displacement string
	Config       string
	Valves       string
	Aspiration   string
	Fuel_type    string
	Power_hp     string
	Torque       string
}

type TransmissionData struct {
	Name         string
	Consumtion   string
	Acceleration string
}

type ModelData struct {
	Name     string
	Version  string
	Year     string
	BrandID  int32
	EngineID int32
	TransID  int32
}
