package db

func (database *DB) PostSales(data *SaleData) (int32, error) {
	return database.Exec(`INSERT INTO sales (gen_id, country_id, year, amount) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING RETURNING id`, data.GenID, data.CountryID, data.Year, data.Amount)
}

type SaleData struct {
	GenID     int32
	CountryID int32
	Year      int
	Amount    int
}
