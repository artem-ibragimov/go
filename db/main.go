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
func (database *DB) ExecRows(cmd string, args ...interface{}) ([]string, error) {
	query, err := database.db.Prepare(cmd)
	if err != nil {
		log.Println(err)
	}
	defer query.Close()
	rows, err := query.Query(args...)
	if err != nil {
		log.Println(err)
	}

	var results []string
	var s string
	for rows.Next() {
		err := rows.Scan(&s)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, s)
	}
	err = rows.Err()
	if err != nil {
		return []string{}, err
	}
	return results, nil
}
func (database *DB) ExecMap(cmd string, args ...interface{}) (map[string]string, error) {
	query, err := database.db.Prepare(cmd)
	if err != nil {
		log.Println(err)
	}
	defer query.Close()
	rows, err := query.Query(args...)
	if err != nil {
		log.Println(err)
	}

	var results map[string]string = make(map[string]string)
	var k, v string
	for rows.Next() {
		err := rows.Scan(&k, &v)
		if err != nil {
			log.Fatal(err)
		}
		results[k] = v
	}
	err = rows.Err()
	if err != nil {
		return make(map[string]string), err
	}
	return results, nil
}
