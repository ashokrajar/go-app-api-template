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

// Package cmd
package cmd

import (
	"os"
	"strings"

	log "github.com/ashokrajar/zerolog_wrapper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags
	cfgFile string
	dryRun  bool

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "go-app-api-template",
		Short: "A simple API Applicaion",
		Long: `A simple hello world API Application 

Written in Golang for a GitHub template repository.`,
		// Uncomment the following line if your bare application has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initBuildInfo)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-app-api.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVar(&dryRun, "plan", false, "Pretend to execute... (DRYRUN)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".go-app-api-template" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-app-api-template")
	}

	viper.AutomaticEnv() // read in environment variables that match
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Configure logging and log level
	logLevel := strings.ToLower(viper.GetString("log.level"))
	if logLevel == "" {
		logLevel = "info"
	}
	appEnv := strings.ToLower(viper.GetString("app.env"))
	if appEnv == "" {
		appEnv = "dev"
	}
	log.InitLog(log.LogLevel(logLevel), log.Env(appEnv))

	// If a config file is found, read from it.
	if err := viper.ReadInConfig(); err != nil {
		log.Info().Msg("Unable to read application config file. Using defaults & environment variables.")
	} else {
		log.Info().Str("ConfigFile", viper.ConfigFileUsed()).Msg("Using config file & environment variables")
	}
}
