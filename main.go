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
	log := helpers.NewLogger()
	db, err := configs.DB()
	if err != nil {
		log.Error(err.Error())
	}
	repository := repositorys.NewRepository(db)
	service := services.NewService(repository, log)
	handler := handlers.NewHandler(service)

	router := gin.Default()

	product := router.Group("/product")
	product.GET("/")
	product.GET("/:id")
	product.PUT("/:id")
	product.PATCH("/:id")
	product.DELETE("/:id")

	category := router.Group("/category")
	category.GET("/")
	category.GET("/:id")
	category.PUT("/:id")
	category.PATCH("/:id")
	category.DELETE("/:id")

	router.Run()
}
