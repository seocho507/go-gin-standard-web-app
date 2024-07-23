package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/seocho507/go-gin-standard-web-app/config"
	"github.com/seocho507/go-gin-standard-web-app/service"
	"github.com/sirupsen/logrus"
)

type router struct {
	cfg     *config.Config
	service service.Service
	log     *logrus.Logger

	engine *gin.Engine
}

func InitRouter(cfg *config.Config, service service.Service, log *logrus.Logger) {
	r := &router{
		cfg:     cfg,
		service: service,
		log:     log,
		engine:  gin.New(),
	}

	r.engine.Use(gin.Logger())

	// TODO : implement error handler
	r.engine.Use(gin.Recovery())

	r.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		ExposeHeaders:    []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	// TODO : Set Controllers

	err := r.engine.Run(cfg.ServerInfo.Port)
	if err != nil {
		log.WithField("error", err).Error("Failed to run server")
		panic(err)
	}
}
