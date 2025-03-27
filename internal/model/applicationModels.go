package model

// SwiftCode represents the data parsed from CSV
type SwiftCode struct {
	ID            uint   `gorm:"primaryKey"`
	CountryISO2   string `gorm:"not null" json:"countryISO2"`
	SwiftCode     string `gorm:"not null;uniqueIndex:idx_swift_code;unique" json:"swiftCode"`
	CodeType      string `gorm:"not null" json:"codeType"`
	BankName      string `gorm:"not null" json:"bankName"`
	Address       string `gorm:"not null" json:"address"`
	TownName      string `gorm:"not null" json:"townName"`
	CountryName   string `gorm:"not null" json:"countryName"`
	IsHeadquarter bool   `gorm:"not null" json:"isHeadquarter"`
	Timezone      string `gorm:"not null" json:"timezone"`
}

type CountryISO2Response struct {
	CountryISO2 string         `json:"countryISO2"`
	CountryName string         `json:"countryName"`
	SwiftCode   []SwiftCodeDto `json:"swiftCodes"`
}

// SwiftCodeDto BankDetail represents the details of a bank branch
type SwiftCodeDto struct {
	Address       string `json:"address,omitempty"`
	BankName      string `json:"bankName"`
	CountryISO2   string `json:"countryISO2"`
	CountryName   string `json:"countryName,omitempty"`
	IsHeadquarter bool   `json:"isHeadquarter"`
	SwiftCode     string `json:"swiftCode"`
}

// HeadquarterResponse represents the response with headquarter and associated branches
type HeadquarterResponse struct {
	Address       string         `json:"address,omitempty"`
	BankName      string         `json:"bankName"`
	CountryISO2   string         `json:"countryISO2"`
	CountryName   string         `json:"countryName"`
	IsHeadquarter bool           `json:"isHeadquarter"`
	SwiftCode     string         `json:"swiftCode"`
	Branches      []SwiftCodeDto `json:"branches,omitempty"`
}
