/*
Package repo provides functionality for communicating between the application and the database.

The `Init` function initializes the database connection, with retry logic, ensuring that the database connection is set up once.

Functions:
  - `Init`: Initializes the database connection and sets up the connection pooling with retry logic.
  - `GetDBInstance`: Retrieves the singleton instance of the database connection.

Example usage:

	repo.Init()
	dbInstance := repo.GetDBInstance()
*/
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
	// Avoid for the global db instance
	instance *gorm.DB
	// Ensure the db connection is initialized once
	once sync.Once
)

// Init initializes the database connection using the provided configuration details.
// It includes a retry mechanism, attempting to connect to the database up to 5 times with a 5-second delay between attempts.
// If the connection is successful, connection pooling parameters are set.
func Init(log *logrus.Logger, cfg *config.AppConfiguration) {
	once.Do(func() {
		log.Info("Initializing database connection...")
		dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
			cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbUser, cfg.DbPass, cfg.SSLMODE)
		var err error
		// Retry logic for establishing the connection
		for i := 0; i < 5; i++ {
			instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
			log.Warnf("Connection attempt %d failed: %v", i+1, err)
			time.Sleep(5 * time.Second)
		}

		// If the connection is not established after retries, log a fatal error
		if err != nil {
			log.Fatal("Database connection failed after retries: ", err)
		}

		// Set up the SQL DB instance and connection pooling
		sqlDB, err := instance.DB()
		if err != nil {
			log.Error("Failed to retrieve SQL DB instance: ", err)
			return
		}
		// Connection Pooling Set up
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
		log.Info("Database connection initialized successfully.")
	})
}

// GetDBInstance retrieves the singleton instance of the database connection.
// This allows the application to perform database operations by obtaining a reference to the initialized instance.
func GetDBInstance() *gorm.DB {
	return instance
}
