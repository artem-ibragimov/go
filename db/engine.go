package db

import (
	"fmt"
)

func (database *DB) GetEngineID(name string) (int32, error) {
	return database.Exec(`SELECT id FROM engine WHERE name= $1 `, name)
}
func (database *DB) GetEngines() (map[string]string, error) {
	return database.ExecMapRows(`SELECT id, name FROM engine`)
}

func (database *DB) GetEngine(engineID int32) (*EngineData, error) {
	query, err := database.db.Prepare(`SELECT name, displacement, cylinders, valves, fuel_type, power_hp, torque, img FROM engine WHERE id = $1`)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	if err != nil {
		return nil, err
	}
	var img interface{}
	var name, fuel_type string
	var displacement, cylinders, valves, power_hp, torque int
	err = query.QueryRow(engineID).Scan(&name, &displacement, &cylinders, &valves, &fuel_type, &power_hp, &torque, &img)
	if err != nil {
		return nil, err
	}
	var img_s string
	if img != nil {
		img_s = fmt.Sprintf("%v", img)
	}
	return &EngineData{
		Name:         name,
		Displacement: displacement,
		Valves:       valves,
		Cylinders:    cylinders,
		Fuel_type:    fuel_type,
		Power_hp:     power_hp,
		Torque:       torque,
		Img:          img_s,
	}, nil
}

func (database *DB) GetEngineByParams(displacement int, valves int, power_hp int, torque int) (int32, error) {
	return database.Exec(`SELECT id FROM engine WHERE displacement= $1 AND valves=$2 AND power_hp=$3 AND torque=$4`,
		displacement, valves, power_hp, torque)
}

func (database *DB) PostEngine(engine *EngineData) (int32, error) {
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
func (database *DB) PatchEngine(id int32, engine *EngineData) (int32, error) {
	return database.Exec(`
	UPDATE 
		engine 
	SET (
		name, displacement, cylinders, valves, fuel_type, power_hp, torque, img
	) = ( $1, $2, $3, $4, $5, $6, $7, $8 )
	WHERE
		id = $9
	RETURNING id`,
		engine.Name,
		engine.Displacement,
		engine.Cylinders,
		engine.Valves,
		engine.Fuel_type,
		engine.Power_hp,
		engine.Torque,
		engine.Img,
		id,
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
