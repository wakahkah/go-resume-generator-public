package model

import (
	"html/template"
	"time"
)

type WorkExperience struct {
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate,omitempty"`
	IsPresent    bool      `json:"isPresent"`
	Position     string    `json:"posistion"`
	Company      string    `json:"company"`
	CompanyUrl   string    `json:"companyUrl"`
	Location     string    `json:"location"`
	Achieved     []string  `json:"achieved"`
	AchievedHtml []template.HTML
	Skills       []string `json:"skills"`
	ReferenceUrl string   `json:"referenceUrl"`
	Hide         bool     `json:"hide"`
	DateRangeStr string
}
