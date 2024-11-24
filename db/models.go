package db

import (
    "gorm.io/gorm"
)

type Receipt struct {
    gorm.Model
    ID           string `gorm:"type:uuid;primaryKey"`
    Retailer     string
    PurchaseDate string
    PurchaseTime string
    Total        float64
    Points       int
}

type Item struct {
    gorm.Model
    ID              string  `gorm:"type:uuid;primaryKey"`
    ReceiptID       string  `gorm:"type:uuid"`
    ShortDescription string
    Price           float64
}