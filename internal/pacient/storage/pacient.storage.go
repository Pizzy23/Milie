package pacient

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
