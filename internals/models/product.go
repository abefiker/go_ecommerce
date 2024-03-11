package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type ImageURLs []string

func (u *ImageURLs) Scan(value interface{}) error {
	if value == nil {
		*u = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSON value for ImageURLs")
	}
	return json.Unmarshal(bytes, u)
}

func (u ImageURLs) Value() (driver.Value, error) {
	if u == nil {
		return nil, nil
	}
	return json.Marshal(u)
}

type Product struct {
    gorm.Model
    Name        string    `json:"name" xml:"name" form:"name" query:"name"`
    Description string    `json:"description" xml:"description" form:"description" query:"description"`
    Price       float64   `json:"price" xml:"price" form:"price" query:"price"`
    Stock       int       `json:"stock" xml:"stock" form:"stock" query:"stock"`
    Images      ImageURLs `json:"images" xml:"images" form:"images" query:"images" gorm:"type:json"`
}
