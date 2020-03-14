package models

type Channel struct {
	// Database fields
	BaseModel
	Name string `json:"name"`
	RxFrequency float32 `json:"rx_frequency"`
	RxWidth ChanWidth `json:"rx_width"`
	RxToneType ToneType `json:"rx_tone_type"`
	RxTone *float32 `json:"rx_tone"`
	TxFrequency *float32 `json:"tx_frequency"`
	TxWidth *ChanWidth `json:"tx_width"`
	TxToneType *ToneType `json:"tx_tone_type"`
	TxTone *float32 `json:"tx_tone"`
	Mode ChanMode `json:"chan_mode"`
	Remarks string `json:"remarks"`
	IsPublic bool `json:"is_public"`
	UserID uint `json:"owner_id"`
	BandID uint `json:"band_id"`
}

// BeforeSave validates tone and Tx channel parameters.
func (c *Channel) BeforeSave() error {
	var v ValidationErrors

	if err := ValidateToneType(&c.RxToneType, c.RxTone); err != nil {
		v = append(v, err.Error())
	}

	if err := ValidateToneType(c.TxToneType, c.TxTone); err != nil {
		v = append(v, err.Error())
	}

	if c.TxFrequency == nil {
		if c.TxWidth != nil || c.TxToneType != nil || c.TxTone != nil {
			v = append(
				v,
				"TxFrequency was nil thus all Tx parameters were expected to be nil, some Tx params were not",
				)

		}
	}

	return v
}
