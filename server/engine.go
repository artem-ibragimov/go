package server

import (
	"encoding/json"
	DB "main/db"
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
func CreateEngineGetter(db IDB) func(ctx *gin.Context) {
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

func CreateEnginePoster(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		jsonData, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		var engine *DB.EngineData
		err = json.Unmarshal(jsonData, &engine)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		engineID, err := db.PostEngine(engine)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, engineID)
	}
}

func CreateEnginePatcher(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		engineID, err := strconv.Atoi(ctx.Param("engineID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		jsonData, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		var engine *DB.EngineData
		err = json.Unmarshal(jsonData, &engine)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		_, err = db.PatchEngine(int32(engineID), engine)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	}
}
