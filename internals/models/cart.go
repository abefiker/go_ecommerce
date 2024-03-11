package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	ProductID int  `json:"product_id"`
	Quantity  int  `json:"quantity"`
	UserID    uint `json:"user_id"`
}