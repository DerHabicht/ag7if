package models

type User struct {
	BaseModel
	Username string `json:"username"`
	Auth0ID string `json:"auth0_id"`
	Zones []Zone `json:"zones,omitempty"`
	Channels []Channel `json:"channels,omitempty"`
}
