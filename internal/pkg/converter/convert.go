package converter

import "github.com/SwanHtetAungPhyo/swifcode/internal/model"

func ConvertToSwiftCodeModel(request model.SwiftCodeAddRequest) *model.SwiftCode {
	return &model.SwiftCode{
		Address:       request.Address,
		BankName:      request.BankName,
		SwiftCode:     request.SwiftCode,
		CountryISO2:   request.CountryISO2,
		CountryName:   request.CountryName,
		IsHeadquarter: request.IsHeadquarter,
	}
}
