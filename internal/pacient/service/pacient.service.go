package pService

import (
	pInter "millieMind/internal/pacient/interface"
	"millieMind/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePacient(c *gin.Context, input pInter.InputPacient) {}

func QuestionsMark(c *gin.Context) {}

func CreatOrientation(c *gin.Context) {}

func PullProfile(c *gin.Context, email string) {}

func PullOrientation(c *gin.Context, email string) {}

func PullByName(c *gin.Context, email string) {}

func PullPacient(c *gin.Context, email string) {
	emailValid := util.IsValidEmail(email)
	if emailValid {

	} else {
		c.Set("Response", "E-mail is invalid")
		c.Status(http.StatusNotAcceptable)
	}
}
