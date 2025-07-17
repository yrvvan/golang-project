package models

import "gorm.io/gorm"
import "time"

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	Profile   Profile        `json:"profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Leave     LeaveBalance   `json:"leave" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type GetUser struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Profile   Profile        `json:"profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Leave     LeaveBalance   `json:"leave" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
