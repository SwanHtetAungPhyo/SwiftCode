package swiftcode_test

import (
	"context"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/repo"
	"github.com/SwanHtetAungPhyo/swifcode/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

// setupPostgresContainer initializes a PostgresSQL container for the tests
func SetupPostgresContainer(ctx context.Context) (testcontainers.Container, string, *gorm.DB, func(), error) {
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, "", nil, nil, err
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, "", nil, nil, err
	}

	instance, err := gorm.Open(driver.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, "", nil, nil, err
	}

	cleanup := func() {
		sqlDB, err := instance.DB()
		if err == nil {
			err := sqlDB.Close()
			if err != nil {
				return
			}
		}

		err = pgContainer.Terminate(ctx)
		if err != nil {
			return
		}
	}
	return pgContainer, connStr, instance, cleanup, nil
}
func truncateTablesAfterTest(db *gorm.DB) error {
	return db.Migrator().DropTable(&model.Country{}, &model.Town{}, &model.SwiftCode{})
}
func TestInsertion(t *testing.T) {
	ctx := context.Background()
	_, _, instance, _, err := SetupPostgresContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	//defer clean()
	err = instance.AutoMigrate(&model.Country{}, &model.Town{}, &model.SwiftCodeModel{})
	if err != nil {
		t.Fatal(err)
	}
	log := logrus.New()
	processor := services.NewBankProcessor(instance, log)
	processor.ProcessData("../data/swift_codes.csv")
	err = truncateTablesAfterTest(instance)
	if err != nil {
		log.Error(err.Error())
		return
	}
}

// TestSwiftCodeRepo performs tests on the repository layer
func TestSwiftCodeRepo(t *testing.T) {
	ctx := context.Background()
	_, _, instance, _, err := SetupPostgresContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	// Perform database migrations
	err = instance.AutoMigrate(&model.Country{}, &model.Town{}, &model.SwiftCodeModel{})
	assert.NoError(t, err)

	// Sample data
	countryToInsert := &model.Country{
		ID:              1,
		CountryIso2Code: "US",
		Name:            "USA",
		TimeZone:        "Europe/London",
	}
	townName := &model.Town{
		ID:        1,
		Name:      "New York",
		CountryId: 1,
	}
	swiftCode := &model.SwiftCodeModel{
		ID:            1,
		Address:       "TEST ADDRESS",
		SwiftCode:     "TESTSWIFTC",
		IsHeadquarter: true,
		CountryID:     1,
		TownNameId:    1,
		CodeType:      "SWIFTCODE",
	}
	swiftCodeNew := &model.SwiftCodeDto{
		Address:       "TESADDRESS",
		SwiftCode:     "TESTSWCXX",
		IsHeadquarter: true,
		CountryName:   "USA",
		CountryISO2:   "US",
		BankName:      "BANK",
	}
	// Log setup
	log := logrus.New()
	repoInst := repo.NewRepository(instance, log)

	// Test case: Insert data
	t.Run("insert", func(t *testing.T) {
		err = instance.Create(countryToInsert).Error
		assert.NoError(t, err)
		err = instance.Create(townName).Error
		assert.NoError(t, err)
		err = instance.Create(swiftCode).Error
		assert.NoError(t, err)
	})

	t.Run("create", func(t *testing.T) {
		err = repoInst.Create(swiftCodeNew)
		assert.Error(t, err)
	})
	// Test case: Get by swift code
	t.Run("getBySwiftCode", func(t *testing.T) {
		var swiftCode []model.SwiftCodeDto
		swiftCode, err = repoInst.GetBySwiftCode("TESTSWIFTC")
		assert.NoError(t, err)
		assert.Len(t, swiftCode, 1)
	})
	t.Run("GetbyISO2Code", func(t *testing.T) {
		swiftCodeDto, country, err := repoInst.GetByCountryISO("US")
		assert.NoError(t, err)
		assert.Len(t, swiftCodeDto, 1)
		assert.Equal(t, country, countryToInsert)
	})
	// Test case: Retrieve Country by ISO code
	t.Run("getCountryByISO", func(t *testing.T) {
		countryFromDB, err := repoInst.GetCountryByISO("US")
		assert.NoError(t, err)
		assert.Equal(t, countryFromDB, countryToInsert)
	})

}
