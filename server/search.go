package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const RESULT_LIMIT = 3

func CreateSearchGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		query, ok := ctx.GetQuery("query")
		if !ok {
			ctx.JSON(http.StatusServiceUnavailable, errors.New("wrong query"))
			return
		}
		brands, err := db.SearchBrands(query, RESULT_LIMIT)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		extra := uint(RESULT_LIMIT - len(brands))
		models, err := db.SearchModels(query, RESULT_LIMIT+extra)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		extra = uint(RESULT_LIMIT - len(models))
		generations, err := db.SearchGenerations(query, RESULT_LIMIT+extra)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}

		var data map[string]map[string]string = make(map[string]map[string]string, 3)
		data["brands"] = brands
		data["models"] = models
		data["generations"] = generations
		// TODO если одна модель возвращать все поколения
		ctx.JSON(http.StatusOK, data)
	}
}
