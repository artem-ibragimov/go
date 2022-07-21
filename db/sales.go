package db

import "strconv"

func (database *DB) GetBrandSales(brand_id int32) (int, error) {
	results, err := database.ExecRows(`SELECT SUM(amount) FROM sales WHERE brand_id=$1`, brand_id)
	if len(results) == 0 || err != nil {
		return 0, err
	}
	return strconv.Atoi(results[0])
}
func (database *DB) PostSales(data *SaleData) (int32, error) {
	return database.Exec(`
		INSERT INTO 
			sales (brand_id, model_id, gen_id, country_id, year, amount) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT
			DO NOTHING
		RETURNING id`, data.BrandID, data.ModelID, data.GenID, data.CountryID, data.Year, data.Amount)
}

type SaleData struct {
	BrandID   int32
	ModelID   int32
	GenID     int32
	CountryID int32
	Year      int
	Amount    int
}
