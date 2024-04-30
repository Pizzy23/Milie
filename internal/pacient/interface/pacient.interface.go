package pInter

import (
	"time"
)

type SearchPacient struct {
	Email string `json:"email"`
}

type Pacient struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Region         string    `json:"region"`
	Age            int       `json:"age"`
	Gender         string    `json:"gender"`
	Score          int       `json:"score"`
	History        []byte    `json:"pdfContent"`
	NeuroDivergent string    `json:"neuroDivergent"`
	LoginID        int       `json:"Login_idLogin"`
	SkillsID       int       `json:"Skills_idSkills"`
	DoctorID       int       `json:"Doctor_idDoctor"`
	FamilyID       int       `json:"Family_idFamily"`
	CreateAt       time.Time `json:"create_at"`
	UpdateAt       time.Time `json:"update_at"`
}

type InputPacient struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Region         string `json:"region"`
	Age            int    `json:"age"`
	Gender         string `json:"gender"`
	Score          int    `json:"score"`
	History        []byte `json:"pdfContent"`
	NeuroDivergent string `json:"neuroDivergent"`
	Family         Family `json:"family"`
}
type Family struct {
	Parents  []Parents  `json:"parents"`
	Brothers []Brothers `json:"brothers"`
}
type Brothers struct {
	Health         string `json:"health"`
	NeuroDivergent string `json:"neuroDivergent"`
}
type Parents struct {
	Relationship   string `json:"relationShip"`
	NeuroDivergent string `json:"neuroDivergent"`
}

type ControllerPullPacient struct {
	Email string `json:"email"`
}

type OutputPacient struct {
	Pacient
}

type QuestionsMark struct {
	FormsName  string `gorm:"column:Forms_name;not null" json:"forms_name"`
	Categories string `gorm:"column:Categories;not null" json:"categories"`
	Question   string `gorm:"column:question;not null" json:"question"`
	Age        string `gorm:"column:Age;not null" json:"age"`
	Answer     string `gorm:"column:Answer;not null" json:"answer"`
}
