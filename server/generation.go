package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateGenerationsListGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		model_id, err := strconv.Atoi(ctx.Query("model_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Parameter model_id is missing")
			return
		}
		data, err := db.GetGenerations(int32(model_id))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
