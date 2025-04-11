package models

import "gorm.io/gorm"
import "time"

type LeaveBalance struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserID      uint           `json:"user_id"`
	AnnualLeave int            `json:"annual_leave"`
	SickLeave   int            `json:"sick_leave"`
	UnpaidLeave int            `json:"unpaid_leave"`
}
