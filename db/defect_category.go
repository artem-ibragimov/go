package db

import "strings"

func (database *DB) GetDefectCategory(category string) (int32, error) {
	return database.Exec(
		`SELECT id FROM defect_category WHERE name = $1`,
		category)
}
func (database *DB) GetDefectCategories() (map[string]string, error) {
	return database.ExecMapRows(
		`SELECT id, name FROM defect_category`)
}

func (database *DB) PostDefectCategory(category string) (int32, error) {
	return database.Exec(
		`INSERT INTO defect_category ( name ) VALUES ( $1 ) ON CONFLICT DO NOTHING RETURNING id`,
		category)
}

func (database *DB) FindDefectCategory(c string) string {
	category := strings.TrimSpace(strings.ToLower(c))
	if strings.Contains(category, "engine") || strings.Contains(category, "oil") {
		return "engine"
	}
	if strings.Contains(category, "equipment") {
		return "equipment"
	}
	if strings.Contains(category, "transmission") || strings.Contains(category, "train") {
		return "transmission"
	}
	if strings.Contains(category, "bag") || strings.Contains(category, "emergency") {
		return "safety"
	}
	if strings.Contains(category, "light") || strings.Contains(category, "beam") {
		return "light"
	}
	if strings.Contains(category, "suspension") {
		return "suspension"
	}
	if strings.Contains(category, "brake") {
		return "brake"
	}
	if strings.Contains(category, "battery") ||
		strings.Contains(category, "wiring") ||
		strings.Contains(category, "horn") ||
		strings.Contains(category, "detection") ||
		strings.Contains(category, "control") ||
		strings.Contains(category, "navigation") ||
		strings.Contains(category, "radio") ||
		strings.Contains(category, "autonomous") ||
		strings.Contains(category, "software") ||
		strings.Contains(category, "assist") ||
		strings.Contains(category, "warning") ||
		strings.Contains(category, "cruise") ||
		strings.Contains(category, "instrument") ||
		strings.Contains(category, "electrical") ||
		strings.Contains(category, "electronics") {
		return "electronics"
	}
	return "other"
}
