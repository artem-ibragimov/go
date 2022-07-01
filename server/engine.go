package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEngineListGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data, err := db.GetEngines()
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
func CreateEngineDescriptionGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		engineID, err := strconv.Atoi(ctx.Param("engineID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter engineID is missing")
			return
		}
		data, err := db.GetEngine(int32(engineID))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
