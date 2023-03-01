package model

type SkillGroup struct {
	GroupName string  `json:"groupName"`
	SkillData []Skill `json:"skillData"`
	SkillStr  string
}

type Skill struct {
	Name  string `json:"name"`
	Level int8   `json:"level"` //1-5: 1 is low and 5 is high
}
