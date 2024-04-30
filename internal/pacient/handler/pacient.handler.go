package pacient

import (
	pInter "millieMind/internal/pacient/interface"
	errors "millieMind/middleware/interfaces/errors"

	"github.com/gin-gonic/gin"
)

// @Summary Recupera um paciente
// @Description Busca um paciente pelo e-mail no banco de dados
// @Tags Patient
// @Accept json
// @Produce json
// @Param request body pInter.ControllerPullPacient true "E-mail do paciente"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} pInter.OutputPacient "Paciente recuperado com sucesso"
// @Failure 404 {object} errors.NotFound "Paciente não encontrado"
// @Router /api/pull-patient [post]
func PullPacient(c *gin.Context) {
	var patient pInter.ControllerPullPacient
	if err := c.BindJSON(&patient); err != nil {
		c.Set("Response", "Invalid parameters, need a JSON with email")
		c.Status(errors.StatusNotAcceptable)
		return
	}

}

// @Summary Recupera um paciente
// @Description Busca um paciente pelo e-mail no banco de dados
// @Tags Patient
// @Accept json
// @Produce json
// @Param request body pInter.QuestionsMark true "Informaçoes do pacientes"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} pInter.OutputPacient "Questions input in db"
// @Failure 404 {object} errors.NotFound "Questions not inputing"
// @Router /api/question-mark [post]
func QuestionsMark(c *gin.Context) {}

func PullProfile(c *gin.Context) {}

func PullOrientation(c *gin.Context) {}

func CreatOrientation(c *gin.Context) {}
