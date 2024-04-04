package storage

import (
	"millieMind/db"

	"gorm.io/gorm"
)

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

func CreateDoctor(q *gorm.DB, data *db.Doctor) (*db.Doctor, error) {
	var login db.Login
	if err := q.Create(data).Error; err != nil {
		return nil, err
	}
	if err := q.Where("email = ?", data.Email).First(&login).Error; err != nil {
		return nil, err
	}
	data.Login = login
	return data, nil
}

func ModifyDoctor(q *gorm.DB, data *db.Doctor) (*db.Doctor, error) {
	if err := db.Repo.Save(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
