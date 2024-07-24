package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/seocho507/go-gin-standard-web-app/entity"
	"github.com/seocho507/go-gin-standard-web-app/router"
	"github.com/seocho507/go-gin-standard-web-app/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserController struct {
	userService service.UserService
	router      *router.Router
	log         *logrus.Logger
	validator   *validator.Validate
}

func NewUserController(userService service.UserService, log *logrus.Logger, router *router.Router, v *validator.Validate) {
	ctl := &UserController{
		userService: userService,
		log:         log,
		router:      router,
		validator:   v,
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

	// TODO : Extract validation logic to middleware or util
	if err := c.validator.Struct(&u); err != nil {
		c.log.WithError(err).Error("Validation failed")
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldError := range validationErrors {
			errorMessages[fieldError.Field()] = fieldError.Error()
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"validation_errors": errorMessages})
		return
	}

	saveUser, err := c.userService.SaveUser(&u)
	if err != nil {
		c.log.WithError(err).Error("Failed to save user")
		return
	}
	ctx.JSON(200, saveUser)
}
