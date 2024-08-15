package config

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/util"
	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	GinMode       string `mapstructure:"GIN_MODE"`
	MysqlDatabase string `mapstructure:"MYSQL_DATABASE"`
	MysqlUser     string `mapstructure:"MYSQL_USER"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MysqlAddr     string `mapstructure:"MYSQL_ADDR"`
}

var AppConfig *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	util.FailOnError(v.BindEnv("PORT"), "Failed on Bind PORT")
	util.FailOnError(v.BindEnv("GIN_MODE"), "Failed on Bind GIN_MODE")
	util.FailOnError(v.BindEnv("MYSQL_DATABASE"), "fail on Bind MYSQL_DATABASE")
	util.FailOnError(v.BindEnv("MYSQL_USER"), "fail on Bind MYSQL_USER")
	util.FailOnError(v.BindEnv("MYSQL_PASSWORD"), "fail on Bind MYSQL_PASSWORD")
	util.FailOnError(v.BindEnv("MYSQL_ADDR"), "fail on Bind MYSQL_ADDR")
	// util.FailOnError(v.BindEnv("DB_URL"), "Failed on Bind DB_URL")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("Load from environment variable")
	}
	err = v.Unmarshal(&AppConfig)
	if err != nil {
		util.FailOnError(err, "Failed to read enivronment")
	}
}
