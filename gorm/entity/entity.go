package entity

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`
	Num      int    `gorm:"AUTO_INCREMENT"`

	CreditCard CreditCard
	Emails     []Email

	BillingAddress   Address
	BillingAddressID sql.NullInt64

	ShippingAddress   Address
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Email struct {
	ID         int
	UserID     int    `gorm:index`
	Email      string `gorm:"type:varchar(100);unique_index"`
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"`
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `go:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"`
	Code string `gorm:"index:idx_name_code"`
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string `gorm:"default:'huawei'"`
}
