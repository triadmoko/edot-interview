package test

import (
	"github.com/gin-gonic/gin"
	"github.com/triadmoko/edot-interview/configs"
	"github.com/triadmoko/edot-interview/handlers"
	"github.com/triadmoko/edot-interview/helpers"
	"github.com/triadmoko/edot-interview/models"
	"github.com/triadmoko/edot-interview/repositorys"
	"github.com/triadmoko/edot-interview/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Testing interface {
	SetupRouter() *gin.Engine
	SetupMongo()
}
type RouterTesting struct {
	testing Testing
}

func NewGinConfig() *gin.Engine {
	log := helpers.NewLogger("./../tmp/logs.log", true)
	db, err := configs.DB()
	if err != nil {
		log.Error(err.Error())
	}
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Category{})
	repository := repositorys.NewRepository(db)
	service := services.NewService(repository, log)
	handler := handlers.NewHandler(service)
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	product := router.Group("/product")
	product.POST("/", handler.CreateProduct)
	product.GET("/", handler.GetProducts)
	product.GET("/:id", handler.GetProductByID)
	product.PUT("/:id", handler.UpdateProduct)
	product.PATCH("/:id", handler.UpdateProduct)
	product.DELETE("/:id", handler.DeleteProduct)

	category := router.Group("/category")
	category.POST("/", handler.CreateCategory)
	category.GET("/", handler.GetCategorys)
	category.GET("/:id", handler.GetCategoryByID)
	category.PUT("/:id", handler.UpdateCategory)
	category.PATCH("/:id", handler.UpdateCategory)
	category.DELETE("/:id", handler.DeleteCategory)

	return router

}
func DB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
