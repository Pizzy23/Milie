package login

import (
	"millieMind/db"

	"gorm.io/gorm"
)

func CreateLogin(q *gorm.DB, email string) (*db.Login, error) {
	newLogin := &db.Login{
		Email:    email,
		IsLogged: false,
	}
	if err := q.Create(newLogin).Error; err != nil {
		return nil, err
	}
	return newLogin, nil
}
