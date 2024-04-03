package inter

import (
	"time"
)

type SearchPacient struct {
	Email string `gorm:"column:email" json:"email"`
}

type OutputPacient struct {
	ID             uint      `gorm:"column:id;primaryKey" json:"id"`
	Name           string    `gorm:"column:name" json:"name"`
	Email          string    `gorm:"column:email" json:"email"`
	Phone          string    `gorm:"column:phone" json:"phone"`
	Region         string    `gorm:"column:region" json:"region"`
	Age            int       `gorm:"column:age" json:"age"`
	Gender         string    `gorm:"column:gender" json:"gender"`
	Score          int       `gorm:"column:score" json:"score"`
	History        []byte    `gorm:"column:content;type:longblob" json:"pdfContent"`
	NeuroDivergent string    `gorm:"column:neuroDivergent" json:"neuroDivergent"`
	LoginID        int       `gorm:"column:Login_idLogin" json:"Login_idLogin"`
	SkillsID       int       `gorm:"column:Skills_idSkills" json:"Skills_idSkills"`
	DoctorID       int       `gorm:"column:Doctor_idDoctor" json:"Doctor_idDoctor"`
	FamilyID       int       `gorm:"column:Family_idFamily" json:"Family_idFamily"`
	CreateAt       time.Time `gorm:"column:create_at" json:"create_at"`
	UpdateAt       time.Time `gorm:"column:update_at" json:"update_at"`
}
