package main

import (
	"fmt"

	"github.com/iamolegga/enviper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/vivekweb2013/deck-api/internal/config"
)

var conf config.Config

func main() {
	initConfig()
	fmt.Printf("%+v", conf)
}

func initConfig() {
	v := enviper.New(viper.New())

	v.AddConfigPath(".")
	v.SetConfigName(".app")

	if err := v.Unmarshal(&conf); err != nil {
		logrus.Fatal("error occurred while parsing config file: ", err)
	}

	logrus.Infof("using the config file: %s", v.ConfigFileUsed())
}
