package sInter

type Skills struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	Note  string `json:"note"`
}

type SkillsTotal struct {
	SkillsID   uint   `json:"Skills_idSkills"`
	Average    string `json:"avarage"`
	AllAverage string `json:"allAverage"`
}
