package models

import (
	"database/sql/driver"
	"github.com/pkg/errors"
)

// ChanWidth is the channel width of a given frequency. W corresponds to wideband (20kHz) and N
// corresponds to narrowband (12.5kHz).
type ChanWidth int
const (
	W = iota + 1
	N
)

func (c ChanWidth) String() string {
	m := map[ChanWidth]string{
		W: "W",
		N: "N",
	}

	s, ok := m[c]
	if ok {
		return s
	}

	return ""
}

func ParseChanWidth(s string) (ChanWidth, error) {
	m := map[string]ChanWidth{
		"W": W,
		"N": N,
	}

	t, ok := m[s]
	if ok {
		return t, nil
	}

	return 0, errors.Errorf("%s is not a valid chan_width", s)
}

func (c ChanWidth) Value() (driver.Value, error) {
	s := c.String()

	if s != "" {
		return s, nil
	}

	return "", errors.Errorf("%v is not a valid value for chan_width", c)
}

func (c *ChanWidth) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.Errorf("value to scan cannot be cast to a string: %v", value)
	}

	var err error
	*c, err = ParseChanWidth(s)
	if err != nil {
		return errors.WithMessage(err, "failed to scan chan_width")
	}

	return nil
}

func (c *ChanWidth) MarshalJSON() ([]byte, error) {
	s := c.String()

	if s != "" {
		return []byte(s), nil
	}

	return nil, errors.Errorf("%v is not a valid value for chan_width", c)
}

func (c *ChanWidth) UnmarshalJSON(value []byte) error {
	s := string(value)

	var err error
	*c, err = ParseChanWidth(s)
	if err != nil {
		return errors.WithMessage(err, "failed to scan chan_width")
	}

	return nil
}

