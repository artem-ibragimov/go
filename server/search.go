package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

const RESULT_LIMIT = 6

func CreateSearchGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		query, ok := ctx.GetQuery("query")
		if !ok {
			ctx.JSON(http.StatusServiceUnavailable, errors.New("wrong query"))
			return
		}

		brand_query := query
		model_query := query
		gen_query := query
		words := strings.Split(query, " ")

		if len(words) > 1 {
			brand_query = words[0]
			model_query = words[1]
		}

		if len(words) > 2 {
			gen_query = words[2]
		}

		brand_id := "0"
		brands, err := db.SearchBrands(brand_query, RESULT_LIMIT)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		if len(brands) != 0 {
			brand_id = maps.Keys(brands)[0]
		}
		extra := uint(RESULT_LIMIT - len(brands))
		models, err := db.SearchModels(brand_id, model_query, RESULT_LIMIT+extra)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		model_id := "0"
		extra = uint(RESULT_LIMIT - len(models))
		if len(models) != 0 {
			model_id = maps.Keys(models)[0]
		}
		generations, err := db.SearchGenerations(model_id, gen_query, RESULT_LIMIT+extra)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}

		var data map[string]map[string]string = make(map[string]map[string]string, 3)
		data["gens"] = generations
		data["models"] = models
		data["brands"] = brands
		// TODO если одна модель возвращать все поколения
		ctx.JSON(http.StatusOK, data)
	}
}
