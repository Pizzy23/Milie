package consults

import (
	"millieMind/db"
	inter "millieMind/internal/consults/interface"

	"gorm.io/gorm"
)

func CreateConsult(q *gorm.DB, data inter.InsertInDB) (*db.Consults, error) {
	newConsult := &db.Consults{
		Date:      data.Date,
		DoctorID:  data.DoctorID,
		PacientID: data.PacientID,
	}
	if err := q.Create(newConsult).Error; err != nil {
		return nil, err
	}
	return newConsult, nil
}

func CheckConsult(q *gorm.DB, data inter.CheckConsult) (*db.Consults, error) {
	var consult db.Consults
	if err := q.Find(&consult, data).Error; err != nil {
		return nil, err
	}
	return &consult, nil
}
