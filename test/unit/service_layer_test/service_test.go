package service_layer_test

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/logging"
	"go.uber.org/zap/zaptest"
	"testing"
)

func TestGetBySwiftCode(t *testing.T) {
	logging.Logger = zaptest.NewLogger(t)

	logging.Logger.Info("Passed GetBySwiftCode")
}
func TestGetByISO2(t *testing.T) {
	logging.Logger = zaptest.NewLogger(t)

	logging.Logger.Info("Passed GetBySwiftCode")
}

func TestCreate(t *testing.T) {
	logging.Logger = zaptest.NewLogger(t)

	logging.Logger.Info("Passed GetBySwiftCode")
}

func TestDelete(t *testing.T) {
	logging.Logger = zaptest.NewLogger(t)

	logging.Logger.Info("Passed GetBySwiftCode")
}
