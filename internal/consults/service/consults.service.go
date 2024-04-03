package consults

import (
	"millieMind/db"
	inter "millieMind/internal/consults/interface"
	consults "millieMind/internal/consults/storage"
	doctor "millieMind/internal/doctor/storage"
	pacient "millieMind/internal/pacient/storage"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateConsult(c *gin.Context, data inter.InputConsult) {
	doctorRes, errD := doctor.ConsultDoctor(db.Repo, data.EmailDoctor)
	if errD != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messge": "Doctor dont found", "err": errD})
	}
	pacientRes, errP := pacient.ConsultPacient(db.Repo, data.EmailPacient)
	if errP != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messge": "Pacient dont found", "err": errP})
	}
	var dataConsult = &inter.InsertInDB{
		Date:      data.Date,
		DoctorID:  doctorRes.ID,
		PacientID: pacientRes.ID,
		CreateAt:  time.Now(),
	}
	consult, errC := consults.CreateConsult(db.Repo, *dataConsult)
	if errC != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messge": "Doctor dont found", "err": errD})
	} else {
		c.JSON(http.StatusCreated, gin.H{"messge": "our consult create", "data": consult})
	}
}
