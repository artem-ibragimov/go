package db

import "log"

func (database *DB) GetVersionID(name string, generation_id int32) (int32, error) {
	return database.Exec(
		`SELECT id FROM version WHERE name = $1 AND generation_id = $2`,
		name, generation_id)
}

func (database *DB) GetVersion(version_id int32) (map[string]string, error) {
	query, err := database.db.Prepare(`SELECT name, transmission_id, engine_id FROM version WHERE id = $1`)
	if err != nil {
		log.Println(err)
	}
	defer query.Close()

	if err != nil {
		log.Println(err)
	}
	var results map[string]string = make(map[string]string)
	var name, transmissionID, engineID string
	err = query.QueryRow(version_id).Scan(&name, &transmissionID, &engineID)
	if err != nil {
		log.Fatal(err)
		return make(map[string]string), err
	}
	results["name"] = name
	results["transmissionID"] = transmissionID
	results["engineID"] = engineID
	return results, nil
}

func (database *DB) GetVersions(generation_id int32) (map[string]string, error) {
	return database.ExecMapRows(
		`SELECT id, name FROM version WHERE generation_id = $1`,
		generation_id)
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
