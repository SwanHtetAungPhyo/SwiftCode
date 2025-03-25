package repo

import (
	"errors"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RepositoryMethods interface {
	Create(req *model.SwiftCodeDto) error
	GetBySwiftCode(swiftCode string) ([]model.SwiftCodeDto, error)
	GetByCountryISO(countryISO2 string) ([]model.BankDetails, *model.Country, error)
	Delete(swiftCode string) error
}

type Repository struct {
	db      *gorm.DB
	repoLog *logrus.Logger
}

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

	bankDetails := model.BankDetails{
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
	var bankDetails []model.BankDetails
	swiftCodePrefix := swiftCode[:8] + "%"
	if err := r.db.Where("swift_code LIKE ?", swiftCodePrefix).Find(&bankDetails).Error; err != nil {
		r.repoLog.Errorln("Failed to fetch bank details:", err)
		return nil, err
	}

	if len(bankDetails) == 0 {
		return nil, errors.New("no bank details found")
	}

	var country model.Country
	if err := r.db.Where("id = ?", bankDetails[0].CountryID).First(&country).Error; err != nil {
		r.repoLog.Errorln("Failed to fetch country:", err)
		return nil, err
	}

	var swiftCodeDtos []model.SwiftCodeDto
	for _, b := range bankDetails {
		swiftCodeDtos = append(swiftCodeDtos, model.SwiftCodeDto{
			Address:       b.Address,
			BankName:      b.Name,
			CountryISO2:   country.CountryIso2Code,
			CountryName:   country.Name,
			IsHeadquarter: b.IsHeadquarter,
			SwiftCode:     b.SwiftCode,
		})
	}
	return swiftCodeDtos, nil
}

func (r *Repository) GetByCountryISO(countryISO2 string) ([]model.BankDetails, *model.Country, error) {
	var country model.Country
	if err := r.db.Where("country_iso2_code = ?", countryISO2).First(&country).Error; err != nil {
		r.repoLog.Errorln("Failed to fetch country:", err)
		return nil, nil, err
	}

	var bankDetails []model.BankDetails
	if err := r.db.Where("country_id = ?", country.ID).Find(&bankDetails).Error; err != nil {
		r.repoLog.Errorln("Failed to fetch bank details:", err)
		return nil, nil, err
	}
	return bankDetails, &country, nil
}

func (r *Repository) Delete(swiftCode string) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("swift_code = ?", swiftCode).Delete(&model.BankDetails{}).Error; err != nil {
		tx.Rollback()
		r.repoLog.Errorln("Failed to delete bank details:", err)
		return err
	}

	return tx.Commit().Error
}
