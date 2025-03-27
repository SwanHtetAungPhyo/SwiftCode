package mocks

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockServiceMethods struct {
	mock.Mock
}

func (m *MockServiceMethods) Create(req *model.SwiftCodeDto) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockServiceMethods) GetBySwiftCode(swiftCode string) (*model.HeadquarterResponse, error) {
	args := m.Called(swiftCode)
	if args.Get(0) != nil {
		return args.Get(0).(*model.HeadquarterResponse), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockServiceMethods) GetByCountryISO(countryISO2Code string) (*model.CountryISO2Response, error) {
	args := m.Called(countryISO2Code)
	if args.Get(0) != nil {
		return args.Get(0).(*model.CountryISO2Response), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockServiceMethods) Delete(swiftCode string) error {
	args := m.Called(swiftCode)
	return args.Error(0)
}
