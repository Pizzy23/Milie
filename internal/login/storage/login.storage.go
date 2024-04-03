package login

import (
	"millieMind/db"
	"time"

	"gorm.io/gorm"
)

func CreateLogin(q *gorm.DB, email string) (*db.Login, error) {
	newLogin := &db.Login{
		Email:    email,
		IsLogged: false,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	if err := q.Create(newLogin).Error; err != nil {
		return nil, err
	}
	return newLogin, nil
}
