package db

import "log"

func (database *DB) GetTransmissionID(brandID int32, name string, gears int32) (int32, error) {
	return database.Exec(
		`SELECT id FROM transmission WHERE brand_id = $1 AND name = $2 AND gears = $3`,
		brandID,
		name,
		gears,
	)
}

func (database *DB) GetTransmissions(brandID int32) (map[string]string, error) {
	return database.ExecMapRows(`SELECT id, name FROM transmission WHERE brand_id = $1`, brandID)
}

func (database *DB) GetTransmission(id int32) (*TransmissionData, error) {
	query, err := database.db.Prepare(`SELECT brand_id, name, consumption, acceleration, gears FROM transmission WHERE id = $1`)
	if err != nil {
		log.Println(err)
	}
	defer query.Close()

	if err != nil {
		log.Println(err)
	}
	var brandID int32
	var gears int
	var name string
	var consumption, acceleration float32
	err = query.QueryRow(id).Scan(&brandID, &name, &consumption, &acceleration, &gears)
	if err != nil {
		log.Fatal(err)
		return &TransmissionData{}, err
	}

	return &TransmissionData{
		BrandID:      brandID,
		Name:         name,
		Gears:        gears,
		Consumtion:   consumption,
		Acceleration: acceleration,
	}, nil
}

func (database *DB) SaveTransmission(trans *TransmissionData) (int32, error) {
	return database.Exec(`INSERT INTO transmission (
		brand_id, name, consumption, acceleration, gears
		) VALUES ( $1, $2, $3, $4, $5 ) ON CONFLICT DO NOTHING RETURNING id`,
		trans.BrandID,
		trans.Name,
		trans.Consumtion,
		trans.Acceleration,
		trans.Gears,
	)
}

type TransmissionData struct {
	BrandID      int32
	Name         string
	Gears        int
	Consumtion   float32
	Acceleration float32
}
