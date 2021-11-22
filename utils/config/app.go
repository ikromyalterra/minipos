package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	//driver
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// set config based on env
	LoadEnvVars()
	OpenDbPool()
}

func LoadEnvVars() {
	// Bind OS environment variable
	viper.SetEnvPrefix("minipos")
	viper.BindEnv("env")
	viper.BindEnv("app_path") //

	dir, _ := os.Getwd()
	AppPath := dir

	cfg := "config"
	var env string
	if viper.Get("env") != nil {
		env = viper.Get("env").(string)
	}
	if strings.HasPrefix(env, "dev") {
		cfg += "_dev"
	} else if strings.HasPrefix(env, "test") {
		cfg += "_test"
		if viper.Get("app_path") != nil {
			AppPath = viper.Get("app_path").(string)
		}
	}
	viper.SetConfigName(cfg)
	viper.SetConfigType("json")
	viper.AddConfigPath(AppPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
