package storage

import (
	"millieMind/db"

	"gorm.io/gorm"
)

func ConsultDoctor(q *gorm.DB, data string) (*db.Doctor, error) {
	var doctor db.Doctor
	if err := q.First(&doctor, data).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func CreateDoctor(q *gorm.DB, data *db.Doctor) (*db.Doctor, error) {
	if err := q.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func ModifyDoctor(q *gorm.DB, data *db.Doctor) (*db.Doctor, error) {
	if err := db.Repo.Save(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
