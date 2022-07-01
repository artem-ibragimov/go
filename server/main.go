package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IDB interface {
	GetBrands() (map[string]string, error)
	SearchBrands(query string, limit uint) (map[string]string, error)

	GetModelsByBrand(brandID int32) (map[string]string, error)
	SearchModels(query string, limit uint) (map[string]string, error)

	GetGeneration(genID int32) (map[string]string, error)
	GetGenerations(modelID int32) (map[string]string, error)
	SearchGenerations(query string, limit uint) (map[string]string, error)

	GetVersions(generationID int32) (map[string]string, error)
	GetVersion(versionID int32) (map[string]string, error)

	GetEngines() (map[string]string, error)
	GetEngine(engineID int32) (map[string]string, error)
}

func Run(port string, db IDB) {
	router := gin.Default()
	data_group := router.Group("/data")
	{
		data_group.GET("/search/", CreateSearchGetter(db))

		data_group.GET("/brand/", CreateBrandsListGetter(db))

		data_group.GET("/model/", CreateModelsListGetter(db))

		data_group.GET("/generation/:genID", CreateGenDescriptionGetter(db))
		data_group.GET("/generation/", CreateGenerationsListGetter(db))

		data_group.GET("/version/:versionID", CreateVersionDescriptionGetter(db))
		data_group.GET("/version/", CreateVersionsListGetter(db))

		data_group.GET("/engine/:engineID", CreateEngineDescriptionGetter(db))
		data_group.GET("/engine/", CreateEngineListGetter(db))
	}
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static")
	})
	router.Static("/static", "./server/static/index.html")
	router.Run(":" + port)
}
