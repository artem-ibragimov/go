package main

import (
	"log"
	"main/autokatalog"
	"main/db"
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
	autokatalog.Parse(db, req)
	// car_complaints.ParseCarComplaints(db)
	// automaniac.Parse(db, req)
	// nhtsa.Parse(db)
	// carproblemzoo.Parse(db, req)
	log.Println("Done.")
	db.Close()
}
