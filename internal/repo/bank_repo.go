package repo

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/logging"
	"gorm.io/gorm"
)

// Repository layer to interact with the database
// @Author: Swan Htet Aung Phyo

// BankRepoMethods
//
//	@Description: Define the Methods to perform the operation with SwiftCode Entities
type BankRepoMethods interface {
	//
	// Create
	//  @Description: Create the new record in the database and return the model to process the further operation
	//  @param model
	//  @return error
	//
	Create(model *model.SwiftCode) error
	//
	// GetBySwiftCode
	//  @Description:  Retrieve the  records which match with swift code
	//  @param swiftCode
	//  @return *model.SwiftCode
	//  @return error
	//
	GetBySwiftCode(swiftCode string) (*model.SwiftCode, error)
	//
	// GetByCountryISO2
	//  @Description: Retrieve the record which matches with countryISO2
	//  @param countryISO2
	//  @return *model.SwiftCode
	//  @return error
	//
	GetByCountryISO2(countryISO2 string) (*model.SwiftCode, error)
	//
	// DeleteBySwiftCode
	//  @Description: Delete the record by the swiftCode
	//  @param swiftCode
	//  @return error
	//
	DeleteBySwiftCode(swiftCode string) error
}

// BankRepoMethodImpl
// @Description: Implement the  methods of the BankRepoMethod to interact with database
type BankRepoMethodImpl struct {
	BankRepoMethods  BankRepoMethods
	DatabaseInstance *gorm.DB
}

// NewBankRepoMethodImpl
//
//	@Description: Create the new BankRepoMethodImpl to follow the dependency injection
//	@return *BankRepoMethodImpl The implementation of the BankRepoMethod
func NewBankRepoMethodImpl(databaseInstance *gorm.DB) *BankRepoMethodImpl {
	return &BankRepoMethodImpl{
		DatabaseInstance: databaseInstance,
	}
}

func (impl *BankRepoMethodImpl) Create(model *model.SwiftCode) error {
	logging.Logger.Info("Create Swift Code")
	if err := impl.DatabaseInstance.Create(model).Error; err != nil {
		logging.Logger.Error("Create Swift Code Failed")
		return err
	}
	return nil
}
func (impl *BankRepoMethodImpl) GetBySwiftCode(swiftCode string) (*model.SwiftCode, error) {
	return nil, nil
}
func (impl *BankRepoMethodImpl) GetByCountryISO2(countryISO2 string) (*model.SwiftCode, error) {
	return nil, nil
}
func (impl *BankRepoMethodImpl) DeleteBySwiftCode(swiftCode string) error {
	return nil
}
