package server

import (
	DB "main/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IDB interface {
	GetBrandByName(string) (int32, error)
	GetBrands() (map[string]string, error)
	PostBrand(string) (int32, error)
	SearchBrands(query string, limit uint) (map[string]string, error)

	GetModelID(brand_id int32, model_name string) (int32, error)
	PostModel(*DB.ModelData) (int32, error)
	GetModelsByBrand(brandID int32) (map[string]string, error)
	SearchModels(brand_id string, query string, limit uint) (map[string]string, error)

	GetGeneration(genID int32) (*DB.GenerationData, error)
	GetGenerations(modelID int32) (map[string]string, error)
	GetGenID(model_id int32, name string) (int32, error)
	PatchGeneration(id int32, data *DB.GenerationData) (int32, error)
	PostGen(data *DB.GenerationData) (int32, error)
	SearchGenerations(model_id string, query string, limit uint) (map[string]string, error)

	GetVersions(generationID int32) (map[string]string, error)
	GetVersion(versionID int32) (*DB.VersionData, error)
	GetVersionID(name string, gen_id int32) (int32, error)
	PatchVersion(id int32, version *DB.VersionData) (int32, error)
	PostVersion(*DB.VersionData) (int32, error)

	PostEngine(*DB.EngineData) (int32, error)
	PatchEngine(id int32, engine *DB.EngineData) (int32, error)
	GetEngines() (map[string]string, error)
	GetEngine(engineID int32) (*DB.EngineData, error)

	GetTranss(brandID int32) (map[string]string, error)
	GetTrans(id int32) (*DB.TransData, error)
	GetTransID(brand_id int32, name string, gears int32) (int32, error)
	PostTrans(*DB.TransData) (int32, error)
	PatchTrans(id int32, trans *DB.TransData) (int32, error)

	GetDefectsAgesByBrand(brand_id int32, norm bool) (map[string]string, error)
	GetDefectsAgesByModel(model_id int32, norm bool) (map[string]string, error)
	GetDefectsAgesByGen(gen_id int32, norm bool) (map[string]string, error)
}

func Run(port string, db IDB) {
	router := gin.Default()
	data_group := router.Group("/data")
	version := data_group.Group("/version")
	generation := data_group.Group("/generation")
	engine := data_group.Group("/engine")
	transmission := data_group.Group("/transmission")
	defect := data_group.Group("/defect")

	{
		data_group.GET("/search/", CreateSearchGetter(db))

		data_group.GET("/brand/", CreateBrandsListGetter(db))

		data_group.GET("/model/", CreateModelsListGetter(db))
		{
			generation.PATCH("/:genID", CreateGenPatcher(db))
			generation.POST("/", CreateGenPoster(db))
			generation.GET("/:genID", CreateGenGetter(db))
			generation.GET("/", CreateGenListGetter(db))
		}
		{
			version.PATCH("/:versionID", CreateVersionPatcher(db))
			version.POST("/", CreateVersionPoster(db))
			version.GET("/:versionID", CreateVersionGetter(db))
			version.GET("/", CreateVersionsListGetter(db))
		}
		{
			engine.PATCH("/:engineID", CreateEnginePatcher(db))
			engine.POST("/", CreateEnginePoster(db))
			engine.GET("/:engineID", CreateEngineGetter(db))
			engine.GET("/", CreateEngineListGetter(db))
		}
		{
			transmission.PATCH("/:transID", CreateTransPatcher(db))
			transmission.POST("/", CreateTransPoster(db))
			transmission.GET("/:transID", CreateTransGetter(db))
			transmission.GET("/", CreateTransListGetter(db))
		}
		{
			defect.GET("/", CreateDefectsGetter(db))
		}
	}
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static")
	})
	router.Static("/static", "./server/static/index.html")
	router.Run(":" + port)
}
