package model

import (
	"time"
)

type Target struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	MissionID uint      `json:"mission_id" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	Country   string    `json:"country" gorm:"not null"`
	Notes     string    `json:"notes"`
	Completed bool      `json:"completed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TargetCreate struct {
	Name    string `json:"name" binding:"required"`
	Country string `json:"country" binding:"required"`
	Notes   string `json:"notes"`
}

type TargetUpdate struct {
	Notes     string `json:"notes"`
	Completed bool   `json:"completed"`
}
