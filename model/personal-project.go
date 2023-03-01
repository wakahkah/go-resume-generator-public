package model

import "time"

type PersonalProject struct {
	Name         string    `json:"name"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	IsPresent    bool      `json:"isPresent"`
	Description  string    `json:"description"`
	ReferenceUrl string    `json:"referenceUrl"`
	DateRangeStr string
}
