package repository

import (
	"github.com/seocho507/go-gin-standard-web-app/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	config *config.Config
	log    *logrus.Logger
	db     *gorm.DB
}

func NewRepository(config *config.Config, log *logrus.Logger) Repository {
	r := &repository{
		config: config,
		log:    log,
	}

	var err error
	if r.db, err = connectDatabase(config, log); err != nil {
		panic(err)
	}

	return r
}

func connectDatabase(cfg *config.Config, log *logrus.Logger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch cfg.Database.Dialect {
	case "sqlite3":
		db, err = gorm.Open(sqlite.Open(cfg.Database.Host), &gorm.Config{})
	case "postgres":
		// TODO : implement postgres connection
	default:
		log.Fatalf("Unsupported database dialect: %s", cfg.Database.Dialect)
	}

	if err != nil {
		log.WithError(err).Fatal("Failed to connect to database")
	}

	log.WithField("dialect", cfg.Database.Dialect).Info("Database connection established")
	return db, err
}
