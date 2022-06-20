package main

import (
	"log"
	"main/db"
	"main/req"
	"main/server"
)

func main() {
	db := new(db.DB)
	err := db.Init()
	if err != nil {
		return
	}

	server.Run("8081", db)

	request := new(req.Req)
	request.Init()

	// autokatalog.Parse(db, func() autokatalog.IReq { return req.New() })
	// car_complaints.ParseCarComplaints(db)
	// automaniac.Parse(db, func() automaniac.IReq { return req.New() })
	// motorreviewer.Parse(db, func() motorreviewer.IReq { return req.New() })
	// nhtsa.Parse(db)
	// carproblemzoo.Parse(db, func() carproblemzoo.IReq { return req.New() })
	log.Println("Done.")
	db.Close()
}
