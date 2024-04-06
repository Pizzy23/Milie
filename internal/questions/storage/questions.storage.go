package storage

import (
	"millieMind/db"

	"gorm.io/gorm"
)

func AddQuestion(q *gorm.DB, data db.BaseQuestions) (*db.BaseQuestions, error) {
	if err := q.Create(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
