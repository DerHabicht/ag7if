package models

type User struct {
	BaseModel
	Username string    `json:"username"`
	Zones    []Zone    `json:"zones,omitempty"`
	Channels []Channel `json:"channels,omitempty"`
}
