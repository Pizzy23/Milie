package storage

import (
	"millieMind/db"
	inter "millieMind/internal/questions/interface"

	"gorm.io/gorm"
)

func AddQuestion(q *gorm.DB, data db.BaseQuestions) (*db.BaseQuestions, error) {
	if err := q.Create(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func SearchQuestions(q *gorm.DB, data inter.SearchQuestions) ([]db.BaseQuestions, error) {
	var questions []db.BaseQuestions
	if err := q.Where("Forms_name = ? AND Categories = ? AND Age = ?", data.FormsName, data.Categories, data.Age).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}
