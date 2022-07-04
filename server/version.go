package server

import (
	"encoding/json"
	DB "main/db"
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
func CreateVersionGetter(db IDB) func(ctx *gin.Context) {
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

func CreateVersionPoster(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		jsonData, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		var version *DB.VersionData
		err = json.Unmarshal(jsonData, &version)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		versionID, err := db.PostVersion(version)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, versionID)
	}
}

func CreateVersionPatcher(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		versionID, err := strconv.Atoi(ctx.Param("versionID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter versionID is missing")
			return
		}
		jsonData, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		var version *DB.VersionData
		err = json.Unmarshal(jsonData, &version)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		_, err = db.PatchVersion(int32(versionID), version)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, versionID)
	}
}
