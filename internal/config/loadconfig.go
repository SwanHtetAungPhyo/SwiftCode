package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

type AppConfiguration struct {
	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	DbName   string
	PORT     string
	SSLMODE  string
	FilePath string
	Mode     string
}

func LoadConfig() *AppConfiguration {
	dbHost := getEnv("DB_HOST", "localhost") // 'localhost' if running locally, 'postgres' inside Docker
	dbPort := getEnv("DB_PORT", "5432")
	dbUsername := getEnv("DB_USER", "swanhtetaungphyo")
	dbPassword := getEnv("DB_PASSWORD", "swanhtet12")
	dbName := getEnv("DB_NAME", "app_db")
	port := getEnv("PORT", "8080")
	sslMode := getEnv("SSLMODE", "disable")
	filePath := getEnv("FILE_PATH", "/Users/swanhtet1aungphyo/IdeaProjects/SwiftCode/data/swif_codes.csv")
	mode := getEnv("MODE", "development")
	return &AppConfiguration{
		DbHost:   dbHost,
		DbPort:   dbPort,
		DbUser:   dbUsername,
		DbPass:   dbPassword,
		DbName:   dbName,
		PORT:     port,
		SSLMODE:  sslMode,
		FilePath: filePath,
		Mode:     mode,
	}
}
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		logrus.Infof("Environment variable %s not set", key)
		return fallback
	}
	return value
}
