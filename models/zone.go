package models

type Zone struct {
	// Database fields
	BaseModelWithUUID
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
	UserID      uint   `json:"owner_id"`

	// Association objects
	Channels []Channel `json:"channels,omitempty"`
}
