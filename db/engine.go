package db

func (database *DB) GetEngine(engine string) (int32, error) {
	return database.Exec(`SELECT id FROM engine WHERE name= $1 `, engine)
}

func (database *DB) SaveEngine(engine *EngineData) (int32, error) {
	return database.Exec(`INSERT INTO engine (
		name, displacement, config, valves, aspiration, fuel_type, power_hp, torque
		) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8 ) ON CONFLICT DO NOTHING RETURNING id`,
		engine.Name,
		engine.Displacement,
		engine.Config,
		engine.Valves,
		engine.Aspiration,
		engine.Fuel_type,
		engine.Power_hp,
		engine.Torque,
	)
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
