package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	Mode       string `mapstructure:"MODE"`
	Port       int    `mapstructure:"PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
}

var config *Configuration
var once sync.Once

var requiredEnvs []string = []string{"MODE", "PORT", "DB_HOST", "DB_USER",
	"DB_PASSWORD", "DB_NAME"}

//Init

func Init() {
	once.Do(func() {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		//Set default values

		// next
		viper.AutomaticEnv()
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error in reading the config. Err: %s", err.Error())
		}
		//Check required Values
		checkRequiredValues()
		config = new(Configuration)
		err = viper.Unmarshal(config)
		if err != nil {
			log.Fatalf("Error in Unmarshalling the config. Err: %s", err.Error())
		}
	})
}

func checkRequiredValues() {
	for _, k := range requiredEnvs {
		if !viper.IsSet(k) {
			panic(k + ": Required key is not present")
		}
	}
}

// Config ...
func Config() *Configuration {
	return config
}

// SetConfig ...
func SetConfig(c *Configuration) {
	config = c
}
