package main

import (
	"github.com/iamolegga/enviper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/vivekweb2013/deck-api/internal/config"
	"github.com/vivekweb2013/deck-api/internal/httpservice"
)

var conf config.Config

func main() {
	initConfig()
	httpservice.Run(conf)
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
