package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"to-do-app/config"
	"to-do-app/logger"
	"to-do-app/models"
)

var DB *gorm.DB

func NewDBConnection(cfg *config.Params) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			cfg.Postgres.Host,
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.DBName,
			cfg.Postgres.Port,
			cfg.Postgres.SSLMode,
			cfg.Postgres.TimeZone)),
		&gorm.Config{})

	if err != nil {
		return nil, err
	}

	logger.Log.Info("Database connection successful")

	err = MigrateModels(db)
	if err != nil {
		return nil, err
	}

	logger.Log.Info("Models migrated successful")

	return db, nil
}

func MigrateModels(psql *gorm.DB) error {
	err := psql.AutoMigrate(&models.Tasks{})
	if err != nil {
		return err
	}

	return nil
}
