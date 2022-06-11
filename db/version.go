package db

func (database *DB) GetVersion(name string, generation_id int32) (int32, error) {
	return database.Exec(
		`SELECT id FROM version WHERE name = $1 AND generation_id = $2`,
		name, generation_id)
}

func (database *DB) SaveVersion(version *VersionData) (int32, error) {
	return database.Exec(`INSERT INTO version (
		name, generation_id, transmission_id, engine_id
		) VALUES ( $1, $2, $3, $4 ) ON CONFLICT DO NOTHING RETURNING id`,
		version.Name,
		version.GenerationID,
		version.TransID,
		version.EngineID,
	)
}

type VersionData struct {
	Name         string
	GenerationID int32
	EngineID     int32
	TransID      int32
}
