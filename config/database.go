package config

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	config := Config
	encodePassword := url.QueryEscape(config.Database.Password)
	url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.Username,
		encodePassword,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConnection)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.Database.MaxIdleTime) * time.Second)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Database.MaxIdleTime) * time.Second)

	return db, nil
}
