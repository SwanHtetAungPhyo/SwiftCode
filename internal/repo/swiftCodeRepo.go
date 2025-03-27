package repo

import (
	"errors"
	"fmt"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// RepositoryMethods interface defines database operations
type RepositoryMethods interface {
	Create(req *model.SwiftCodeDto) error
	GetBySwiftCode(swiftCode string) ([]model.SwiftCodeDto, error)
	GetByCountryISO(countryISO2 string) ([]model.SwiftCodeModel, *model.Country, error)
	Delete(swiftCode string) error
}

// Repository struct
type Repository struct {
	db      *gorm.DB
	repoLog *logrus.Logger
}

// Ensure For Repository implements RepositoryMethods
var _ RepositoryMethods = (*Repository)(nil)

// NewRepository Constructor function
func NewRepository(db *gorm.DB, repoLog *logrus.Logger) *Repository {
	return &Repository{db: db, repoLog: repoLog}
}

func (r *Repository) Create(req *model.SwiftCodeDto) error {
	r.repoLog.Infoln("Creating new swift code...")
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	countryID, err := r.getOrCreateCountryID(tx, req.CountryISO2, req.CountryName)
	if err != nil {
		tx.Rollback()
		r.repoLog.Errorln("Failed to process country:", err)
		return err
	}

	bankDetails := model.SwiftCodeModel{
		Address:       req.Address,
		Name:          req.BankName,
		SwiftCode:     req.SwiftCode,
		IsHeadquarter: req.IsHeadquarter,
		CountryID:     countryID,
	}

	if err := tx.Create(&bankDetails).Error; err != nil {
		tx.Rollback()
		r.repoLog.Errorln("Failed to create bank details:", err)
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) getOrCreateCountryID(tx *gorm.DB, countryISO2, countryName string) (int, error) {
	var country model.Country
	if err := tx.Where("country_iso2_code = ?", countryISO2).First(&country).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newCountry := model.Country{
				CountryIso2Code: countryISO2,
				Name:            countryName,
				TimeZone:        "UTC",
			}
			if err := tx.Create(&newCountry).Error; err != nil {
				return 0, err
			}
			return newCountry.ID, nil
		}
		return 0, err
	}
	return country.ID, nil
}

func (r *Repository) GetBySwiftCode(swiftCode string) ([]model.SwiftCodeDto, error) {
	if len(swiftCode) < 8 {
		return nil, errors.New("invalid swift code format")
	}

	swiftCodePrefix := fmt.Sprintf("%s%%", swiftCode[:8])
	var bankDetails []model.SwiftCodeModel
	if err := r.db.Where("swift_code LIKE ?", swiftCodePrefix).Find(&bankDetails).Error; err != nil {
		r.repoLog.Errorln("Failed to fetch bank details:", err)
		return nil, err
	}

	if len(bankDetails) == 0 {
		return nil, errors.New("no bank details found")
	}

	country, err := r.GetCountryByBankDetails(bankDetails[0])
	if err != nil {
		return nil, err
	}

	return convertToSwiftCodeDTOS(bankDetails, country), nil
}

func (r *Repository) GetByCountryISO(countryISO2 string) ([]model.SwiftCodeModel, *model.Country, error) {
	country, err := r.GetCountryByISO(countryISO2)
	if err != nil {
		return nil, nil, err
	}

	var bankDetails []model.SwiftCodeModel
	if err := r.db.Where("country_id = ?", country.ID).Find(&bankDetails).Error; err != nil {
		r.repoLog.Errorln("Failed to fetch bank details:", err)
		return nil, nil, err
	}
	return bankDetails, country, nil
}

func (r *Repository) Delete(swiftCode string) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("swift_code = ?", swiftCode).Delete(&model.SwiftCodeModel{}).Error; err != nil {
		tx.Rollback()
		r.repoLog.Errorln("Failed to delete bank details:", err)
		return err
	}

	return tx.Commit().Error
}
