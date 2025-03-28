package repo

import "github.com/SwanHtetAungPhyo/swifcode/internal/model"

// RepositoryMethods interface defines database operations
type RepositoryMethods interface {
	Create(req *model.SwiftCodeDto) error
	GetBySwiftCode(swiftCode string) ([]model.SwiftCodeDto, error)
	GetByCountryISO(countryISO2 string) ([]model.SwiftCodeModel, *model.Country, error)
	Delete(swiftCode string) error
}
