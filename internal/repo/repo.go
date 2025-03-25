package repo

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/sirupsen/logrus"
)

// MigrateTables ensures tables exist
func MigrateTables(log *logrus.Logger) error {
	log.Info("Running database migrations...")
	return DbInstance.AutoMigrate(&model.Country{}, &model.BankDetails{})
}

// InsertCountries saves country records to DB
func InsertCountries(countries []model.Country, log *logrus.Logger) {
	if len(countries) == 0 {
		log.Warn("No countries to insert")
		return
	}
	if err := DbInstance.Create(&countries).Error; err != nil {
		log.Fatal("Failed to insert countries: ", err)
	}
}

// FetchCountryIDs maps country ISO codes to DB IDs
func FetchCountryIDs(countries []model.Country, log *logrus.Logger) map[string]int {
	idMap := make(map[string]int)
	for _, country := range countries {
		var dbCountry model.Country
		if err := DbInstance.Where("countryiso2code = ?", country.CountryIso2Code).First(&dbCountry).Error; err != nil {
			log.Warnf("Country ID not found for %s: %v", country.CountryIso2Code, err)
			continue
		}
		idMap[country.CountryIso2Code] = dbCountry.ID
	}
	return idMap
}

// InsertBankDetails inserts bank details into DB
func InsertBankDetails(banks []model.BankDetails, log *logrus.Logger) {
	if len(banks) == 0 {
		log.Info("No bank data to insert")
		return
	}
	for _, bank := range banks {
		if err := DbInstance.Create(&bank).Error; err != nil {
			log.Errorf("Failed to insert bank: %v", err)
		}
	}
}
