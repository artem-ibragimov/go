package db

import (
	"database/sql"
	"strings"
)

func (database *DB) GetLastYearsDefect(brand_id int32, model_id int32) ([]string, error) {
	return database.ExecRows(`SELECT year FROM defect WHERE brand_id=$1 AND model_id=$2 GROUP BY year ORDER BY year DESC`, brand_id, model_id)
}

func (database *DB) GetDefectsAgesByBrand(brand_id int32, categories []string, norm bool) (map[string]string, error) {
	y := "count(id)::decimal as n"
	if norm {
		y = " count(id)::decimal / (SELECT SUM(amount) as total FROM sales WHERE brand_id=$1) * 100000  as n"
	}
	with_categories := ""
	if len(categories) > 1 || len(categories) == 1 && categories[0] != "" {
		with_categories = " AND (category_id=" + strings.Join(categories, " OR category_id=") + ")"
	}
	return database.ExecMapRows(
		`SELECT
			age,`+y+`
		FROM defect
		WHERE brand_id=$1 AND age>=0 AND age<=22`+with_categories+`
		GROUP BY age
		ORDER BY age`, brand_id)
}
func (database *DB) GetDefectsMileageByBrand(brand_id int32, norm bool) (map[string]string, error) {
	y := "count(id)::decimal as n"
	if norm {
		y = "COALESCE( count(defect.id)::decimal / (SELECT SUM(amount) as total FROM sales WHERE brand_id=$1) * 100000, 0) as n"
	}
	return database.ExecMapRows(
		`SELECT mileage,`+y+`
		FROM defect
		WHERE brand_id=$1 AND defect.mileage>0
		GROUP BY mileage
		ORDER BY mileage`, brand_id)
}

func (database *DB) GetDefectsAgesByModel(model_id int32, norm bool) (map[string]string, error) {
	if norm {
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
	return database.ExecMapRows(
		`SELECT
			age,
			count(id)::decimal as n
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
func (database *DB) GetDefectsMileageByModel(model_id int32, norm bool) (map[string]string, error) {
	y := "count(id)::decimal as n"
	if norm {
		y = "COALESCE( count(defect.id)::decimal / (SELECT SUM(amount) as total FROM sales WHERE model_id=$1) * 100000, 0) as n"
	}
	return database.ExecMapRows(
		`SELECT mileage,`+y+`
		FROM defect
		WHERE model_id=$1 AND defect.mileage>0
		GROUP BY mileage
		ORDER BY mileage`, model_id)
}

func (database *DB) GetDefectsAgesByGen(gen_id int32, norm bool) (map[string]string, error) {
	if norm {
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
	return database.ExecMapRows(
		`SELECT
			age,
			count(id)::decimal as n
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
func (database *DB) GetDefectsMileageByGen(gen_id int32, norm bool) (map[string]string, error) {
	y := "count(id)::decimal as n"
	if norm {
		y = "COALESCE( count(defect.id)::decimal / (SELECT SUM(amount) as total FROM sales WHERE gen_id=$1) * 100000, 0) as n"
	}
	return database.ExecMapRows(
		`SELECT mileage,`+y+`
		FROM defect
		WHERE gen_id=$1 AND defect.mileage>0
		GROUP BY mileage
		ORDER BY mileage`, gen_id)
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
		rating, age, year, freq, category_id
		) 
		VALUES (
			$1, $2, $3,
			$4, $5, $6,
			$7, $8, $9,
			$10, $11, $12,
			$13
		) 
		ON CONFLICT (description) DO UPDATE 
		SET (
			brand_id, model_id, gen_id, version_id, description,
			mileage, cost, country_id,
			rating, age, year, freq, category_id
		) = (
			$1, $2, $3,
			$4, $5, $6,
			$7, $8, $9,
			$10, $11, $12,
			$13
		)
		RETURNING id`,
		d.BrandID, d.ModelID, gen_id, version_id, d.Desc,
		d.Mileage, d.Cost, d.CountryID,
		d.Rating, d.Age, d.Year, d.Freq,
		d.CategoryID,
	)
}

type Defect struct {
	BrandID   int32
	ModelID   int32
	GenID     int32
	VersionID int32
	CountryID int32

	CategoryID int32

	Age     int
	Year    int
	Cost    float32
	Rating  float32
	Mileage int
	Freq    int
	Desc    string
}
