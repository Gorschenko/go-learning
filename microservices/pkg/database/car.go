package database

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	UserID    uint   `gorm:"index;not null"`
	UserName  string `gorm:"not null;default:''"`
	ModelName string `gorm:"not null;default:''"`
	Brand     string `gorm:"not null;default:''"`
	Color     string `gorm:"not null;default:''"`
}
