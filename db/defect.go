package db

import (
	"database/sql"
)

func (database *DB) SaveDefect(d *Defect) (int32, error) {
	var id sql.NullInt32
	tx, err := database.db.Begin()
	if err != nil {
		return 0, err
	}
	query, err := tx.Prepare(`INSERT INTO defect (
		brand_id,	model_id,
		miles,		desc,			cost, 
		rating,		date,			freq,
		major_category_id, minor_category_id, category_id
		) 
		VALUES (
			$1, $2, $3, 
			$4, $5, $6, 
			$7, $8, $9
		) ON CONFLICT DO NOTHING RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(
		d.BrandID, d.ModelID,
		d.Miles, d.Desc, d.Cost,
		d.Rating, d.Date, d.Freq,
		d.MajorCategoryID, d.MinorCategoryID, d.CategoryID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id.Int32, nil
}

type Defect struct {
	BrandID         int32
	ModelID         int32
	Date            string
	MajorCategoryID int32
	MinorCategoryID int32
	CategoryID      int32
	Cost            string
	Rating          float32
	Miles           sql.NullInt32
	Freq            sql.NullInt32
	Desc            string
}
