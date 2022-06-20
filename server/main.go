package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IDB interface {
	GetBrands() (map[string]string, error)
	GetModelsByBrand(brand_id int32) (map[string]string, error)
	GetGenerations(model_id int32) (map[string]string, error)
	GetVersions(generation_id int32) (map[string]string, error)
}

func Run(port string, db IDB) {
	router := gin.Default()
	// router.Use(allowCORS)
	data_group := router.Group("/data")
	{
		data_group.GET("/brand/", CreateBrandsListGetter(db))
		data_group.GET("/model/", CreateModelsListGetter(db))
		data_group.GET("/generation/", CreateGenerationsListGetter(db))
		data_group.GET("/version/", CreateVersionsListGetter(db))
	}
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static")
	})
	router.Static("/static", "./server/static/index.html")
	router.Run(":" + port)
}

// func allowCORS(ctx *gin.Context) {
// 	ctx.Header("Access-Control-Allow-Origin", "*")
// 	ctx.Next()
// }
