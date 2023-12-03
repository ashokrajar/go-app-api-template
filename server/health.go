package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping @BasePath /api/ping
//
// @Summary Health check api
// @Tags HealthCheck
// @Produce json
// @Success 200 {string} pong
// @Router /api/ping [get]
func Ping(g *gin.Context) {
	g.String(http.StatusOK, "pong")
}
