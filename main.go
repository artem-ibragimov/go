package main

import (
	// "main/automaniac"
	"main/db"
	"main/nhtsa"
	"main/req"
)

func main() {
	db := new(db.DB)
	err := db.Init()
	if err != nil {
		return
	}
	req := new(req.Req)
	req.Init()
	// car_complaints.ParseCarComplaints(db)
	// automaniac.Parse(db, req)
	nhtsa.Parse()
}
