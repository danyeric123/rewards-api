package db

import (
    "gorm.io/gorm"
)

type Receipt struct {
    gorm.Model
    ID           string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Retailer     string
    PurchaseDate string
    PurchaseTime string
    Total        float64
    Points       int
}

type Item struct {
    gorm.Model
    ID              string  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    ReceiptID       string  `gorm:"type:uuid"`
    ShortDescription string
    Price           float64
}