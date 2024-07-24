package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/seocho507/go-gin-standard-web-app/config"
	"github.com/sirupsen/logrus"
)

type Router struct {
	cfg *config.Config
	log *logrus.Logger

	Engine *gin.Engine
}

func InitRouter(cfg *config.Config, log *logrus.Logger) *Router {
	r := &Router{
		cfg:    cfg,
		log:    log,
		Engine: gin.New(),
	}

	r.Engine.Use(gin.Logger())

	// TODO : implement error handler
	r.Engine.Use(gin.Recovery())

	r.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		ExposeHeaders:    []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	return r
}

func (r *Router) Run() error {
	return r.Engine.Run(r.cfg.ServerInfo.Port)
}
