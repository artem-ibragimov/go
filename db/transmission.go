package db

func (database *DB) GetTransmission(brand_id int32, name string, gears int32) (int32, error) {
	return database.Exec(
		`SELECT id FROM transmission WHERE brand_id = $1 AND type = $2 AND gears = $3`,
		brand_id,
		name,
		gears,
	)
}

func (database *DB) SaveTransmission(trans *TransmissionData) (int32, error) {
	return database.Exec(`INSERT INTO transmission (
		brand_id, type, consumption, acceleration, gears
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
