package postgres

import "time"

type Message struct {
	ID        uint `gorm:"primaryKey"`
	Message   string
	Sender    string
	Is_group  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
