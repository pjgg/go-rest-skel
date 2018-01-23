//go:generate swagger generate spec
// Package main Skeleton API.
//
// the purpose of this application is to provide an skeleton application
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package main

import (
	"github.com/pjgg/go-rest-skel/configuration"
	"github.com/pjgg/go-rest-skel/cli"

	"os"

	"github.com/spf13/viper"
)

func main() {
	configuration.NewConfigurationManager()
	startServer()
}

func startServer() {
	cli.Execute()
}

func init() {

	viper.SetConfigName("config")
	configPath, exist := os.LookupEnv("CONFIG_PATH")
	if exist {
		viper.AddConfigPath(configPath)
	}
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
