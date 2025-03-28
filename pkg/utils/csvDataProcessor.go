package utils

import (
	"encoding/csv"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
)

func LoadCSV(path string, log *logrus.Logger) ([]model.SwiftCode, error) {
	file, err := os.Open(filepath.Join(path))
	if err != nil {
		log.Error("Failed to open file", zap.String("path", path), zap.Error(err))
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Error("Failed to close file", zap.String("path", path), zap.Error(err))
		}
	}()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	reader.ReuseRecord = true

	if _, err := reader.Read(); err != nil {
		log.Error("Failed to read header", zap.Error(err))
		return nil, err
	}

	banksArray := make([]model.SwiftCode, 0)
	seen := make(map[string]struct{})

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Error("Failed to read record", zap.Error(err))
			continue
		}

		swiftCode := record[1]
		if _, exists := seen[swiftCode]; exists {
			log.Warn("Skipping duplicate SwiftCode", zap.String("swiftCode", swiftCode))
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
	for i := range banksArray {
		banksArray[i].IsHeadquarter = IsHeadquarter(banksArray[i].SwiftCode)
	}
	return banksArray, nil
}

func IsHeadquarter(swiftCode string) bool {
	return strings.HasSuffix(swiftCode, "XXX")
}

func Parse(path string, log *logrus.Logger) []model.SwiftCode {
	data, err := LoadCSV(path, log)
	if err != nil {
		log.Fatal("Failed to load CSV", zap.String("path", path), zap.Error(err))
	}

	return data
}
