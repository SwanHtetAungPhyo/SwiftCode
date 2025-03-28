/*
Package config provides functionality for loading and managing application configurations.

The LoadConfig function loads application configuration from environment variables, with fallback values if the environment variables are not set.

The AppConfiguration struct holds the necessary configuration details for connecting to the database and setting up other application parameters.

Functions:
  - LoadConfig: Loads configuration from environment variables with fallback values.
  - getEnv: Retrieves the value of an environment variable or returns a fallback value if not set.

Example usage:

	config := config.LoadConfig()
	fmt.Println(config.DbHost)  // Outputs the DB host (either from environment variable or fallback)
*/
package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

// AppConfiguration holds the configuration parameters for the application
type AppConfiguration struct {
	DbHost   string // Database host
	DbPort   string // Database port
	DbUser   string // Database username
	DbPass   string // Database password
	DbName   string // Database name
	PORT     string // Application port
	SSLMODE  string // SSL mode for database connection
	FilePath string // File path for storing data
	Mode     string // Application mode (development/production)
}

// LoadConfig loads the application configuration from environment variables.
// If an environment variable is not set, it falls back to a default value.
func LoadConfig() *AppConfiguration {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUsername := getEnv("DB_USER", "swanhtetaungphyo")
	dbPassword := getEnv("DB_PASSWORD", "swanhtet12")
	dbName := getEnv("DB_NAME", "app_db")
	port := getEnv("PORT", "8080")
	sslMode := getEnv("SSLMODE", "disable")
	filePath := getEnv("FILE_PATH", "/Users/swanhtet1aungphyo/IdeaProjects/SwiftCode/data/swift_codes.csv")
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

// getEnv retrieves the value of an environment variable or returns a fallback value if not set.
// Logs an info message if the environment variable is not set.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		logrus.Infof("Environment variable %s not set", key)
		return fallback
	}
	return value
}
