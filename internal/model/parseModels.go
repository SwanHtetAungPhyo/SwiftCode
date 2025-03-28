package model

import "time"

// DetailsDto For CSV parsing and data preparation
type DetailsDto struct {
	Name            string
	Address         string
	SwiftCode       string
	TownName        string
	IsHeadquarter   bool
	CountryIso2Code string
}

// Details For CSV parsing and data preparation
type Details struct {
	Name            string
	Address         string
	SwiftCode       string
	TownName        string
	IsHeadquarter   bool
	CountryIso2Code string
	CountryId       int
}

// TownNameDto For CSV parsing and data preparation
type TownNameDto struct {
	TownName    string
	CountryName string
}

// Country for gorm model
type Country struct {
	ID              int    `gorm:"primaryKey;autoIncrement"`
	CountryIso2Code string `gorm:"unique;size:2"`
	Name            string `gorm:"unique"`
	TimeZone        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Town for gorm model
type Town struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	CountryId int `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SwiftCodeModel for the gorm model
type SwiftCodeModel struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	Name          string
	Address       string
	SwiftCode     string `gorm:"unique;size:11"`
	IsHeadquarter bool
	CountryID     int `gorm:"index"`
	TownNameId    int `gorm:"index"`
	CodeType      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TownDto TownDTOs for data processing to be ready to insert to town table
type TownDto struct {
	TownName        string
	CountryIso2Code string
}

// BankDto for data processing to be ready to insert to swift-code table
type BankDto struct {
	Name            string
	Address         string
	SwiftCode       string
	TownName        string
	IsHeadquarter   bool
	CountryIso2Code string
}

func (Country) TableName() string        { return "countries" }
func (Town) TableName() string           { return "towns" }
func (SwiftCodeModel) TableName() string { return "swiftcodes" }
