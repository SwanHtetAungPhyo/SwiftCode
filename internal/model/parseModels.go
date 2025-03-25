package model

// TableName Define table names explicitly
func (Details) TableName() string { return "bank_details" }
func (Country) TableName() string { return "countries" }

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
type Country struct {
	ID              int           `gorm:"primaryKey;autoIncrement;column:id"`
	CountryIso2Code string        `gorm:"column:countryiso2code;type:char(2);not null;unique"`
	Name            string        `gorm:"column:name;type:varchar(255);not null;unique"`
	TimeZone        string        `gorm:"column:timezone;type:varchar(255);not null"`
	BankDetails     []BankDetails `gorm:"foreignKey:CountryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// use for the data insertion
type BankDetails struct {
	ID            int     `gorm:"primaryKey;autoIncrement;column:id"`
	Name          string  `gorm:"column:name;type:varchar(255);not null"`
	Address       string  `gorm:"column:address;type:varchar(255);not null"`
	TownName      string  `gorm:"column:town_name;type:varchar(255);not null"`
	SwiftCode     string  `gorm:"column:swift_code;type:varchar(12);not null"`
	IsHeadquarter bool    `gorm:"column:is_headquarter;type:boolean;not null"`
	CountryID     int     `gorm:"column:countryid;not null"`
	Country       Country `gorm:"foreignKey:CountryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
