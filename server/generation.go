package server

import (
	DB "main/db"

	"encoding/json"
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
func CreateGenDescriptionSetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		genID, err := strconv.Atoi(ctx.Param("genID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		println(genID)
		jsonData, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		var gen *DB.GenerationData
		err = json.Unmarshal(jsonData, &gen)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if genID == 0 {
			_, err = db.PostGeneration(gen)
		} else {
			_, err = db.PatchGeneration(int32(genID), gen)
		}
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	}
}
