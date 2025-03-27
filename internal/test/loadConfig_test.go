package test

import (
	config2 "github.com/SwanHtetAungPhyo/swifcode/internal/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func failOnError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Test setup failed: %v", err.Error())
	}
}

// setUpEnv sets environment variables for testing
func setUpEnv() {
	failOnError(nil, os.Setenv("DB_HOST", "test_host"))
	failOnError(nil, os.Setenv("DB_PORT", "1234"))
	failOnError(nil, os.Setenv("DB_USER", "test_user"))
	failOnError(nil, os.Setenv("DB_PASSWORD", "test_password"))
	failOnError(nil, os.Setenv("DB_NAME", "test_db"))
	failOnError(nil, os.Setenv("PORT", "9999"))
	failOnError(nil, os.Setenv("SSLMODE", "require"))
	failOnError(nil, os.Setenv("FILE_PATH", "/tmp/test.csv"))
}

func cleanUpEnv() {
	failOnError(nil, os.Unsetenv("DB_HOST"))
	failOnError(nil, os.Unsetenv("DB_PORT"))
	failOnError(nil, os.Unsetenv("DB_USER"))
	failOnError(nil, os.Unsetenv("DB_PASSWORD"))
	failOnError(nil, os.Unsetenv("DB_NAME"))
	failOnError(nil, os.Unsetenv("PORT"))
	failOnError(nil, os.Unsetenv("SSLMODE"))
	failOnError(nil, os.Unsetenv("FILE_PATH"))
}

func TestLoadConfig_WithEnvVars(t *testing.T) {
	setUpEnv()
	defer cleanUpEnv()

	config := config2.LoadConfig()

	assert.Equal(t, "test_host", config.DbHost)
	assert.Equal(t, "1234", config.DbPort)
	assert.Equal(t, "test_user", config.DbUser)
	assert.Equal(t, "test_password", config.DbPass)
	assert.Equal(t, "test_db", config.DbName)
	assert.Equal(t, "9999", config.PORT)
	assert.Equal(t, "require", config.SSLMODE)
	assert.Equal(t, "/tmp/test.csv", config.FilePath)
}
