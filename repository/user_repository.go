package repository

import (
	"github.com/seocho507/go-gin-standard-web-app/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	FindAllUser() ([]entity.User, error)
}

type userRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewUserRepository(db *gorm.DB, log *logrus.Logger) UserRepository {
	repo := &userRepository{
		db:  db,
		log: log,
	}

	err := repo.db.AutoMigrate(&entity.User{})

	if err != nil {
		log.WithError(err).Fatal("Failed to create user table")
		panic(err)
	}

	return repo
}

func (r *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		r.log.WithError(err).Error("Failed to create user")
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindAllUser() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	if err != nil {
		r.log.WithError(err).Error("Failed to find all users")
		return nil, err
	}

	return users, nil
}
