package db

func (database *DB) PostDefect(d *Defect) (int32, error) {
	return database.Exec(`INSERT INTO defect (
		brand_id, model_id, description,
		mileage, cost, country_id,
		rating, age, freq,
		major_category_id, minor_category_id, category_id
		) 
		VALUES (
			$1, $2, $3,
			$4, $5, $6,
			$7, $8, $9,
			$10, $11, $12
		) ON CONFLICT DO NOTHING RETURNING id`,
		d.BrandID, d.ModelID, d.Desc,
		d.Mileage, d.Cost, d.Country_ID,
		d.Rating, d.Age, d.Freq,
		d.MajorCategoryID, d.MinorCategoryID, d.CategoryID,
	)
}

type Defect struct {
	BrandID         int32
	ModelID         int32
	GenID           int32
	VersionID       int32
	Age             int
	MajorCategoryID int32
	MinorCategoryID int32
	CategoryID      int32
	Country_ID      int32
	Cost            float32
	Rating          float32
	Mileage         int
	Freq            int
	Desc            string
}
