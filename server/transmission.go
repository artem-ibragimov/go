package server

import (
	"encoding/json"
	DB "main/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTransListGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		brandID, err := strconv.Atoi(ctx.Query("brandID"))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		data, err := db.GetTranss(int32(brandID))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
func CreateTransGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		transID, err := strconv.Atoi(ctx.Param("transID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter transID is missing")
			return
		}
		data, err := db.GetTrans(int32(transID))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}

func CreateTransPoster(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		jsonData, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		var trans *DB.TransData
		err = json.Unmarshal(jsonData, &trans)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		id, err := db.PostTrans(trans)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, id)
	}
}

func CreateTransPatcher(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		transID, err := strconv.Atoi(ctx.Param("transID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		jsonData, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		var trans *DB.TransData
		err = json.Unmarshal(jsonData, &trans)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		_, err = db.PatchTrans(int32(transID), trans)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	}
}
