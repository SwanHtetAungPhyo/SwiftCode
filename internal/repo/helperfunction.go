package repo

import "github.com/SwanHtetAungPhyo/swifcode/internal/model"

// GetCountryByBankDetails Helper functions For the repository Layer
func (r *Repository) GetCountryByBankDetails(bd model.SwiftCodeModel) (*model.Country, error) {
	var country model.Country
	if err := r.db.Where("id = ?", bd.CountryID).First(&country).Error; err != nil {
		r.repoLog.Errorln("Failed to fetch country:", err)
		return nil, err
	}
	return &country, nil
}

func (r *Repository) GetCountryByISO(isoCode string) (*model.Country, error) {
	var country model.Country
	if err := r.db.Where("country_iso2_code = ?", isoCode).First(&country).Error; err != nil {
		r.repoLog.Errorln("Failed to fetch country:", err)
		return nil, err
	}
	return &country, nil
}

func convertToSwiftCodeDTOS(banks []model.SwiftCodeModel, country *model.Country) []model.SwiftCodeDto {
	dataDto := make([]model.SwiftCodeDto, len(banks))
	for i, b := range banks {
		dataDto[i] = model.SwiftCodeDto{
			Address:       b.Address,
			BankName:      b.Name,
			CountryISO2:   country.CountryIso2Code,
			CountryName:   country.Name,
			IsHeadquarter: b.IsHeadquarter,
			SwiftCode:     b.SwiftCode,
		}
	}
	return dataDto
}
