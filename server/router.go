package server

import (
	"github.com/gin-gonic/gin"
	swagFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"

	"go-app-api-template/docs"
)

func setupRouter() *gin.Engine {
	r := getGinEngine()

	api := r.Group("/api")
	{
		// Get api health
		api.GET("/ping", Ping)
		v1 := api.Group("/v1")
		{
			// Ping test
			v1.GET("/ping", Ping)

			// Say Hello !
			v1.GET("/hello", SayHelloWorld())
		}
	}

	// configure swagger
	docs.SwaggerInfo.BasePath = "/"
	api.GET("/docs/*any", ginSwag.WrapHandler(swagFiles.Handler))

	return r
}
