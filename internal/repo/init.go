package repo

import (
	"fmt"
	"github.com/SwanHtetAungPhyo/swifcode/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	instance *gorm.DB
	// TO ensure the db connection is init once
	once sync.Once
)

func Init(log *logrus.Logger, cfg *config.AppConfiguration) {
	once.Do(func() {
		log.Info("Initializing database connection...")
		dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
			cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbUser, cfg.DbPass, cfg.SSLMODE)
		var err error
		for i := 0; i < 5; i++ {
			instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
			log.Warnf("Connection attempt %d failed: %v", i+1, err)
			time.Sleep(5 * time.Second)
		}

		if err != nil {
			log.Fatal("Database connection failed after retries: ", err)
		}

		sqlDB, err := instance.DB()
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

func GetDBInstance() *gorm.DB {
	return instance
}
