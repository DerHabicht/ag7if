package models

import (
	"database/sql/driver"
	"fmt"
	"github.com/pkg/errors"
	"math"
)

// ToneType represents the tone needed to open the squelch of a radio's receiver. CSQ represents carrier squelch (i.e.
// no tone is needed), CTCSS represents an analog tone at a given frequency, and DCS represents a 9-bit digital sequence
// to be sent at the beginning of each transmission (represented as an octal number).
type ToneType int
const (
	CSQ ToneType = iota + 1
	CTCSS
	DCS
)

func (t ToneType) String() string {
	m := map[ToneType]string{
		CSQ: "CSQ",
		CTCSS: "CTCSS",
		DCS: "DCS",
	}

	s, ok := m[t]
	if ok {
		return s
	}

	return ""
}

func ParseToneType(s string) (ToneType, error) {
	m := map[string]ToneType{
		"CSQ": CSQ,
		"CTCSS": CTCSS,
		"DCS": DCS,
	}

	t, ok := m[s]
	if ok {
		return t, nil
	}

	return 0, errors.Errorf("%s is not a valid tone_type", s)
}

func (t ToneType) Value() (driver.Value, error) {
	s := t.String()

	if s != "" {
		return s, nil
	}

	return "", errors.Errorf("%v is not a valid value for tone_type", t)
}

func (t *ToneType) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.Errorf("value to scan cannot be cast to a string: %v", value)
	}

	var err error
	*t, err = ParseToneType(s)
	if err != nil {
		return errors.WithMessage(err, "failed to scan tone_type")
	}

	return nil
}

func (t *ToneType) MarshalJSON() ([]byte, error) {
	s := t.String()

	if s != "" {
		return []byte(s), nil
	}

	return nil, errors.Errorf("%v is not a valid value for tone_type", t)
}

func (t *ToneType) UnmarshalJSON(value []byte) error {
	s := string(value)

	var err error
	*t, err = ParseToneType(s)
	if err != nil {
		return errors.WithMessage(err, "failed to scan tone_type")
	}

	return nil
}

func ValidateToneType(t *ToneType, v *float32) error {
	if t == nil {
		if v != nil {
			 return fmt.Errorf("ToneType is nil thus a nil tone was expected, tone was %v instead", *v)
		}
	} else {
		if v == nil {
			return fmt.Errorf("ToneType was not nil thus a non-nil tone was expected, got nil instead")
		} else if *t == CSQ {
			return fmt.Errorf("ToneType was CSQ thus a nil tone was expected, tone was %v instead", *v)
		} else if *t == DCS && (math.Floor(float64(*v)) != float64(*v)) {
			return fmt.Errorf("ToneType was DCS thus an integer tone was expected, tone was %v instead", *v)
		}
	}

	return nil
}
