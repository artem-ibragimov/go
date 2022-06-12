package db

func (database *DB) GetGeneration(model_id int32, name string) (int32, error) {
	return database.Exec(`SELECT id FROM generation WHERE name = $1 AND model_id = $2 `, name, model_id)
}
func (database *DB) GetGenerationByStartYear(model_id int32, start int32) (int32, error) {
	return database.Exec(`SELECT id FROM generation WHERE model_id = $1 AND start = $2`, model_id, start)
}

func (database *DB) SaveGeneration(data *GenerationData) (int32, error) {
	return database.Exec(`INSERT INTO generation (
		name, model_id,  img, start, finish
		) VALUES ( $1, $2, $3, $4, $5 ) ON CONFLICT DO NOTHING RETURNING id`,
		data.Name,
		data.ModelID,
		data.Img,
		data.Start,
		data.Finish,
	)
}

type GenerationData struct {
	Name    string
	ModelID int32
	Start   int
	Finish  int
	Img     string
}
