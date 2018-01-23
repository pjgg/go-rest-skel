package configuration

import (
	"os"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigurationManager struct {
	SentryDsn     string
	Enviroment    string
	Version       string
	ServerAddress string
}

type ConfigurationManagerBehavior interface{}

var once sync.Once
var ConfigurationManagerInstance *ConfigurationManager

func NewConfigurationManager() *ConfigurationManager {

	once.Do(func() {
		//TODO write here all your conf stuff
		sentryDsn := viper.GetString("server.sentryDSN")
		serverAddress := viper.GetString("server.address")
		env, existEnv := os.LookupEnv("ENVIROMENT")
		version, existVersion := os.LookupEnv("VERSION")

		if !existEnv {
			env = "LOCAL"
		}

		if !existVersion {
			version = "0.0.1-snapshot"
		}

		ConfigurationManagerInstance = &ConfigurationManager{
			SentryDsn:     sentryDsn,
			Enviroment:    env,
			Version:       version,
			ServerAddress: serverAddress}

		log.WithFields(log.Fields{
			"SentryDsn":     ConfigurationManagerInstance.SentryDsn,
			"Enviroment":    ConfigurationManagerInstance.Enviroment,
			"Version":       ConfigurationManagerInstance.Version,
			"ServerAddress": ConfigurationManagerInstance.ServerAddress,
		}).Info("configurationManager loaded")
	})

	return ConfigurationManagerInstance
}
