package service

import (
	"github.com/seocho507/go-gin-standard-web-app/config"
	"github.com/seocho507/go-gin-standard-web-app/repository"
	"github.com/sirupsen/logrus"
)

type Service interface {
}

type service struct {
	cfg  *config.Config
	repo repository.Repository
	log  *logrus.Logger
}

func NewService(cfg *config.Config, repo repository.Repository, log *logrus.Logger) Service {
	return &service{
		cfg:  cfg,
		repo: repo,
		log:  log,
	}
}
