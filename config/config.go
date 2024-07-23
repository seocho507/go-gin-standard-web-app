package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Database struct {
		Dialect   string
		Host      string
		Port      string
		Dbname    string
		Username  string
		Password  string
		Migration bool
	}

	ServerInfo struct {
		Port string
	}
}

func NewConfig(path string, log *logrus.Logger) *Config {
	cfg := new(Config)
	_, err := toml.DecodeFile(path, cfg)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatalln("Failed to load config file")
		panic(err)
	}
	return cfg
}
