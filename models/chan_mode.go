package models

import (
	"database/sql/driver"
	"github.com/pkg/errors"
)

// ChanMode refers to the operating mode on this channel. A represents analog modes, D represents digital modes, and M
// represents mixed-modes (when analog and digital modes are both present in a duplex or trunked configuration).
type ChanMode int

const (
	A ChanMode = iota + 1
	D
	M
)

func (c ChanMode) String() string {
	m := map[ChanMode]string{
		A: "A",
		D: "D",
		M: "M",
	}

	s, ok := m[c]
	if ok {
		return s
	}

	return ""
}

func ParseChanMode(s string) (ChanMode, error) {
	m := map[string]ChanMode{
		"A": A,
		"D": D,
		"M": M,
	}

	t, ok := m[s]
	if ok {
		return t, nil
	}

	return 0, errors.Errorf("%s is not a valid chan_mode", s)
}

func (c ChanMode) Value() (driver.Value, error) {
	s := c.String()

	if s != "" {
		return s, nil
	}

	return "", errors.Errorf("%v is not a valid value for chan_mode", c)
}

func (c *ChanMode) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.Errorf("value to scan cannot be cast to a string: %v", value)
	}

	var err error
	*c, err = ParseChanMode(s)
	if err != nil {
		return errors.WithMessage(err, "failed to scan chan_mode")
	}

	return nil
}

func (c *ChanMode) MarshalJSON() ([]byte, error) {
	s := c.String()

	if s != "" {
		return []byte(s), nil
	}

	return nil, errors.Errorf("%v is not a valid value for chan_mode", c)
}

func (c *ChanMode) UnmarshalJSON(value []byte) error {
	s := string(value)

	var err error
	*c, err = ParseChanMode(s)
	if err != nil {
		return errors.WithMessage(err, "failed to scan chan_mode")
	}

	return nil
}
