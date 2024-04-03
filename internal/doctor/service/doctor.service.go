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
		login, err := login.CreateLogin(db.Repo, data.Email)
		if err != nil {
			c.Set("Response", "Internal Server error")
			c.Status(http.StatusInternalServerError)
		}
		encryptedPassword, err := util.Encrypt(data.Password)
		if err != nil {
			c.Set("Response", "Error encrypting password")
			c.Status(http.StatusInternalServerError)
			return
		}
		encryptedEmail, err := util.Encrypt(data.Password)
		if err != nil {
			c.Set("Response", "Error encrypting email")
			c.Status(http.StatusInternalServerError)
			return
		}

		newDoctor := db.Doctor{
			Name:     data.Name,
			Email:    encryptedEmail,
			Password: encryptedPassword,
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
	}
}

func GetDoctor(c *gin.Context, data inter.SearchDoctors) {
	doctor, err := storage.ConsultDoctor(db.Repo, data.Email)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}
	decryptedEmail, err := util.Decrypt(doctor.Email)
	if err != nil {
		c.Set("Response", "Error decrypting email")
		c.Status(http.StatusInternalServerError)
		return
	}
	doctor.Email = decryptedEmail

	c.Set("Response", doctor)
	c.Status(http.StatusOK)
}

func ModifyDoctor(c *gin.Context, data inter.ControllerDoctor) {
	var existingDoctor db.Doctor
	if err := db.Repo.First(&existingDoctor, data.Email).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Médico não encontrado"})
		return
	}

	encryptedPassword, err := util.Encrypt(data.Password)
	if err != nil {
		c.Set("Response", "Error encrypting password")
		c.Status(http.StatusInternalServerError)
		return
	}
	encryptedEmail, err := util.Encrypt(data.Password)
	if err != nil {
		c.Set("Response", "Error encrypting email")
		c.Status(http.StatusInternalServerError)
		return
	}

	existingDoctor.Name = data.Name
	existingDoctor.Email = encryptedEmail
	existingDoctor.Password = encryptedPassword
	existingDoctor.ZipCode = data.ZipCode
	existingDoctor.Street = data.Street
	existingDoctor.District = data.District
	existingDoctor.City = data.City
	existingDoctor.UpdateAt = time.Now()

	doctor, err := storage.ModifyDoctor(db.Repo, &existingDoctor)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
	}
	c.Set("Response", doctor)
	c.Status(http.StatusOK)
}
