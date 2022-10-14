package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	Total      int      `json:"total"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignKey:category_id"`
}
