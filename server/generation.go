package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateGenerationsListGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		modelID, err := strconv.Atoi(ctx.Query("modelID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter modelID is missing")
			return
		}
		data, err := db.GetGenerations(int32(modelID))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
func CreateGenDescriptionGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		genID, err := strconv.Atoi(ctx.Param("genID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter genID is missing")
			return
		}
		data, err := db.GetGeneration(int32(genID))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
