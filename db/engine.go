package db

func (database *DB) GetEngine(name string) (int32, error) {
	return database.Exec(`SELECT id FROM engine WHERE name= $1 `, name)
}
func (database *DB) GetEngineByParams(displacement int, valves int, power_hp int, torque int) (int32, error) {
	return database.Exec(`SELECT id FROM engine WHERE displacement= $1 AND valves=$2 AND power_hp=$3 AND torque=$4`,
		displacement, valves, power_hp, torque)
}

func (database *DB) SaveEngine(engine *EngineData) (int32, error) {
	return database.Exec(`INSERT INTO engine (
		name, displacement, cylinders, valves, fuel_type, power_hp, torque
		) VALUES ( $1, $2, $3, $4, $5, $6, $7 ) ON CONFLICT DO NOTHING RETURNING id`,
		engine.Name,
		engine.Displacement,
		engine.Cylinders,
		engine.Valves,
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
	Fuel_type    string
	Power_hp     int
	Torque       int
}
