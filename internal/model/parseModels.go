package model

import "time"

type DetailsDto struct {
	Name            string
	Address         string
	SwiftCode       string
	TownName        string
	IsHeadquarter   bool
	CountryIso2Code string
}
type Details struct {
	Name            string
	Address         string
	SwiftCode       string
	TownName        string
	IsHeadquarter   bool
	CountryIso2Code string
	CountryId       int
}

type TownNameDto struct {
	TownName    string
	CountryName string
}
type Country struct {
	ID              int    `gorm:"primaryKey;autoIncrement"`
	CountryIso2Code string `gorm:"unique;size:2"`
	Name            string `gorm:"unique"`
	TimeZone        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Town struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	CountryId int `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BankDetails struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	Name          string
	Address       string
	SwiftCode     string `gorm:"size:11"`
	IsHeadquarter bool
	CountryID     int `gorm:"index"`
	TownNameId    int `gorm:"index"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// DTOs for data processing
type TownDto struct {
	TownName        string
	CountryIso2Code string
}

type BankDto struct {
	Name            string
	Address         string
	SwiftCode       string
	TownName        string
	IsHeadquarter   bool
	CountryIso2Code string
}

func (Country) TableName() string     { return "countries" }
func (Town) TableName() string        { return "towns" }
func (BankDetails) TableName() string { return "bank_details" }
