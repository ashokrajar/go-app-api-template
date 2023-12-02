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

// Package config
package config

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	LogLevel   = "info"
	AppEnv     = "dev"
	AppSvcPort = "8080"
)

func init() {
	viper.AutomaticEnv() // read in environment variables that match
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	logLevel := viper.GetString("log.level")
	if logLevel != "" {
		LogLevel = logLevel
	}

	appEnv := viper.GetString("app.env")
	if appEnv != "" {
		AppEnv = appEnv
	}

	appSvcPort := viper.GetString("app.service.port")
	if appSvcPort != "" {
		AppSvcPort = appSvcPort
	}
}
