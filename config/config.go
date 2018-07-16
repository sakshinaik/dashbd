package config

import (
	"github.com/spf13/viper"
)

type config struct {
	MySQL_connection string `yaml:"mysql_connection"`
}

var (
	Conf *config
)

func Setup() {
	var err error
	Conf, err = getConf()
	if err != nil {
		panic(err)
	}
}

func getConf() (*config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	conf := &config{}
	err = viper.Unmarshal(conf)

	if err != nil {
		return nil, err
	}

	return conf, nil
}
