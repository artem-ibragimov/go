package db

func (database *DB) GetEngine(name string) (int32, error) {
	return database.Exec(`SELECT id FROM engine WHERE name= $1 `, name)
}

func (database *DB) SaveEngine(engine *EngineData) (int32, error) {
	return database.Exec(`INSERT INTO engine (
		name, displacement, cylinders, valves, aspiration, fuel_type, power_hp, torque
		) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8 ) ON CONFLICT DO NOTHING RETURNING id`,
		engine.Name,
		engine.Displacement,
		engine.Cylinders,
		engine.Valves,
		engine.Aspiration,
		engine.Fuel_type,
		engine.Power_hp,
		engine.Torque,
	)
}

type EngineData struct {
	Name         string
	Displacement int
	Valves       int
	Cylinders    int
	Aspiration   string
	Fuel_type    string
	Power_hp     int
	Torque       int
}
