package model

import "time"

type Certificate struct {
	Name         string    `json:"name"`
	Date         time.Time `json:"date"`
	ReferenceUrl string    `json:"referenceUrl"`
	Provider     string    `json:"provider"`
	YearStr      string
}
