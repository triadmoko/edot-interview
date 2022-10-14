package main

import (
	"net/http"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	_ "github.com/triadmoko/edot-interview/docs"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/triadmoko/edot-interview/configs"
	"github.com/triadmoko/edot-interview/handlers"
	"github.com/triadmoko/edot-interview/helpers"
	"github.com/triadmoko/edot-interview/repositorys"
	"github.com/triadmoko/edot-interview/services"
)

// @title           Backend Golang Coding Interview
// @version         1.0
// @description     Documentation APIs triadmoko12@gmail.com for client.

// @contact.name   Triadmoko Support
// @contact.email  triadmoko@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
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

	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	router.Use(mw)
	product := router.Group("/product")
	product.POST("/", handler.CreateProduct)
	product.GET("/", handler.GetProducts)
	product.GET("/:id", handler.GetProductByID)
	product.PUT("/:id", handler.UpdateProduct)
	product.PATCH("/:id", handler.UpdateProductPatch)
	product.DELETE("/:id", handler.DeleteProduct)

	category := router.Group("/category")
	category.POST("/", handler.CreateCategory)
	category.GET("/", handler.GetCategorys)
	category.GET("/:id", handler.GetCategoryByID)
	category.PUT("/:id", handler.UpdateCategory)
	category.PATCH("/:id", handler.UpdateCategory)
	category.DELETE("/:id", handler.DeleteCategory)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}
func errorHandler(c *gin.Context, info ratelimit.Info) {
	response := helpers.ResponseApi("Too many requests, Max request 5 per minute. Try again in "+time.Until(info.ResetTime).String(), http.StatusTooManyRequests, "request limit", nil)
	c.JSON(http.StatusTooManyRequests, response)
}
