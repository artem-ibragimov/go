package db

import "database/sql"

func (database *DB) GetLastYearsDefect(brand_id int32, model_id int32) ([]string, error) {
	return database.ExecRows(`SELECT year FROM defect WHERE brand_id=$1 AND model_id=$2 GROUP BY year ORDER BY year DESC`, brand_id, model_id)
}

func (database *DB) GetDefectsAgesByBrand(brand_id int32) (map[string]string, error) {
	return database.ExecMapRows(
		`SELECT
			age,
			COALESCE( count(id)::decimal / (SELECT SUM(amount) as total FROM sales WHERE brand_id=$1) * 100000, 0) as n
		FROM
			defect
		WHERE
			brand_id=$1
		AND
			age>=0
		GROUP BY
			age
		ORDER BY
			age`, brand_id)
}

func (database *DB) GetDefectsAgesByModel(model_id int32) (map[string]string, error) {
	return database.ExecMapRows(
		`SELECT
			age,
			COALESCE( count(id)::decimal / (SELECT SUM(amount) as total FROM sales WHERE brand_id=$1) * 100000, 0) as n
		FROM
			defect
		WHERE
		model_id=$1
		AND
			age>=0
		GROUP BY
			age
		ORDER BY
			age`, model_id)
}
func (database *DB) GetDefectsAgesByGen(gen_id int32) (map[string]string, error) {
	return database.ExecMapRows(
		`SELECT
			age,
			COALESCE( count(id)::decimal / (SELECT SUM(amount) as total FROM sales WHERE brand_id=$1) * 100000, 0) as n
		FROM
			defect
		WHERE
		gen_id=$1
		AND
			age>=0
		GROUP BY
			age
		ORDER BY
			age`, gen_id)
}

func (database *DB) PostDefect(d *Defect) (int32, error) {
	var gen_id sql.NullInt32 = sql.NullInt32{Int32: d.GenID}
	if d.GenID == 0 {
		gen_id = sql.NullInt32{Valid: false}
	}
	var version_id sql.NullInt32 = sql.NullInt32{Int32: d.VersionID}
	if d.VersionID == 0 {
		version_id = sql.NullInt32{Valid: false}
	}
	return database.Exec(`INSERT INTO defect (
		brand_id, model_id, gen_id, version_id, description,
		mileage, cost, country_id,
		rating, age, year, freq,
		major_category_id, minor_category_id, category_id
		) 
		VALUES (
			$1, $2, $3,
			$4, $5, $6,
			$7, $8, $9,
			$10, $11, $12,
			$13, $14, $15
		) 
		ON CONFLICT (description) DO UPDATE 
		SET (
			brand_id, model_id, gen_id, version_id, description,
			mileage, cost, country_id,
			rating, age, year, freq,
			major_category_id, minor_category_id, category_id
		) = (
			$1, $2, $3,
			$4, $5, $6,
			$7, $8, $9,
			$10, $11, $12,
			$13, $14, $15
		)
		RETURNING id`,
		d.BrandID, d.ModelID, gen_id, version_id, d.Desc,
		d.Mileage, d.Cost, d.CountryID,
		d.Rating, d.Age, d.Year, d.Freq,
		d.MajorCategoryID, d.MinorCategoryID,
		d.CategoryID,
	)
}

type Defect struct {
	BrandID   int32
	ModelID   int32
	GenID     int32
	VersionID int32
	CountryID int32

	MajorCategoryID int32
	MinorCategoryID int32
	CategoryID      int32

	Age     int
	Year    int
	Cost    float32
	Rating  float32
	Mileage int
	Freq    int
	Desc    string
}
