package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateModelsListGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		brand_id, err := strconv.Atoi(ctx.Query("brand_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter brand_id is missing")
			return
		}
		data, err := db.GetModelsByBrand(int32(brand_id))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
