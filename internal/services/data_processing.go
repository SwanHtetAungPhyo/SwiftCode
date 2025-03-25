package services

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/pkg/utils"
	"github.com/sirupsen/logrus"
)

func ParseAndProcessData(filePath string, log *logrus.Logger) ([]model.Country, []model.DetailsDto) {
	data := utils.Parse(filePath, log)
	if len(data) == 0 {
		log.Fatal("No data parsed from CSV")
	}

	// Extracting Countries and Banks
	countries := ExtractCountries(data)
	banks := ExtractBanksDto(data)

	log.Infof("Extracted %d countries and %d banks", len(countries), len(banks))
	return countries, banks
}

//	func ExtractCountries(data []model.SwiftCode) []model.Country {
//		countryMap := make(map[string]model.Country)
//		for _, code := range data {
//			if _, exists := countryMap[code.CountryISO2]; !exists {
//				countryMap[code.CountryISO2] = model.Country{
//					CountryIso2Code: code.CountryISO2,
//					Name:            code.CountryName,
//					TimeZone:        code.Timezone,
//				}
//			}
//		}
//		return mapToSlice(countryMap)
//	}
//
//	func ExtractBanksDto(data []model.SwiftCode) []model.DetailsDto {
//		bankMap := make(map[string]model.DetailsDto)
//		for _, code := range data {
//			if _, exists := bankMap[code.SwiftCode]; !exists {
//				bankMap[code.SwiftCode] = model.DetailsDto{
//					Name:            code.BankName,
//					Address:         code.Address,
//					SwiftCode:       code.SwiftCode,
//					TownName:        code.TownName,
//					IsHeadquarter:   utils.IsHeadquarter(code.SwiftCode),
//					CountryIso2Code: code.CountryISO2,
//				}
//			}
//		}
//		return mapToSlice(bankMap)
//	}
func mapToSlice[T any](m map[string]T) []T {
	slice := make([]T, 0, len(m))
	for _, v := range m {
		slice = append(slice, v)
	}
	return slice
}
