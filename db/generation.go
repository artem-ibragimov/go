package db

import "log"

func (database *DB) GetGenerationID(model_id int32, name string) (int32, error) {
	return database.Exec(`SELECT id FROM generation WHERE name = $1 AND model_id = $2 `, name, model_id)
}
func (database *DB) GetGenerations(model_id int32) (map[string]string, error) {
	return database.ExecMapRows(`SELECT id, name FROM generation WHERE model_id = $1`, model_id)
}
func (database *DB) GetGeneration(gen_id int32) (*GenerationData, error) {
	query, err := database.db.Prepare(`SELECT name, img, start, finish FROM generation WHERE id = $1`)
	if err != nil {
		log.Println(err)
	}
	defer query.Close()

	if err != nil {
		log.Println(err)
	}
	var name, img string
	var start, finish int
	err = query.QueryRow(gen_id).Scan(&name, &img, &start, &finish)
	if err != nil {
		log.Fatal(err)
		return new(GenerationData), err
	}
	return &GenerationData{
		Name:   name,
		Img:    img,
		Start:  start,
		Finish: finish,
	}, nil
}
func (database *DB) GetGenerationByStartYear(model_id int32, start int32) (int32, error) {
	return database.Exec(`SELECT id FROM generation WHERE model_id = $1 AND start = $2`, model_id, start)
}

func (database *DB) SearchGenerations(query string, limit uint) (map[string]string, error) {
	return database.ExecMapRows(`
	SELECT 
	generation.id, model.name || ' ' || generation.name 
	FROM 
		generation 
	LEFT JOIN 
		model ON model.id = generation.model_id 
	WHERE 
		generation.name 
	LIKE 
		$1
	LIMIT $2`, query+"%", limit)
}

// func (database *DB) GetGenerationStartByYear(model_id int32, year int32) (int32, error) {
// 	database.ExecRows(`SELECT id FROM generation WHERE model_id = $1 AND start = $2`, model_id, start)
// }

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
