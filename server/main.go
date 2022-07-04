package server

import (
	DB "main/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IDB interface {
	GetBrandByName(string) (int32, error)
	GetBrands() (map[string]string, error)
	SaveBrand(string) (int32, error)
	SearchBrands(query string, limit uint) (map[string]string, error)

	GetModelID(brand_id int32, model_name string) (int32, error)
	SaveModel(*DB.ModelData) (int32, error)
	GetModelsByBrand(brandID int32) (map[string]string, error)
	SearchModels(query string, limit uint) (map[string]string, error)

	GetGeneration(genID int32) (*DB.GenerationData, error)
	GetGenerations(modelID int32) (map[string]string, error)
	GetGenerationID(model_id int32, name string) (int32, error)
	PatchGeneration(id int32, data *DB.GenerationData) (int32, error)
	PostGeneration(data *DB.GenerationData) (int32, error)
	SearchGenerations(query string, limit uint) (map[string]string, error)

	GetVersions(generationID int32) (map[string]string, error)
	GetVersion(versionID int32) (*DB.VersionData, error)
	GetVersionID(name string, generation_id int32) (int32, error)
	SaveVersion(*DB.VersionData) (int32, error)

	GetEngineID(name string) (int32, error)
	SaveEngine(*DB.EngineData) (int32, error)
	GetEngines() (map[string]string, error)
	GetEngine(engineID int32) (*DB.EngineData, error)

	GetTransmissions(brandID int32) (map[string]string, error)
	GetTransmission(id int32) (*DB.TransmissionData, error)
	GetTransmissionID(brand_id int32, name string, gears int32) (int32, error)
	SaveTransmission(*DB.TransmissionData) (int32, error)
}

func Run(port string, db IDB) {
	router := gin.Default()
	data_group := router.Group("/data")
	{
		data_group.GET("/search/", CreateSearchGetter(db))

		data_group.GET("/brand/", CreateBrandsListGetter(db))

		data_group.GET("/model/", CreateModelsListGetter(db))

		data_group.POST("/generation/:genID", CreateGenDescriptionSetter(db))
		data_group.GET("/generation/:genID", CreateGenDescriptionGetter(db))
		data_group.GET("/generation/", CreateGenerationsListGetter(db))

		data_group.GET("/version/:versionID", CreateVersionDescriptionGetter(db))
		data_group.GET("/version/", CreateVersionsListGetter(db))

		data_group.GET("/engine/:engineID", CreateEngineDescriptionGetter(db))
		data_group.GET("/engine/", CreateEngineListGetter(db))

		data_group.GET("/transmission/:transmissionID", CreateTransmissionDescriptionGetter(db))
		data_group.GET("/transmission/", CreateTransmissionListGetter(db))
	}
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static")
	})
	router.Static("/static", "./server/static/index.html")
	router.Run(":" + port)
}
