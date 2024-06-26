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
	Doctor         string `json:"doctor"`
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

type CotrollerQuestionsMarker struct {
	DoctorEmail  string          `json:"doctor_Email"`
	PacientEmail string          `json:"pacient_Email"`
	FormsName    string          `json:"forms_name"`
	Categories   string          `json:"categories"`
	Questions    []QuestionsMark `json:"questions"`
}

type QuestionsMark struct {
	QuestionID int    `json:"question_id"`
	Question   string `json:"question"`
	Answer     bool   `json:"answer"`
}

type OrientationController struct {
	PacientEmail string                   `json:"pacient_Email"`
	Precautions  []OrientationPrecautions `json:"precautions"`
}

type OrientationPrecautions struct {
	Orietation string `json:"orietation"`
	CheckBox   bool   `json:"checkBox"`
}

type OrientationOutPut struct {
	PacientId   string                   `json:"pacientId"`
	Precautions []OrientationPrecautions `json:"precautions"`
}

type ControllerGetName struct {
	Name string `json:"name"`
}
