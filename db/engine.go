package db

import "database/sql"

func (database *DB) GetEngine(engine string) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}
	query, err := tx.Prepare(`SELECT id FROM engine WHERE name= $1 `)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(engine).Scan(&id)

	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

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
	defer query.Close()

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

	return id.Int32, nil
}

type EngineData struct {
	Name         string
	Displacement int32
	Config       string
	Valves       string
	Aspiration   string
	Fuel_type    string
	Power_hp     int32
	Torque       int32
}
