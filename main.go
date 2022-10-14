package main

import (
	"github.com/gin-gonic/gin"
	"github.com/triadmoko/edot-interview/configs"
	"github.com/triadmoko/edot-interview/handlers"
	"github.com/triadmoko/edot-interview/helpers"
	"github.com/triadmoko/edot-interview/repositorys"
	"github.com/triadmoko/edot-interview/services"
)

func main() {
	log := helpers.NewLogger("./tmp/logs.log", true)
	db, err := configs.DB()
	if err != nil {
		log.Error(err.Error())
	}
	// db.AutoMigrate(&models.Product{})
	// db.AutoMigrate(&models.Category{})
	repository := repositorys.NewRepository(db)
	service := services.NewService(repository, log)
	handler := handlers.NewHandler(service)

	router := gin.Default()

	product := router.Group("/product")
	product.POST("/", handler.CreateProduct)
	product.GET("/", handler.GetProducts)
	product.GET("/:id", handler.GetProductByID)
	product.PUT("/:id", handler.UpdateProduct)
	product.PATCH("/:id", handler.UpdateProduct)
	product.DELETE("/:id", handler.DeleteProduct)

	category := router.Group("/category")
	category.POST("/")
	category.GET("/")
	category.GET("/:id")
	category.PUT("/:id")
	category.PATCH("/:id")
	category.DELETE("/:id")

	router.Run()
}
