package login

import (
	"millieMind/db"
	"time"

	"gorm.io/gorm"
)

func CreateLogin(q *gorm.DB, email string) (*db.Login, error) {
	CreateAt := time.Now()
	UpdateAt := time.Now()
	newLogin := &db.Login{
		Email:    email,
		IsLogged: false,
		CreateAt: CreateAt,
		UpdateAt: UpdateAt,
	}
	if err := q.Create(newLogin).Error; err != nil {
		return nil, err
	}
	return newLogin, nil
}
