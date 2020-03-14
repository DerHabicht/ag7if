package models

type Service struct {
	BaseModel
	Name string `json:"name"`
	CFRPart string `json:"cfr_part"`
	License bool `json:"license"`
	Bands []Band `json:"bands,omitempty"`
}
