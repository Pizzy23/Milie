package inter

import "time"

type InputDoctor struct {
	Name     string    `gorm:"column:name" json:"name"`
	Email    string    `gorm:"column:email" json:"email"`
	Password string    `gorm:"column:password" json:"password"`
	ZipCode  string    `gorm:"column:zipCode" json:"zipCode"`
	Street   string    `gorm:"column:street" json:"street"`
	District string    `gorm:"column:district" json:"district"`
	City     string    `gorm:"column:city" json:"city"`
	LoginID  int       `gorm:"column:Login_idLogin" json:"Login_idLogin"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at" json:"update_at"`
}

type SearchDoctors struct {
	Email string `gorm:"column:email" json:"email"`
}

type OutputDoctor struct {
	ID       uint      `gorm:"column:id;primaryKey" json:"id"`
	Name     string    `gorm:"column:name" json:"name"`
	Email    string    `gorm:"column:email" json:"email"`
	Password string    `gorm:"column:password" json:"password"`
	ZipCode  string    `gorm:"column:zipCode" json:"zipCode"`
	Street   string    `gorm:"column:street" json:"street"`
	District string    `gorm:"column:district" json:"district"`
	City     string    `gorm:"column:city" json:"city"`
	LoginID  int       `gorm:"column:Login_idLogin" json:"Login_idLogin"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at" json:"update_at"`
}

type ControllerDoctor struct {
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	ZipCode  string `gorm:"column:zipCode" json:"zipCode"`
	Street   string `gorm:"column:street" json:"street"`
	District string `gorm:"column:district" json:"district"`
	City     string `gorm:"column:city" json:"city"`
}
