package models

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    UserID     uint        `json:"userId"`
    User       User        `gorm:"foreignKey:UserID"` // Define the User model as a foreign key
    Items      []OrderItem `json:"items"`
    TotalPrice float64     `json:"totalPrice"`
    Status     string      `json:"status"`
}

type OrderItem struct {
    gorm.Model
    OrderID   uint   `json:"orderId"`
    ProductID uint   `json:"productId"`
    Quantity  int    `json:"quantity"`
}
