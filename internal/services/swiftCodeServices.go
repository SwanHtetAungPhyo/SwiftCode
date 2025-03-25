package services

import (
	"errors"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/repo"
	"github.com/SwanHtetAungPhyo/swifcode/pkg/utils"
	"github.com/sirupsen/logrus"
)

type ServiceMethods interface {
	Create(req *model.SwiftCodeDto) error
	GetBySwiftCode(swiftCode string) (*model.HeadquarterResponse, error)
	GetByCountryISO(countryISO2Code string) (*model.CountryISO2Response, error)
	Delete(swiftCode string) error
}

type SwiftCodeServices struct {
	repo   repo.RepositoryMethods
	logger *logrus.Logger
}

var _ ServiceMethods = (*SwiftCodeServices)(nil)

func NewService(repo repo.RepositoryMethods, logger *logrus.Logger) *SwiftCodeServices {
	return &SwiftCodeServices{
		repo:   repo,
		logger: logger,
	}
}

func (s *SwiftCodeServices) Create(req *model.SwiftCodeDto) error {
	s.logger.Infof("[Service] Creating Swift Code: %s", req.SwiftCode)
	if err := s.repo.Create(req); err != nil {
		s.logger.Errorf("[Service] Failed to create Swift Code: %v", err)
		return err
	}
	s.logger.Infof("[Service] Successfully created Swift Code: %s", req.SwiftCode)
	return nil
}

func (s *SwiftCodeServices) GetBySwiftCode(swiftCode string) (*model.HeadquarterResponse, error) {
	s.logger.Infof("[Service] Fetching details for Swift Code: %s", swiftCode)
	swiftCodes, err := s.repo.GetBySwiftCode(swiftCode)
	if err != nil {
		s.logger.Errorf("[Service] Error fetching Swift Code: %v", err)
		return nil, err
	}

	if len(swiftCodes) == 0 {
		s.logger.Errorf("[Service] Swift Code not found: %s", swiftCode)
		return nil, errors.New("swift Code not found")
	}
	if len(swiftCodes) == 1 {
		s.logger.Errorf("[Service] Found more than one Swift Code: %s", swiftCode)
		return &model.HeadquarterResponse{
			Address:       swiftCodes[0].Address,
			SwiftCode:     swiftCodes[0].SwiftCode,
			BankName:      swiftCodes[0].BankName,
			CountryName:   swiftCodes[0].CountryName,
			CountryISO2:   swiftCodes[0].CountryISO2,
			IsHeadquarter: swiftCodes[0].IsHeadquarter,
		}, nil
	}

	var headquarter model.SwiftCodeDto
	for i, value := range swiftCodes {
		if utils.IsHeadquarter(value.SwiftCode) {
			headquarter = value
			swiftCodes = append(swiftCodes[:i], swiftCodes[i+1:]...)
			break
		}
	}

	s.logger.Infof("[Service] Found Headquarter: %s", headquarter.Address)
	return &model.HeadquarterResponse{
		Address:       headquarter.Address,
		SwiftCode:     headquarter.SwiftCode,
		BankName:      headquarter.BankName,
		CountryName:   headquarter.CountryName,
		CountryISO2:   headquarter.CountryISO2,
		IsHeadquarter: headquarter.IsHeadquarter,
		Branches:      swiftCodes,
	}, nil
}

func (s *SwiftCodeServices) GetByCountryISO(countryISO2Code string) (*model.CountryISO2Response, error) {
	s.logger.Infof("[Service] Fetching Swift Codes for country: %s", countryISO2Code)
	swiftCodes, country, err := s.repo.GetByCountryISO(countryISO2Code)
	if err != nil {
		s.logger.Errorf("[Service] Error fetching Swift Codes for country %s: %v", countryISO2Code, err)
		return nil, err
	}

	var swiftCodeDtos []model.SwiftCodeDto
	for _, value := range swiftCodes {
		swiftCodeDtos = append(swiftCodeDtos, model.SwiftCodeDto{
			Address:       value.Address,
			BankName:      value.Name,
			CountryISO2:   countryISO2Code,
			SwiftCode:     value.SwiftCode,
			IsHeadquarter: value.IsHeadquarter,
		})
	}

	s.logger.Infof("[Service] Successfully fetched Swift Codes for country: %s", countryISO2Code)
	return &model.CountryISO2Response{
		CountryISO2: country.CountryIso2Code,
		CountryName: country.Name,
		SwiftCode:   swiftCodeDtos,
	}, nil
}

func (s *SwiftCodeServices) Delete(swiftCode string) error {
	s.logger.Infof("[Service] Deleting Swift Code: %s", swiftCode)
	if err := s.repo.Delete(swiftCode); err != nil {
		s.logger.Errorf("[Service] Failed to delete Swift Code %s: %v", swiftCode, err)
		return err
	}
	s.logger.Infof("[Service] Successfully deleted Swift Code: %s", swiftCode)
	return nil
}
