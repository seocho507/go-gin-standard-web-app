package service

import (
	"github.com/seocho507/go-gin-standard-web-app/entity"
	"github.com/seocho507/go-gin-standard-web-app/repository"
	"github.com/seocho507/go-gin-standard-web-app/util"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	SaveUser(*entity.User) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	log      *logrus.Logger
}

func NewUserService(userRepo repository.UserRepository, log *logrus.Logger) UserService {
	return &userService{
		userRepo: userRepo,
		log:      log,
	}
}

func (s *userService) SaveUser(user *entity.User) (*entity.User, error) {
	hashedPassword, _ := util.HashPassword(user.Password)
	user.Password = hashedPassword
	return s.userRepo.CreateUser(user)
}
