package marshallers

import (
	"time"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

type MarshalledEntry struct {
	ID           uint       `json:"id"`
	UserID       uint       `json:"user_id"`
	UserName     string     `json:"username"`
	ClockInTime  time.Time  `json:"clock_in"`
	ClockOutTime *time.Time `json:"clock_out"`
}

func MarshalEntry(e models.Entry) MarshalledEntry {
	return MarshalledEntry{
		ID:           e.ID,
		UserID:       e.UserID,
		UserName:     e.User.Username,
		ClockInTime:  e.ClockInTime,
		ClockOutTime: e.ClockOutTime,
	}
}

func MarshalEntries(entries []models.Entry) []MarshalledEntry {
	en := make([]MarshalledEntry, 0)

	for _, e := range entries {
		en = append(en, MarshalEntry(e))
	}

	return en
}
