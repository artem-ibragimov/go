package db

import "log"

func (database *DB) GetEngineID(name string) (int32, error) {
	return database.Exec(`SELECT id FROM engine WHERE name= $1 `, name)
}
func (database *DB) GetEngines() (map[string]string, error) {
	return database.ExecMapRows(`SELECT id, name FROM engine`)
}

func (database *DB) GetEngine(engineID int32) (map[string]string, error) {
	query, err := database.db.Prepare(`SELECT name, displacement, cylinders, valves, fuel_type, power_hp, torque, img FROM engine WHERE id = $1`)
	if err != nil {
		log.Println(err)
	}
	defer query.Close()

	if err != nil {
		log.Println(err)
	}
	var results map[string]string = make(map[string]string)
	var img interface{}
	var name, displacement, cylinders, valves, fuel_type, power_hp, torque string
	err = query.QueryRow(engineID).Scan(&name, &displacement, &cylinders, &valves, &fuel_type, &power_hp, &torque, &img)
	if err != nil {
		log.Fatal(err)
		return make(map[string]string), err
	}
	results["name"] = name
	results["displacement"] = displacement
	results["cylinders"] = cylinders
	results["valves"] = valves
	results["fuel_type"] = fuel_type
	results["power_hp"] = power_hp
	results["torque"] = torque
	// results["img"] = img
	return results, nil
}

func (database *DB) GetEngineByParams(displacement int, valves int, power_hp int, torque int) (int32, error) {
	return database.Exec(`SELECT id FROM engine WHERE displacement= $1 AND valves=$2 AND power_hp=$3 AND torque=$4`,
		displacement, valves, power_hp, torque)
}

func (database *DB) SaveEngine(engine *EngineData) (int32, error) {
	return database.Exec(`INSERT INTO engine (
		name, displacement, cylinders, valves, fuel_type, power_hp, torque, img
		) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8 ) 
		ON CONFLICT (name) DO UPDATE 
		SET (
		name, displacement, cylinders, valves, fuel_type, power_hp, torque, img
		) = ( $1, $2, $3, $4, $5, $6, $7, $8 )
		RETURNING id`,
		engine.Name,
		engine.Displacement,
		engine.Cylinders,
		engine.Valves,
		engine.Fuel_type,
		engine.Power_hp,
		engine.Torque,
		engine.Img,
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
	Img          string
}
