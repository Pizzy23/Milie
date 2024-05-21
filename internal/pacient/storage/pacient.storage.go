package storage

import (
	"millieMind/db"

	"gorm.io/gorm"
)

func ConsultPacient(q *gorm.DB, data string) (*db.Pacient, error) {
	var pacient db.Pacient
	if err := q.First(&pacient, data).Error; err != nil {
		return nil, err
	}
	return &pacient, nil
}

func ConsultOrientation(q *gorm.DB, data string) (*db.Orientation, error) {
	var orientation db.Orientation
	if err := q.First(&orientation, data).Error; err != nil {
		return nil, err
	}
	return &orientation, nil
}

func CreatePacient(q *gorm.DB, data db.Pacient) (*db.Pacient, error) {
	if err := q.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func ConsultDoctor(q *gorm.DB, email string) (*db.Doctor, error) {
	var doctor db.Doctor
	var login db.Login
	if err := q.Where("email = ?", email).First(&doctor).Error; err != nil {
		return nil, err
	}
	if err := q.Where("email = ?", email).First(&login).Error; err != nil {
		return nil, err
	}
	doctor.Login = login
	return &doctor, nil
}

func CreateOrientations(q *gorm.DB, orientations []db.Orientation) error {
	return q.Create(&orientations).Error
}
