package dto

type SessionRequest struct {
	UserID     int    `json:"user_id"`
	DeviceName string `json:"device_name"`
	IPAddress  string `json:"ip_address"`
	UserAgent  string `json:"user_agent"`
}
