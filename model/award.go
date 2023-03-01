package model

import "time"

type Award struct {
	Name        string    `json:"name"`
	Competition string    `json:"competition"`
	Date        time.Time `json:"date"`
	YearStr     string
}
