package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTransmissionListGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		brandID, err := strconv.Atoi(ctx.Query("brandID"))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		data, err := db.GetTransmissions(int32(brandID))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
func CreateTransmissionDescriptionGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		transmissionID, err := strconv.Atoi(ctx.Param("transmissionID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter transmissionID is missing")
			return
		}
		data, err := db.GetTransmission(int32(transmissionID))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
