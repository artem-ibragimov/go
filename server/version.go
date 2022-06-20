package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateVersionsListGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		generation_id, err := strconv.Atoi(ctx.Query("generation_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter generation_id is missing")
			return
		}
		data, err := db.GetVersions(int32(generation_id))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
