package postgres

import "time"

type Contact struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Last_name    string
	Phone_Number string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
