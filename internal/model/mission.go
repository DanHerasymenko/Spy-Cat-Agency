package model

import (
	"time"
)

type Mission struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CatID     uint      `json:"cat_id" gorm:"not null"`
	Cat       Cat       `json:"cat" gorm:"foreignKey:CatID"`
	Targets   []Target  `json:"targets" gorm:"foreignKey:MissionID"`
	Completed bool      `json:"completed" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MissionCreate struct {
	Name    string         `json:"name" binding:"required"`
	CatID   uint           `json:"cat_id" binding:"required"`
	Targets []TargetCreate `json:"targets" binding:"required,min=1,max=3,dive"`
}

type MissionUpdate struct {
	Completed bool `json:"completed"`
}

type CatAssign struct {
	CatID uint `json:"cat_id" binding:"required"`
}
