package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBrandsListGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data, err := db.GetBrands()
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
