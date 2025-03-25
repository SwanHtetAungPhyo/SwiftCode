package services

import (
	"fmt"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BankProcessor struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewBankProcessor(db *gorm.DB, log *logrus.Logger) *BankProcessor {
	return &BankProcessor{
		db:  db,
		log: log,
	}
}

func (b *BankProcessor) ProcessData(filePath string) {
	data := utils.Parse(filePath, b.log)
	if len(data) == 0 {
		b.log.Fatal("No data parsed from CSV")
	}

	if err := b.migrateModels(); err != nil {
		b.log.Fatalf("Migration failed: %v", err)
	}

	countryMap, err := b.processCountries(data)
	if err != nil {
		b.log.Fatalf("Country processing failed: %v", err)
	}

	townIDMap, err := b.processTowns(data, countryMap)
	if err != nil {
		b.log.Fatalf("Town processing failed: %v", err)
	}

	if err := b.processBanks(data, countryMap, townIDMap); err != nil {
		b.log.Fatalf("Bank processing failed: %v", err)
	}
}

func (b *BankProcessor) migrateModels() error {
	return b.db.AutoMigrate(
		&model.Country{},
		&model.Town{},
		&model.BankDetails{},
	)
}

func (b *BankProcessor) processCountries(data []model.SwiftCode) (map[string]int, error) {
	countries := extractUniqueCountries(data)
	countryMap := make(map[string]int)

	for _, country := range countries {
		var dbCountry model.Country
		result := b.db.Where(model.Country{CountryIso2Code: country.CountryIso2Code}).
			FirstOrCreate(&dbCountry, country)
		if result.Error != nil {
			return nil, fmt.Errorf("country upsert failed: %v", result.Error)
		}
		countryMap[country.CountryIso2Code] = dbCountry.ID
	}

	b.log.Infof("Processed %d countries", len(countryMap))
	return countryMap, nil
}

func extractUniqueCountries(data []model.SwiftCode) []model.Country {
	countryMap := make(map[string]model.Country)
	for _, record := range data {
		key := record.CountryISO2
		if _, exists := countryMap[key]; !exists {
			countryMap[key] = model.Country{
				CountryIso2Code: record.CountryISO2,
				Name:            record.CountryName,
				TimeZone:        record.Timezone,
			}
		}
	}

	countries := make([]model.Country, 0, len(countryMap))
	for _, country := range countryMap {
		countries = append(countries, country)
	}
	return countries
}

func (b *BankProcessor) processTowns(data []model.SwiftCode, countryMap map[string]int) (map[string]int, error) {
	townDtos := extractUniqueTowns(data)
	var townsToInsert []model.Town

	for _, dto := range townDtos {
		countryID, exists := countryMap[dto.CountryIso2Code]
		if !exists {
			b.log.Warnf("Skipping town %s - country %s not found", dto.TownName, dto.CountryIso2Code)
			continue
		}
		townsToInsert = append(townsToInsert, model.Town{
			Name:      dto.TownName,
			CountryId: countryID,
		})
	}

	for _, town := range townsToInsert {
		var dbTown model.Town
		result := b.db.Where(model.Town{Name: town.Name, CountryId: town.CountryId}).
			FirstOrCreate(&dbTown, town)
		if result.Error != nil {
			return nil, fmt.Errorf("town upsert failed: %v", result.Error)
		}
	}

	return b.buildTownIDMap(), nil
}

func extractUniqueTowns(data []model.SwiftCode) []model.TownDto {
	townMap := make(map[string]model.TownDto)
	for _, record := range data {
		key := record.TownName + record.CountryISO2
		if _, exists := townMap[key]; !exists {
			townMap[key] = model.TownDto{
				TownName:        record.TownName,
				CountryIso2Code: record.CountryISO2,
			}
		}
	}

	towns := make([]model.TownDto, 0, len(townMap))
	for _, town := range townMap {
		towns = append(towns, town)
	}
	return towns
}

func (b *BankProcessor) buildTownIDMap() map[string]int {
	var towns []model.Town
	if err := b.db.Find(&towns).Error; err != nil {
		b.log.Errorf("Failed to fetch towns: %v", err)
		return nil
	}

	townIDMap := make(map[string]int)
	for _, town := range towns {
		key := fmt.Sprintf("%s|%d", town.Name, town.CountryId)
		townIDMap[key] = town.ID
	}
	return townIDMap
}

func (b *BankProcessor) processBanks(
	data []model.SwiftCode,
	countryMap map[string]int,
	townIDMap map[string]int,
) error {
	bankDtos := extractUniqueBanks(data)
	var banks []model.BankDetails

	for _, dto := range bankDtos {
		countryID, exists := countryMap[dto.CountryIso2Code]
		if !exists {
			b.log.Warnf("Skipping bank %s - country %s not found", dto.SwiftCode, dto.CountryIso2Code)
			continue
		}

		townKey := fmt.Sprintf("%s|%d", dto.TownName, countryID)
		townID, exists := townIDMap[townKey]
		if !exists {
			b.log.Warnf("Skipping bank %s - town %s not found", dto.SwiftCode, dto.TownName)
			continue
		}

		banks = append(banks, model.BankDetails{
			Name:          dto.Name,
			Address:       dto.Address,
			SwiftCode:     dto.SwiftCode,
			IsHeadquarter: dto.IsHeadquarter,
			CountryID:     countryID,
			TownNameId:    townID,
		})
	}

	if err := b.db.CreateInBatches(banks, 100).Error; err != nil {
		return fmt.Errorf("bank batch insert failed: %v", err)
	}

	b.log.Infof("Inserted %d bank records", len(banks))
	return nil
}

func extractUniqueBanks(data []model.SwiftCode) []model.BankDto {
	bankMap := make(map[string]model.BankDto)
	for _, record := range data {
		if _, exists := bankMap[record.SwiftCode]; !exists {
			bankMap[record.SwiftCode] = model.BankDto{
				Name:            record.BankName,
				Address:         record.Address,
				SwiftCode:       record.SwiftCode,
				TownName:        record.TownName,
				IsHeadquarter:   utils.IsHeadquarter(record.SwiftCode),
				CountryIso2Code: record.CountryISO2,
			}
		}
	}

	banks := make([]model.BankDto, 0, len(bankMap))
	for _, bank := range bankMap {
		banks = append(banks, bank)
	}
	return banks
}
