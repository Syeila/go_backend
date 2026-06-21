package dao

import "time"

type RefreshToken struct {
	ID         int
	UserID     int
	Token      string
	DeviceName string
	IPAddress  string
	UserAgent  string
	ExpiredAt  time.Time
	CreatedAt  time.Time
}
