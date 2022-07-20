package db

func (database *DB) GetVersionID(name string, gen_id int32) (int32, error) {
	return database.Exec(
		`SELECT id FROM version WHERE name = $1 AND gen_id = $2`,
		name, gen_id)
}

func (database *DB) GetVersion(version_id int32) (*VersionData, error) {
	query, err := database.db.Prepare(`SELECT name, trans_id, engine_id, gen_id FROM version WHERE id = $1`)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	if err != nil {
		return nil, err
	}
	var name string
	var transID, engineID, generationID int32
	err = query.QueryRow(version_id).Scan(&name, &transID, &engineID, &generationID)
	if err != nil {
		return nil, err
	}
	return &VersionData{
		Name:     name,
		EngineID: engineID,
		TransID:  transID,
		GenID:    generationID,
	}, nil
}

func (database *DB) GetVersions(gen_id int32) (map[string]string, error) {
	return database.ExecMapRows(
		`SELECT id, name FROM version WHERE gen_id = $1`,
		gen_id)
}

func (database *DB) PostVersion(version *VersionData) (int32, error) {
	return database.Exec(`INSERT INTO version (
		name, gen_id, trans_id, engine_id
		) VALUES ( $1, $2, $3, $4 ) ON CONFLICT DO NOTHING RETURNING id`,
		version.Name,
		version.GenID,
		version.TransID,
		version.EngineID,
	)
}
func (database *DB) PatchVersion(id int32, version *VersionData) (int32, error) {
	return database.Exec(`
	UPDATE
		version 
	SET
		name = $1, 
		gen_id = $2,
		trans_id = $3,
		engine_id = $4
	WHERE
		id = $5
	RETURNING id`,
		version.Name,
		version.GenID,
		version.TransID,
		version.EngineID,
		id,
	)
}

type VersionData struct {
	Name     string
	GenID    int32
	EngineID int32
	TransID  int32
}
