package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type ValidationErrors []string

func (v ValidationErrors) Error() string {
	verrs, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Errorf("marshalling of validation errors failed: %v", err))
	}

	return string(verrs)
}

type ModelBase struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
