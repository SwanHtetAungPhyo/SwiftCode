package repo

import (
	"fmt"
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

var DbInstance *gorm.DB
var err error

func Init() {
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"app_db", port, user, password, dbName)
	DbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logging.Logger.Error(err.Error())
		panic(err)
	}
	sqlDB, err := DbInstance.DB()
	if err != nil {
		logging.Logger.Error(err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	err = ModelMigration(DbInstance)
	if err != nil {
		return
	}
}
