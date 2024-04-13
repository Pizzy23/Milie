package response

import (
	base "millieMind/middleware/interfaces"
	"time"
)

type OutputDoctor struct {
	ID       uint      `gorm:"column:id;primaryKey" json:"id"`
	Name     string    `gorm:"column:name" json:"name"`
	Email    string    `gorm:"column:email" json:"email"`
	Password string    `gorm:"column:password" json:"password"`
	ZipCode  string    `gorm:"column:zipCode" json:"zipCode"`
	Street   string    `gorm:"column:street" json:"street"`
	District string    `gorm:"column:district" json:"district"`
	City     string    `gorm:"column:city" json:"city"`
	LoginID  int       `gorm:"column:Login_idLogin" json:"Login_idLogin"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at" json:"update_at"`
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

type BaseQuestions struct {
	ID         uint   `gorm:"column:id" json:"id"`
	FormsName  string `gorm:"column:Forms_name;not null" json:"forms_name"`
	Categories string `gorm:"column:Categories;not null" json:"categories"`
	Age        string `gorm:"column:Age;not null" json:"age"`
	Question   string `gorm:"column:question;not null" json:"question"`
}

type ResBaseQuestions struct {
	Response interface{}  `json:"Response"`
	BaseReq  base.BaseReq `json:"BaseReq"`
}
