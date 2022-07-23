package main

import (
	"log"
	"main/autokatalog"
	"main/automaniac"
	"main/car_complaints"
	"main/carproblemzoo"
	"main/db"
	"main/drom"
	"main/motorreviewer"
	"main/nhtsa"
	"main/req"
	"main/server"
	"os"
)

func main() {
	db := new(db.DB)
	err := db.Init()
	request := new(req.Req)
	request.Init()
	if err != nil {
		return
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("SERVER") == "1" {
		server.Run("8081", db)
	}

	if os.Getenv("DROM") == "1" {
		drom.Parse(db, func() drom.IReq { return req.New() })
	}
	if os.Getenv("AUTOKATALOG") == "1" {
		autokatalog.Parse(db, func() autokatalog.IReq { return req.New() })
	}
	if os.Getenv("CARPROBLEMZOO") == "1" {
		carproblemzoo.Parse(db, func() carproblemzoo.IReq { return req.New() })
	}
	if os.Getenv("CARCOMPLAINTS") == "1" {
		car_complaints.Parse(db)
	}
	if os.Getenv("AUTOMANIAC") == "1" {
		automaniac.Parse(db, func() automaniac.IReq { return req.New() })
	}
	if os.Getenv("MOTORREVIEVER") == "1" {
		motorreviewer.Parse(db, func() motorreviewer.IReq { return req.New() })
	}
	if os.Getenv("NHTSA") == "1" {
		nhtsa.Parse(db)
	}
	log.Println("Done.")
	db.Close()
}
