package db

// func (database *DB) GetVersion(model *ModelData) (int32, error) {
// 	var id sql.NullInt32
// 	tx, err := database.db.Begin()
// 	if err != nil {
// 		return 0, err
// 	}

// 	query, err := tx.Prepare(`SELECT id FROM model WHERE name = $1 AND version = $2 AND brand_id = $3`)
// 	if err != nil {
// 		return 0, err
// 	}

// 	err = query.QueryRow(
// 		model.Name,
// 		model.Version,
// 		model.BrandID,
// 	).Scan(&id)
// 	if err != nil {
// 		return 0, err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return 0, err
// 	}

// 	defer query.Close()
// 	return id.Int32, nil
// }

func (database *DB) SaveVersion(version *VersionData) (int32, error) {
	return database.Exec(`INSERT INTO version (
		name, model_id, transmission_id, engine_id
		) VALUES ( $1, $2, $3, $4 ) ON CONFLICT DO NOTHING RETURNING id`,
		version.Name,
		version.ModelID,
		version.TransID,
		version.EngineID,
	)
}

type VersionData struct {
	Name     string
	ModelID  int32
	BrandID  int32
	EngineID int32
	TransID  int32
}
