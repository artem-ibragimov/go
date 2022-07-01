package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateVersionsListGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		generationID, err := strconv.Atoi(ctx.Query("genID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter genID is missing")
			return
		}
		data, err := db.GetVersions(int32(generationID))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
func CreateVersionDescriptionGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		versionID, err := strconv.Atoi(ctx.Param("versionID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter versionID is missing")
			return
		}
		data, err := db.GetVersion(int32(versionID))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
