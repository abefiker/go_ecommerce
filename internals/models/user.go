package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string   `json:"email"`
	Password        string   `json:"-"`
	FirstName       string   `json:"firstName,omitempty"`
	LastName        string   `json:"lastName,omitempty"`
	ShippingAddress *Address `json:"shippingAddress,omitempty"`
}

type Address struct {
	gorm.Model
	UserID  uint   `json:"userId"`
	User    User   `gorm:"foreignKey:UserID"` // Define the User model as a foreign key
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
	Country string `json:"country"`
}
