package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateDefectsGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		brandID, err := strconv.Atoi(ctx.Query("brandID"))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		if brandID != 0 {
			data, err := db.GetDefectsAgesByBrand(int32(brandID))
			if err != nil {
				ctx.JSON(http.StatusServiceUnavailable, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, data)
			return
		}
	}
}
