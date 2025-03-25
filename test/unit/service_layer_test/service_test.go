package service_layer_test

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/log"
	"go.uber.org/zap/zaptest"
	"testing"
)

func TestGetBySwiftCode(t *testing.T) {
	log.Logger = zaptest.NewLogger(t)

	log.Logger.Info("Passed GetBySwiftCode")
}
func TestGetByISO2(t *testing.T) {
	log.Logger = zaptest.NewLogger(t)

	log.Logger.Info("Passed GetBySwiftCode")
}

func TestCreate(t *testing.T) {
	log.Logger = zaptest.NewLogger(t)

	log.Logger.Info("Passed GetBySwiftCode")
}

func TestDelete(t *testing.T) {
	log.Logger = zaptest.NewLogger(t)

	log.Logger.Info("Passed GetBySwiftCode")
}
