package repo

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	DbInstance *gorm.DB
	once       sync.Once
)

func Init(log *logrus.Logger) {
	once.Do(func() {
		log.Info("Initializing database connection...")

		dsn := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable",
			"localhost", "5432", "app_db")

		var err error
		DbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Database connection failed: ", err)
			return
		}

		sqlDB, err := DbInstance.DB()
		if err != nil {
			log.Error("Failed to retrieve SQL DB instance: ", err)
			return
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)

		log.Info("Database connection initialized successfully.")
	})
}
