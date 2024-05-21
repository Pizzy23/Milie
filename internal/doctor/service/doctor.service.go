package service

import (
	"millieMind/db"
	inter "millieMind/internal/doctor/interface"
	storage "millieMind/internal/doctor/storage"
	login "millieMind/internal/login/storage"
	"net/http"
	"time"

	util "millieMind/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddDoctor(c *gin.Context, data inter.ControllerDoctor) {
	valid := util.InputValidator{
		Password: data.Password,
		Email:    data.Email,
		Phone:    "0130121321432",
		ZipCode:  data.ZipCode,
	}
	validations := util.UltimateValidator(valid)
	if validations == "" {

		_, errDB := storage.ConsultDoctor(db.Repo, data.Email)

		if errDB == gorm.ErrRecordNotFound {
			login, err := login.CreateLogin(db.Repo, data.Email)
			if err != nil {
				c.Set("Response", "Internal Server error")
				c.Status(http.StatusInternalServerError)
			}
			data = encryptData(c, data)
			newDoctor := db.Doctor{
				Name:     data.Name,
				Email:    data.Email,
				Password: data.Password,
				ZipCode:  data.ZipCode,
				Street:   data.Street,
				District: data.District,
				City:     data.City,
			}
			newDoctor.LoginID = login.ID
			newDoctor.CreateAt = time.Now()
			newDoctor.UpdateAt = time.Now()
			doctor, err := storage.CreateDoctor(db.Repo, &newDoctor)
			if err != nil {
				c.Set("Response", "Internal Server error")
				c.Status(http.StatusInternalServerError)
			}
			c.Set("Response", doctor)
			c.Status(http.StatusOK)
		} else {
			if errDB != nil {
				c.Set("Response", errDB)
				c.Status(http.StatusInternalServerError)
			}
			c.Set("Response", "Email already registered")
			c.Status(http.StatusConflict)

		}
	} else {
		c.Set("Response", validations)
		c.Status(http.StatusOK)
	}
}

func GetDoctor(c *gin.Context, data string) {

	doctor, err := storage.ConsultDoctor(db.Repo, data)
	if err != nil {
		c.Set("Response", "Don't found email")
		c.Status(http.StatusInternalServerError)
		return
	}
	decryptPassword, err := util.Decrypt(doctor.Password)

	if err != nil {
		c.Set("Response", "Error decrypt email")
		c.Status(http.StatusInternalServerError)

	}
	doctor.Password = decryptPassword

	decryptZipCode, err := util.Decrypt(doctor.ZipCode)
	if err != nil {
		c.Set("Response", "Error decrypt ZipCode")
		c.Status(http.StatusInternalServerError)

	}
	doctor.ZipCode = decryptZipCode
	decryptStreet, err := util.Decrypt(doctor.Street)
	if err != nil {
		c.Set("Response", "Error decrypt Street")
		c.Status(http.StatusInternalServerError)

	}
	doctor.Street = decryptStreet
	decryptDistrict, err := util.Decrypt(doctor.District)
	if err != nil {
		c.Set("Response", "Error decrypt District")
		c.Status(http.StatusInternalServerError)

	}
	doctor.District = decryptDistrict
	decryptCity, err := util.Decrypt(doctor.City)
	if err != nil {
		c.Set("Response", "Error decrypt City")
		c.Status(http.StatusInternalServerError)

	}
	doctor.City = decryptCity
	c.Set("Response", doctor)
	c.Status(http.StatusOK)

}

func ModifyDoctor(c *gin.Context, data inter.ControllerDoctor) {
	existingDoctor, err := storage.ConsultDoctor(db.Repo, data.Email)
	if err != nil {
		c.Set("Response", "Don't found email")
		c.Status(http.StatusInternalServerError)
		return
	}
	data = encryptData(c, data)
	switch {
	case data.Name != "":
		existingDoctor.Name = data.Name
		fallthrough
	case data.Email != "":
		existingDoctor.Email = data.Email
		fallthrough
	case data.ZipCode != "":
		existingDoctor.ZipCode = data.ZipCode
		fallthrough
	case data.Street != "":
		existingDoctor.Street = data.Street
		fallthrough
	case data.District != "":
		existingDoctor.District = data.District
		fallthrough
	case data.City != "":
		existingDoctor.City = data.City
	}

	existingDoctor.UpdateAt = time.Now()

	doctor, err := storage.ModifyDoctor(db.Repo, existingDoctor)

	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
	}
	c.Set("Response", doctor)
	c.Status(http.StatusOK)
}

func encryptData(c *gin.Context, data inter.ControllerDoctor) inter.ControllerDoctor {
	encryptedPassword, err := util.Encrypt(data.Password)

	if err != nil {
		c.Set("Response", "Error encrypting Password")
		c.Status(http.StatusInternalServerError)

	}

	encryptedZipCode, err := util.Encrypt(data.ZipCode)
	if err != nil {
		c.Set("Response", "Error encrypting ZipCode")
		c.Status(http.StatusInternalServerError)

	}

	encryptedStreet, err := util.Encrypt(data.Street)
	if err != nil {
		c.Set("Response", "Error encrypting Street")
		c.Status(http.StatusInternalServerError)

	}

	encryptedDistrict, err := util.Encrypt(data.District)
	if err != nil {
		c.Set("Response", "Error encrypting District")
		c.Status(http.StatusInternalServerError)

	}

	encryptedCity, err := util.Encrypt(data.City)
	if err != nil {
		c.Set("Response", "Error encrypting City")
		c.Status(http.StatusInternalServerError)

	}
	data = inter.ControllerDoctor{
		Name:     data.Name,
		Email:    data.Email,
		Password: encryptedPassword,
		ZipCode:  encryptedZipCode,
		Street:   encryptedStreet,
		District: encryptedDistrict,
		City:     encryptedCity,
	}
	return data
}
