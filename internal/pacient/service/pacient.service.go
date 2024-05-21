package pService

import (
	"millieMind/db"
	login "millieMind/internal/login/storage"
	pInter "millieMind/internal/pacient/interface"
	storage "millieMind/internal/pacient/storage"
	"millieMind/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePacient(c *gin.Context, input pInter.InputPacient) {
	family := db.Family{
		Parents:  make([]db.Parents, len(input.Family.Parents)),
		Brothers: make([]db.Brothers, len(input.Family.Brothers)),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	doctor, err := storage.ConsultDoctor(db.Repo, input.Doctor)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}
	if doctor == nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}

	loginMake, err := login.CreateLogin(db.Repo, input.Email)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}
	if loginMake == nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}

	for i, parent := range input.Family.Parents {
		family.Parents[i] = db.Parents{
			Relationship:   parent.Relationship,
			NeuroDivergent: parent.NeuroDivergent,
		}
	}

	for i, brother := range input.Family.Brothers {
		family.Brothers[i] = db.Brothers{
			Health:         brother.Health,
			NeuroDivergent: brother.NeuroDivergent,
		}
	}

	pacient := db.Pacient{
		Name:           input.Name,
		Email:          input.Email,
		Phone:          input.Phone,
		Region:         input.Region,
		Age:            input.Age,
		Gender:         input.Gender,
		Score:          uint(input.Score),
		History:        input.History,
		NeuroDivergent: input.NeuroDivergent,
		Family:         family,
		Doctor:         doctor,
		Login:          loginMake,
		CreateAt:       time.Now(),
		UpdateAt:       time.Now(),
	}

	newP, err := storage.CreatePacient(db.Repo, pacient)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Set("Response", newP)
	c.Status(http.StatusOK)
}

func QuestionsMark(c *gin.Context, input pInter.CotrollerQuestionsMarker) {
	pacient, err := storage.ConsultPacient(db.Repo, input.PacientEmail)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}
	doctor, err := storage.ConsultDoctor(db.Repo, input.DoctorEmail)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}

	answers := make([]db.AnsweredQuestion, len(input.Questions))
	for i, q := range input.Questions {
		answers[i] = db.AnsweredQuestion{
			QuestionID: q.QuestionID,
			Question:   q.Question,
			Answer:     q.Answer,
		}
	}

	inputDb := db.MakeQuestions{
		DoctorID:   doctor.ID,
		PacientID:  pacient.ID,
		FormsName:  input.FormsName,
		Categories: input.Categories,
		Answers:    answers,
	}

	if err := db.Repo.Create(&inputDb).Error; err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Set("Response", inputDb)
	c.Status(http.StatusOK)
}

func CreateOrientation(c *gin.Context, input pInter.OrientationController) {
	pacient, err := storage.ConsultPacient(db.Repo, input.PacientEmail)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}
	if pacient == nil {
		c.Set("Response", "Paciente não encontrado")
		c.Status(http.StatusNotFound)
		return
	}

	orientations := make([]db.Orientation, len(input.Precautions))
	for i, q := range input.Precautions {
		orientations[i] = db.Orientation{
			Text:      q.Orietation,
			Checkbox:  q.CheckBox,
			Pacient:   pacient,
			PacientID: pacient.ID,
		}
	}

	if err := storage.CreateOrientations(db.Repo, orientations); err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Set("Response", "Orientations created successfully")
	c.Status(http.StatusOK)
}

func PullProfile(c *gin.Context, email string) {
	res, err := storage.ConsultPacient(db.Repo, email)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func PullOrientation(c *gin.Context, email string) {
	res, err := storage.ConsultPacient(db.Repo, email)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func PullByName(c *gin.Context, name string) {
	res, err := storage.ConsultPacient(db.Repo, name)
	if err != nil {
		c.Set("Response", "Internal Server error")
		c.Status(http.StatusInternalServerError)
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func PullPacient(c *gin.Context, email string) {
	emailValid := util.IsValidEmail(email)
	if emailValid {
		res, err := storage.ConsultPacient(db.Repo, email)
		if err != nil {
			c.Set("Response", "Internal Server error")
			c.Status(http.StatusInternalServerError)
		}
		c.Set("Response", res)
		c.Status(http.StatusOK)
	} else {
		c.Set("Response", "E-mail is invalid")
		c.Status(http.StatusNotAcceptable)
	}
}
