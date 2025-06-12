package model

import (
	"time"
)

type Cat struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"not null"`
	YearsExperience int       `json:"years_experience" gorm:"not null"`
	Breed           string    `json:"breed" gorm:"not null"`
	Salary          float64   `json:"salary" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CatCreate struct {
	Name            string  `json:"name" binding:"required"`
	YearsExperience int     `json:"years_experience" binding:"required,min=0"`
	Breed           string  `json:"breed" binding:"required"`
	Salary          float64 `json:"salary" binding:"required,min=0"`
}

type CatUpdate struct {
	Salary float64 `json:"salary" binding:"required,min=0"`
}
