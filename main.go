package main

import (
	"flag"
	"github.com/go-playground/validator/v10"
	"github.com/seocho507/go-gin-standard-web-app/config"
	. "github.com/seocho507/go-gin-standard-web-app/constant"
	"github.com/seocho507/go-gin-standard-web-app/controller"
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

	db, err := repository.ConnectDatabase(cfg, log)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to database")
	}

	v := validator.New()
	userRepo := repository.NewUserRepository(db, log)
	userService := service.NewUserService(userRepo, log)
	r := router.InitRouter(cfg, log)
	controller.NewUserController(userService, log, r, v)
	err = r.Engine.Run(cfg.ServerInfo.Port)
	if err != nil {
		log.WithError(err).Fatal("Failed to start server")
		panic(err)
	}
}
