package test

import (
	"context"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/repo"
	"github.com/SwanHtetAungPhyo/swifcode/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwiftCodeServices(t *testing.T) {
	ctx := context.Background()
	_, _, instance, err := SetupPostgresContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	log := logrus.New()
	repoInst := repo.NewRepository(instance, log)
	logger := logrus.New()
	service := services.NewService(repoInst, logger)
	swiftCodeDto := &model.SwiftCodeDto{
		SwiftCode:     "TESTSWIFTC",
		Address:       "123 Test Address",
		BankName:      "Test Bank",
		CountryName:   "USA",
		CountryISO2:   "US",
		IsHeadquarter: true,
	}
	t.Run("Create Swift Code - Success", func(t *testing.T) {

		err := service.Create(swiftCodeDto)
		assert.NoError(t, err)

	})

	t.Run("Create Swift Code - Failure", func(t *testing.T) {

		err := service.Create(swiftCodeDto)
		assert.Error(t, err)
	})

	t.Run("GetBySwiftCode - Found", func(t *testing.T) {

		result, err := service.GetBySwiftCode("TESTSWIFTC")
		assert.NoError(t, err)
		assert.Equal(t, "TESTSWIFTC", result.SwiftCode)
		assert.Equal(t, "123 Test Address", result.Address)

	})
	t.Run("GetBySwiftCode - Not Found", func(t *testing.T) {

		result, err := service.GetBySwiftCode("INVALID")
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("GetByCountryISO - Found", func(t *testing.T) {

		result, err := service.GetByCountryISO("US")
		assert.NoError(t, err)
		assert.Equal(t, "USA", result.CountryName)

	})
	t.Run("Delete Swift Code - Success", func(t *testing.T) {

		err := service.Delete("TESTSWIFTC")
		assert.NoError(t, err)

	})
	t.Run("Delete Swift Code - Failure", func(t *testing.T) {
		err := service.Delete("TESTSWIFTC")
		assert.NoError(t, err)
	})
}
