package main

import (
	"flag"
	"github.com/seocho507/go-gin-standard-web-app/config"
	. "github.com/seocho507/go-gin-standard-web-app/constant"
	"github.com/seocho507/go-gin-standard-web-app/repository"
	"github.com/seocho507/go-gin-standard-web-app/router"
	"github.com/seocho507/go-gin-standard-web-app/service"
	"github.com/sirupsen/logrus"
	"os"
)

var envFlag = flag.String("env", "dev", "environment (dev|prod)")

func main() {
	flag.Parse()
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.InfoLevel
	log.Out = os.Stdout

	var configPath string
	switch *envFlag {
	case DEV:
		configPath = LOCAL_CONFIG_PATH
	case PROD:
		configPath = PROD_CONFIG_PATH
	default:
		log.WithField("env", *envFlag).Info("Env not found, Loading dev config")
		configPath = LOCAL_CONFIG_PATH
	}

	cfg := config.NewConfig(configPath, log)
	log.WithField("config", cfg).Info("Config file loaded successfully")

	repo := repository.NewRepository(cfg, log)
	serv := service.NewService(cfg, repo, log)
	router.InitRouter(cfg, serv, log)
}
