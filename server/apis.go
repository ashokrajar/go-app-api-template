package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SayHelloWorld @BasePath /api/v1/hello
//
// @version 1.0
// @Summary Say Hello World !
// @Tags API V1
// @Produce json
// @schemes http https
// @Success 200 {string} Hello world!
// @Router /api/v1/hello [get]
func SayHelloWorld() gin.HandlerFunc {
	return func(g *gin.Context) {
		g.String(http.StatusOK, "Hello World !!!")
	}
}
