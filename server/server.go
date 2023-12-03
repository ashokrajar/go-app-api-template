/*
MIT License

# Copyright © Ashok Raja <ashokrajar@gmail.com>

Authors: Ashok Raja <ashokrajar@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package server
package server

import (
	"fmt"

	log "github.com/ashokrajar/zerolog_wrapper"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"go-app-api-template/config"
)

func getGinEngine() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	if config.AppEnv == "prod" {
		r = gin.New()
		r.ForwardedByClientIP = true
		// r.SetTrustedProxies([]string{"127.0.0.1"})
		r.Use(logger.SetLogger(
			logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
				return l.Output(gin.DefaultWriter).With().Logger()
			}),
		))
	}

	return r
}

func StartServer() {
	log.Info().Msg("Starting API Server .....")

	r := setupRouter()
	// Start server and listen on 0.0.0.0:$(AppSvcPort)
	if err := r.Run(fmt.Sprintf(":%s", config.AppSvcPort)); err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("can't start server on port : %s", config.AppSvcPort))
	}

	log.Info().Msg("API Server start-up succeeded .....")
}
