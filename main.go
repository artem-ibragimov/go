package main

import (
	"main/automaniac"
	"main/db"
)

func main() {
	db := new(db.DB)
	err := db.Init()
	if err != nil {
		return
	}
	// car_complaints.ParseCarComplaints(db)
	automaniac.Parse(db)
}
