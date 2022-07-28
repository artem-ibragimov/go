package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateDefectsCategoriesGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data, err := db.GetDefectCategories()
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}

func CreateDefectsAgeGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		norm := ctx.Query("norm") == "true"
		categories := strings.Split(ctx.Query("categories"), ",")
		brandID := ctx.Query("brandID")
		if brandID != "" {
			brandID, err := strconv.Atoi(brandID)
			if err != nil {
				ctx.JSON(http.StatusServiceUnavailable, err.Error())
				return
			}
			if brandID != 0 {
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				data, err := db.GetDefectsAgesByBrand(int32(brandID), categories, norm)
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				ctx.JSON(http.StatusOK, data)
				return
			}
		}

		modelID := ctx.Query("modelID")
		if modelID != "" {
			modelID, err := strconv.Atoi(modelID)
			if err != nil {
				ctx.JSON(http.StatusServiceUnavailable, err.Error())
				return
			}
			if modelID != 0 {
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				data, err := db.GetDefectsAgesByModel(int32(modelID), norm)
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				ctx.JSON(http.StatusOK, data)
				return
			}
		}

		genID := ctx.Query("genID")
		if genID != "" {
			genID, err := strconv.Atoi(genID)
			if err != nil {
				ctx.JSON(http.StatusServiceUnavailable, err.Error())
				return
			}
			if genID != 0 {
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				data, err := db.GetDefectsAgesByGen(int32(genID), norm)
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				ctx.JSON(http.StatusOK, data)
				return
			}
		}
	}
}
func CreateDefectsMileageGetter(db IDB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		norm := ctx.Query("norm") == "true"
		brandID := ctx.Query("brandID")
		if brandID != "" {
			brandID, err := strconv.Atoi(brandID)
			if err != nil {
				ctx.JSON(http.StatusServiceUnavailable, err.Error())
				return
			}
			if brandID != 0 {
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				data, err := db.GetDefectsMileageByBrand(int32(brandID), norm)
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				ctx.JSON(http.StatusOK, data)
				return
			}
		}

		modelID := ctx.Query("modelID")
		if modelID != "" {
			modelID, err := strconv.Atoi(modelID)
			if err != nil {
				ctx.JSON(http.StatusServiceUnavailable, err.Error())
				return
			}
			if modelID != 0 {
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				// TODO
				ctx.JSON(http.StatusServiceUnavailable, "not implemented")
				return
				data, err := db.GetDefectsAgesByModel(int32(modelID), norm)
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				ctx.JSON(http.StatusOK, data)
				return
			}
		}

		genID := ctx.Query("genID")
		if genID != "" {
			genID, err := strconv.Atoi(genID)
			if err != nil {
				ctx.JSON(http.StatusServiceUnavailable, err.Error())
				return
			}
			if genID != 0 {
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				// TODO
				ctx.JSON(http.StatusServiceUnavailable, "not implemented")
				return
				data, err := db.GetDefectsAgesByGen(int32(genID), norm)
				if err != nil {
					ctx.JSON(http.StatusServiceUnavailable, err.Error())
					return
				}
				ctx.JSON(http.StatusOK, data)
				return
			}
		}
	}
}
