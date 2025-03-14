package model

// SwiftCode represents the data parsed from CSV
type SwiftCode struct {
	ID            uint   `gorm:"primaryKey"`
	CountryISO2   string `gorm:"not null" json:"countryISO2"`
	SwiftCode     string `gorm:"not null" json:"swiftCode"`
	CodeType      string `gorm:"not null" json:"codeType"`
	BankName      string `gorm:"not null" json:"bankName"`
	Address       string `gorm:"not null" json:"address"`
	TownName      string `gorm:"not null" json:"townName"`
	CountryName   string `gorm:"not null" json:"countryName"`
	IsHeadquarter bool   `gorm:"not null" json:"isHeadquarter"`
	Timezone      string `gorm:"not null" json:"timezone"`
}

// SwiftCodeAddRequest represents the request payload to create a new SwiftCode
type SwiftCodeAddRequest struct {
	Address       string `json:"address" validate:"required,min=5,max=255"`
	BankName      string `json:"bankName" validate:"required,min=3,max=255"`
	CountryISO2   string `json:"countryISO2" validate:"required,len=2,iso3166_1_alpha2"`
	CountryName   string `json:"countryName" validate:"required,min=2,max=100"`
	IsHeadquarter bool   `json:"isHeadquarter" validate:"required"`
	SwiftCode     string `json:"swiftCode" validate:"required,len=8,alphanum"`
}

// BankDetail represents the details of a bank branch
type BankDetail struct {
	Address       string `json:"address"`
	BankName      string `json:"bankName"`
	CountryISO2   string `json:"countryISO2"`
	CountryName   string `json:"countryName"`
	IsHeadquarter bool   `json:"isHeadquarter"`
	SwiftCode     string `json:"swiftCode"`
}

// HeadquarterResponse represents the response with headquarter and associated branches
type HeadquarterResponse struct {
	Address       string       `json:"address"`
	BankName      string       `json:"bankName"`
	CountryISO2   string       `json:"countryISO2"`
	CountryName   string       `json:"countryName"`
	IsHeadquarter bool         `json:"isHeadquarter"`
	SwiftCode     string       `json:"swiftCode"`
	BankDetails   []BankDetail `json:"branches"`
}
