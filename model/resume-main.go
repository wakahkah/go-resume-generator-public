package model

type ResumeData struct {
	PersonalInfo    PersonalInfo      `json:"personalInfo"`
	Skills          []SkillGroup      `json:"skills"`
	WorkExperience  []WorkExperience  `json:"workExperience"`
	Certificates    []Certificate     `json:"certificates"`
	PersonalProject []PersonalProject `json:"personalProject"`
	Awards          []Award           `json:"awards"`
	Education       []Education       `json:"education"`
	Interest        []Interest        `json:"interest"`
}
