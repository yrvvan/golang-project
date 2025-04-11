package models

import "gorm.io/gorm"
import "time"

type Profile struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserID     uint           `json:"user_id"`
	NationalID string         `json:"national_id"` // e.g., NIK or SIM
	BirthPlace string         `json:"birth_place"`
	BirthDate  string         `json:"birth_date"` // e.g., YYYY-MM-DD
	Address    string         `json:"address"`
}
