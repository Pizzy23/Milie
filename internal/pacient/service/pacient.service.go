package pService

import (
	pInter "millieMind/internal/pacient/interface"
	"millieMind/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PullPacient(c *gin.Context, input pInter.ControllerPullPacient) {
	emailValid := util.IsValidEmail(input.Email)
	if emailValid {

	} else {
		c.Set("Response", "E-mail is invalid")
		c.Status(http.StatusNotAcceptable)
	}
}

func QuestionsMark(c *gin.Context) {}

func PullActivity(c *gin.Context) {}

func PullProfile(c *gin.Context) {}

func PullOrientation(c *gin.Context) {}

func CreatOrientation(c *gin.Context) {}

func CreatActivity(c *gin.Context) {}
