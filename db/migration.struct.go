package db

import (
	"time"
)

type Login struct {
	ID       uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Email    string    `gorm:"column:email" json:"email"`
	IsLogged bool      `gorm:"column:isLogged;not null" json:"isLogged"`
	CreateAt time.Time `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;not null" json:"update_at"`
}

type Skills struct {
	ID       uint      `gorm:"column:id;primaryKey" json:"id"`
	Type     string    `gorm:"column:type;not null" json:"type"`
	Note     string    `gorm:"column:note;not null" json:"note"`
	CreateAt time.Time `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;not null" json:"update_at"`
}
type SkillsTotal struct {
	ID         uint      `gorm:"column:id;primaryKey" json:"id"`
	SkillsID   uint      `gorm:"column:Skills_idSkills;not null" json:"Skills_idSkills"`
	Average    string    `gorm:"column:avarage;not null" json:"avarage"`
	AllAverage string    `gorm:"column:allAverage;not null" json:"allAverage"`
	CreateAt   time.Time `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt   time.Time `gorm:"column:update_at;not null" json:"update_at"`
}

type Doctor struct {
	ID       uint      `gorm:"column:id;primaryKey" json:"id"`
	Name     string    `gorm:"column:name;not null" json:"name"`
	Email    string    `gorm:"column:email;unique;not null" json:"email"`
	Password string    `gorm:"column:password;not null" json:"password"`
	ZipCode  string    `gorm:"column:zipCode;not null" json:"zipCode"`
	Street   string    `gorm:"column:street;not null" json:"street"`
	District string    `gorm:"column:district;not null" json:"district"`
	City     string    `gorm:"column:city;not null" json:"city"`
	Login    Login     `gorm:"foreignKey:LoginID" json:"login"`
	LoginID  uint      `gorm:"column:Login_idLogin;not null" json:"Login_idLogin"`
	CreateAt time.Time `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;not null" json:"update_at"`
}

type Consults struct {
	ID        uint      `gorm:"column:id;primaryKey;not null" json:"id"`
	Date      string    `gorm:"column:Date;not null" json:"Date"`
	Doctor    Doctor    `gorm:"foreignKey:DoctorID" json:"doctor"`
	DoctorID  uint      `gorm:"column:Doctor_idDoctor;not null" json:"Doctor_idDoctor"`
	Pacient   Pacient   `gorm:"foreignKey:PacientID" json:"pacient"`
	PacientID uint      `gorm:"column:Pacient_idPacient;not null" json:"Pacient_idPacient"`
	CreateAt  time.Time `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;not null" json:"update_at"`
}

type Family struct {
	ID        uint       `gorm:"column:id;primaryKey;not null" json:"id"`
	PacientID uint       `gorm:"column:Pacient_idPacient;not null" json:"Pacient_idPacient"`
	Pacient   *Pacient   `gorm:"foreignKey:PacientID" json:"pacient"`
	Parents   []Parents  `gorm:"foreignKey:FamilyID" json:"parents"`
	Brothers  []Brothers `gorm:"foreignKey:FamilyID" json:"brothers"`
	CreateAt  time.Time  `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt  time.Time  `gorm:"column:update_at;not null" json:"update_at"`
}

type Parents struct {
	ID             uint   `gorm:"column:id;primaryKey;not null" json:"id"`
	FamilyID       uint   `gorm:"column:Family_id;not null" json:"family_id"`
	Relationship   string `gorm:"column:relationShip;not null" json:"relationShip"`
	NeuroDivergent string `gorm:"column:neuroDivergent;not null" json:"neuroDivergent"`
}

type Brothers struct {
	ID             uint   `gorm:"column:id;primaryKey;not null" json:"id"`
	FamilyID       uint   `gorm:"column:Family_id;not null" json:"family_id"`
	Health         string `gorm:"column:healt;not null" json:"health"`
	NeuroDivergent string `gorm:"column:neuroDivergent;not null" json:"neuroDivergent"`
}

type Pacient struct {
	ID             uint      `gorm:"column:id;primaryKey;not null" json:"id"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	Email          string    `gorm:"column:email;unique;not null" json:"email"`
	Phone          string    `gorm:"column:phone;unique;not null" json:"phone"`
	Region         string    `gorm:"column:region;not null" json:"region"`
	Age            int       `gorm:"column:age;not null" json:"age"`
	Gender         string    `gorm:"column:gender;not null" json:"gender"`
	Score          uint      `gorm:"column:score;not null" json:"score"`
	History        []byte    `gorm:"column:content;type:longblob;not null" json:"pdfContent"`
	NeuroDivergent string    `gorm:"column:neuroDivergen;not null" json:"neuroDivergent"`
	LoginID        uint      `gorm:"column:Login_idLogin;not null" json:"login_id"`
	Login          *Login    `gorm:"foreignKey:LoginID" json:"login"`
	SkillsId       uint      `gorm:"many2many:Skills_idSkills;" json:"Skills_idSkills"`
	Skills         *Skills   `gorm:"many2many:patient_skills;" json:"skills"`
	DoctorID       uint      `gorm:"column:Doctor_idDoctor;not null" json:"doctor_id"`
	Doctor         *Doctor   `gorm:"foreignKey:DoctorID" json:"doctor"`
	FamilyID       uint      `gorm:"column:Family_id;not null" json:"family_id"`
	Family         *Family   `gorm:"foreignKey:FamilyID" json:"family"`
	CreateAt       time.Time `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt       time.Time `gorm:"column:update_at;not null" json:"update_at"`
}

type BaseQuestions struct {
	FormsName  string `gorm:"column:Forms_name;not null" json:"forms_name"`
	Categories string `gorm:"column:Categories;not null" json:"categories"`
	Age        string `gorm:"column:Age;not null" json:"age"`
	Question   string `gorm:"column:question;not null" json:"question"`
}

type MakeQuestions struct {
	ID         uint               `gorm:"column:id" json:"id"`
	DoctorID   uint               `gorm:"column:Doctor_idDoctor;not null" json:"Doctor_idDoctor"`
	PacientID  uint               `gorm:"column:Pacient_idPacient;not null" json:"Pacient_idPacient"`
	FormsName  string             `gorm:"column:Forms_name;not null" json:"forms_name"`
	Categories string             `gorm:"column:Categories;not null" json:"categories"`
	Answers    []AnsweredQuestion `gorm:"foreignKey:MakeQuestionsID" json:"answers"`
}

type AnsweredQuestion struct {
	ID              uint   `gorm:"column:id;primaryKey" json:"id"`
	MakeQuestionsID uint   `gorm:"column:MakeQuestions_idMakeQuestions;not null" json:"MakeQuestions_idMakeQuestions"`
	QuestionID      int    `gorm:"column:Question_id;not null" json:"question_id"`
	Question        string `gorm:"column:Questions;not null" json:"question"`
	Answer          bool   `gorm:"column:Answer;not null" json:"answer"`
}

type Orientation struct {
	Text      string   `gorm:"type:TEXT;not null" json:"text"`
	Checkbox  bool     `gorm:"type:BOOLEAN;not null;default:false" json:"check"`
	Pacient   *Pacient `gorm:"foreignKey:PacientID" json:"pacient"`
	PacientID uint     `gorm:"column:Pacient_idPacient;not null" json:"Pacient_idPacient"`
}
