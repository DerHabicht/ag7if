package models

type Band struct {
	// Database fields
	BaseModelWithUUID
	Name           string  `json:"name"`
	LowerFrequency float32 `json:"lower_frequency"`
	UpperFrequency float32 `json:"upper_frequency"`
	ServiceID      uint    `json:"service_id"`
}
