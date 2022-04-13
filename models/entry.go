package models

import (
	"time"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	UserID       uint       `json:"user_id"`
	User         User       `json:"user"`
	ClockInTime  time.Time  `json:"clock_in" gorm:"not null"`
	ClockOutTime *time.Time `json:"clock_out"`
}

func (e Entry) IsOnGoing() bool {
	return e.ClockOutTime == nil
}
