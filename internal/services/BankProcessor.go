package services

import (
	"fmt"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/repo"
	"github.com/SwanHtetAungPhyo/swifcode/pkg/utils"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
)

func DataProcessing(log *logrus.Logger, filepath string) {
	data := utils.Parse(filepath, log)
	if len(data) == 0 {
		log.Fatal("No data parsed from CSV")
	}

	err := repo.DbInstance.AutoMigrate(&model.Country{}, &model.BankDetails{})
	if err != nil {
		return
	}

	countries := ExtractCountries(data)
	bankDto := ExtractBanksDto(data)
	var banks []model.Details

	result := repo.DbInstance.Create(&countries)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Fatal(result.Error)
	}
	for _, dto := range bankDto {
		bank := model.Details{
			Name:            dto.Name,
			Address:         dto.Address,
			SwiftCode:       dto.SwiftCode,
			TownName:        dto.TownName,
			IsHeadquarter:   dto.IsHeadquarter,
			CountryIso2Code: dto.CountryIso2Code,
		}
		banks = append(banks, bank)
	}
	fmt.Println("Total Countries:", len(countries))
	fmt.Println("Total Banks:", len(banks))

	mapISO := fetchingIsoCodeFromDB(countries)
	if mapISO == nil {
		log.Fatal("No ISO code in DB")
	}
	banksDataReadForInsertion := assignTheCountryID(banks, mapISO)
	if banksDataReadForInsertion == nil {
		log.Fatal("Data is not ready in DB")
	}
	dataInsertion(banksDataReadForInsertion)
}
func fetchingIsoCodeFromDB(countries []model.Country) map[string]int {
	idINDbByCountryIsoCode := make(map[string]int)
	for _, country := range countries {
		var dbCountry model.Country
		if err := repo.DbInstance.Where("countryiso2code = ?", country.CountryIso2Code).First(&dbCountry).Error; err != nil {
			log.Errorf("Failed to fetch country ID for %s: %v", country.CountryIso2Code, err)
		}
		idINDbByCountryIsoCode[country.CountryIso2Code] = dbCountry.ID
	}
	return idINDbByCountryIsoCode
}

func assignTheCountryID(banks []model.Details, idInDbByCountryIsoCode map[string]int) []model.Details {
	for i := range banks {
		if countryID, ok := idInDbByCountryIsoCode[banks[i].CountryIso2Code]; ok {
			banks[i].CountryId = countryID
		} else {
			log.Warnf("Country ID not found for %s", banks[i].CountryIso2Code)
		}
	}
	return banks
}

func dataInsertion(banks []model.Details) {
	var insertion = make([]model.BankDetails, 0, len(banks))
	for _, bank := range banks {
		insert := model.BankDetails{
			Name:          bank.Name,
			Address:       bank.Address,
			SwiftCode:     bank.SwiftCode,
			TownName:      bank.TownName,
			IsHeadquarter: bank.IsHeadquarter,
			CountryID:     bank.CountryId,
		}
		insertion = append(insertion, insert)
	}
	if len(insertion) == 0 {
		log.Info("No data to insert")
		return
	}
	for _, insertBank := range insertion {
		if err := repo.DbInstance.Create(&insertBank).Error; err != nil {
			log.Errorf("Failed to insert bank details: %v", err)
		}
	}
}

// ExtractCountries Extract unique countries from data
func ExtractCountries(data []model.SwiftCode) []model.Country {
	countryMap := make(map[string]model.Country)
	for _, code := range data {
		if _, exists := countryMap[code.CountryISO2]; !exists {
			countryMap[code.CountryISO2] = model.Country{
				CountryIso2Code: code.CountryISO2,
				Name:            code.CountryName,
				TimeZone:        code.Timezone,
			}
		}
	}
	countries := make([]model.Country, 0, len(countryMap))
	for _, country := range countryMap {
		countries = append(countries, country)
	}
	return countries
}

// ExtractBanksDto Extract unique banks from data
func ExtractBanksDto(data []model.SwiftCode) []model.DetailsDto {
	bankMap := make(map[string]model.DetailsDto)
	for _, code := range data {
		if _, exists := bankMap[code.SwiftCode]; !exists {
			bankMap[code.SwiftCode] = model.DetailsDto{
				Name:            code.BankName,
				Address:         code.Address,
				SwiftCode:       code.SwiftCode,
				TownName:        code.TownName,
				IsHeadquarter:   utils.IsHeadquarter(code.SwiftCode),
				CountryIso2Code: code.CountryISO2,
			}
		}
	}
	banks := make([]model.DetailsDto, 0, len(bankMap))
	for _, bank := range bankMap {
		banks = append(banks, bank)
	}
	return banks
}
