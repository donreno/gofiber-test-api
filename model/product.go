package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey autoIncrement" json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
