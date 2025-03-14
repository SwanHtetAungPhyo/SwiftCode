package services

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/repo"
)

// Service Layer to process the data manipulation, interact with repo (database layer) and cast the application logic to meet the Business Logic
// Author: Swan Htet Aung Phyo

// SwiftCodeService
// @Description:  Service Layer Interface to make the abstraction for the logic Layer(service Layer) for the swiftCode
type SwiftCodeService interface {
	//
	// GetBySwiftCode
	//  @Description:  Get SwiftCode by the swift code by sending the signal (code) to the repo
	//  @param swiftCode
	//  @return model.SwiftCode
	//  @return error
	//
	GetBySwiftCode(swiftCode string) (model.SwiftCode, error)
	//
	// GetWithISO2
	//  @Description: Get with countryISO 2 and return array of swiftcode
	//  @param ISO2
	//  @return []model.SwiftCode
	//
	GetWithISO2(ISO2 string) []model.SwiftCode
	//
	// Create
	//  @Description: Create the SwiftCode
	//  @param swiftCode
	//
	Create(swiftCode *model.SwiftCode)
	//
	// Update
	//  @Description: Delete the
	//  @param swiftCode
	//  @return *model.SwiftCode
	//
	Update(swiftCode *model.SwiftCode) *model.SwiftCode
}

// SwiftCodeServiceImpl
// @Description:
type SwiftCodeServiceImpl struct {
	swiftCodeService SwiftCodeService
	swiftCodeRepo    *repo.BankRepoMethodImpl
}

func NewSwiftCodeService(repo *repo.BankRepoMethodImpl) *SwiftCodeServiceImpl {
	return &SwiftCodeServiceImpl{
		swiftCodeRepo: repo,
	}
}

func (impl *SwiftCodeServiceImpl) GetBySwiftCode(swiftCode *model.SwiftCode) {
	//var swiftCodeModel model.SwiftCode

	return
}

func (impl *SwiftCodeServiceImpl) GetWithISO2(ISO2 string) []model.SwiftCode {
	var swiftCode model.SwiftCode
	var array []model.SwiftCode
	//impl.DB.First(&swiftCode, "country_iso2 = ?", ISO2)
	array = append(array, swiftCode)
	return array
}

func (impl *SwiftCodeServiceImpl) Create(swiftCode *model.SwiftCode) {
	//impl.DB.Create(swiftCode)
}

func (impl *SwiftCodeServiceImpl) Update(swiftCode *model.SwiftCode) *model.SwiftCode {
	return nil
}
