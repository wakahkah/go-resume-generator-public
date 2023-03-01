package model

type PersonalInfo struct {
	PhotoUrl    string      `json:"photoUrl"`
	Name        string      `json:"name"`
	Position    string      `json:"position"`
	Objective   string      `json:"objective"`
	ContactInfo ContactInfo `json:"contact"`
}

type ContactInfo struct {
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	LinkedlnUrl string `json:"linkedlnUrl"`
	WebsiteUrl  string `json:"websiteUrl"`
	GithubUrl   string `json:"githubUrl"`
}
