package model

import "time"

type Education struct {
	UniversityName string    `json:"universityName"`
	Major          string    `json:"major"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
	IsPresent      bool      `json:"isPresent"`
	DateRangeStr   string
}
