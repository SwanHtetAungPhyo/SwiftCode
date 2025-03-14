package utils

import (
	"encoding/csv"
	"fmt"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/logging"
	"github.com/SwanHtetAungPhyo/swifcode/internal/repo"
	"go.uber.org/zap"
	"log"
	"os"
	"strings"
)

func loadCSV(path string) ([]model.SwiftCode, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Failed to close file: %v", err)
		}
	}(file)

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	reader.ReuseRecord = true

	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	banksArray := make([]model.SwiftCode, 0)
	seen := make(map[string]struct{})

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		swiftCode := record[1]
		if _, exists := seen[swiftCode]; exists {
			log.Printf("Skipping duplicate SwiftCode: %s", swiftCode)
			continue
		}
		seen[swiftCode] = struct{}{}

		banksArray = append(banksArray, model.SwiftCode{
			CountryISO2: record[0],
			SwiftCode:   swiftCode,
			CodeType:    record[2],
			BankName:    record[3],
			Address:     record[4],
			TownName:    record[5],
			CountryName: record[6],
			Timezone:    record[7],
		})
	}
	return banksArray, nil
}

func InsertData(data []model.SwiftCode) {
	for i := range data {
		data[i].IsHeadquarter = isHeadquarter(data[i].SwiftCode)
	}

	count := 0
	for i, row := range data {
		if err := repo.DbInstance.Create(&row).Error; err != nil {
			if strings.Contains(err.Error(), "SQLSTATE 23505") {
				logging.Logger.Warn("Duplicate entry, skipping row", zap.Int("row_number", i+1), zap.String("swiftcode", row.SwiftCode))
				continue
			}
			logging.Logger.Error("Failed to create row", zap.Int("row_number", i+1), zap.Error(err))
			continue
		}
		count++
		logging.Logger.Info("Created row", zap.Int("count", count))
	}
	logging.Logger.Info(fmt.Sprintf("%d rows created", count))
}

func isHeadquarter(swiftCode string) bool {
	return strings.HasSuffix(swiftCode, "XXX")
}

func Parse(path string) []model.SwiftCode {
	data, err := loadCSV(path)
	if err != nil {
		log.Fatalf("Failed to load CSV: %v", err)
	}

	return data
}
