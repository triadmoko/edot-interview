package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	Price      int
	Total      int
	CategoryID int
	Category   Category
}
