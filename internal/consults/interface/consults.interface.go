package inter

import "time"

type InputConsult struct {
	Date         string `gorm:"column:Date" json:"Date"`
	EmailDoctor  string ` json:"EmailDoctor"`
	EmailPacient string `json:"EmailPacient"`
}

type OutputConsult struct {
	ID        uint      `gorm:"column:id;primaryKey" json:"id"`
	Date      string    `gorm:"column:Date" json:"Date"`
	DoctorID  uint      `gorm:"column:Doctor_idDoctor" json:"Doctor_idDoctor"`
	PacientID uint      `gorm:"column:Pacient_idPacient" json:"Pacient_idPacient"`
	CreateAt  time.Time `gorm:"column:create_at" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at" json:"update_at"`
}

type CheckConsult struct {
	Date        string `gorm:"column:Date" json:"Date"`
	EmailDoctor string ` json:"EmailDoctor"`
}

type InsertInDB struct {
	Date      string    `gorm:"column:Date" json:"Date"`
	DoctorID  uint      `gorm:"column:Doctor_idDoctor" json:"Doctor_idDoctor"`
	PacientID uint      `gorm:"column:Pacient_idPacient" json:"Pacient_idPacient"`
	CreateAt  time.Time `gorm:"column:create_at" json:"create_at"`
}
