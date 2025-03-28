package services

import "github.com/SwanHtetAungPhyo/swifcode/internal/model"

type ServiceMethods interface {
	Create(req *model.SwiftCodeDto) error
	GetBySwiftCode(swiftCode string) (*model.HeadquarterResponse, error)
	GetByCountryISO(countryISO2Code string) (*model.CountryISO2Response, error)
	Delete(swiftCode string) error
}
