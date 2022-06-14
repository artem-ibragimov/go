package server

import (
	"github.com/gin-gonic/gin"
)

type IDB interface {
	GetBrands() ([]string, error)
}

const DATA_ROUTE = "/data"
const BRAND_DATA_ROUTE = DATA_ROUTE + "/brand/*	"

func Run(port string, db IDB) {
	router := gin.Default()
	router.Static("/", "./server/static")
	router.Run(":" + port)
}
