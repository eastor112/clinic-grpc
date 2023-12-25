package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID        int64
	FirstName string
	LastName  string
	Phone     string
	SSN       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Patient) TableName() string {
	return "patients"
}
