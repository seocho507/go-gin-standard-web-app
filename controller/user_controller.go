package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/seocho507/go-gin-standard-web-app/entity"
	"github.com/seocho507/go-gin-standard-web-app/router"
	"github.com/seocho507/go-gin-standard-web-app/service"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	userService service.UserService
	router      *router.Router
	log         *logrus.Logger
}

func NewUserController(userService service.UserService, log *logrus.Logger, router *router.Router) {
	ctl := &UserController{
		userService: userService,
		log:         log,
		router:      router,
	}
	ctl.setRoutes(router, userService, log)
}

func (c *UserController) setRoutes(router *router.Router, userService service.UserService, log *logrus.Logger) {
	c.router = router
	c.userService = userService
	c.log = log

	c.router.Engine.POST("/user", c.SaveUser)
}

func (c *UserController) SaveUser(ctx *gin.Context) {
	var u entity.User
	if err := ctx.ShouldBindBodyWithJSON(&u); err != nil {
		c.log.WithError(err).Error("Failed to bind json")
		ctx.JSON(400, gin.H{"error": "Failed to bind json"})
		return
	}

	saveUser, err := c.userService.SaveUser(&u)
	if err != nil {
		c.log.WithError(err).Error("Failed to save user")
		return
	}
	ctx.JSON(200, saveUser)
}
