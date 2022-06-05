package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func (database *DB) Init() error {
	db, err := sql.Open("postgres", "user=artem dbname=artem password=123 sslmode=disable")
	if err != nil {
		return err
	}
	database.db = db
	return nil
}
func (database *DB) Close() error {
	return database.db.Close()
}

func (database *DB) Exec(cmd string, args ...interface{}) (int32, error) {
	var id sql.NullInt32
	query, err := database.db.Prepare(cmd)
	if err != nil {
		log.Println(err)
	}
	defer query.Close()

	err = query.QueryRow(args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id.Int32, nil
}
