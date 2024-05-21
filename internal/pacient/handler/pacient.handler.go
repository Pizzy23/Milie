package pacient

import (
	pInter "millieMind/internal/pacient/interface"
	pService "millieMind/internal/pacient/service"
	errors "millieMind/middleware/interfaces/errors"

	"github.com/gin-gonic/gin"
)

// @Summary Recupera um paciente
// @Description Busca um paciente pelo e-mail no banco de dados
// @Tags Pacient
// @Accept json
// @Produce json
// @Param Email header string true  "E-mail do paciente"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} pInter.OutputPacient "Paciente recuperado com sucesso"
// @Failure 404 {object} errors.NotFound "Paciente não encontrado"
// @Router /api/pull-patient [get]
func PullPacient(c *gin.Context) {
	email := c.GetHeader("Email")
	pService.PullPacient(c, email)

}

// @Summary Adiciona questoes
// @Description Adiciona questoes ao banco
// @Tags Pacient
// @Accept json
// @Produce json
// @Param request body pInter.CotrollerQuestionsMarker true "Todas as questoes."
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} pInter.OutputPacient "Questions input in db"
// @Failure 404 {object} errors.NotFound "Questions not inputing"
// @Router /api/question-mark [post]
func QuestionsMark(c *gin.Context) {
	var patient pInter.CotrollerQuestionsMarker
	if err := c.BindJSON(&patient); err != nil {
		c.Set("Response", "Invalid parameters, need a JSON with email")
		c.Status(errors.StatusNotAcceptable)
		return
	}
}

// @Summary Puxar infos paciente
// @Description Puxa todas as informaçoes do usuario.
// @Tags Pacient
// @Accept json
// @Produce json
// @Param Email header string true  "E-mail do paciente"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} pInter.OutputPacient "Paciente recuperado com sucesso"
// @Failure 404 {object} errors.NotFound "Paciente não encontrado"
// @Router /api/pull-profile [get]
func PullProfile(c *gin.Context) {
	email := c.GetHeader("Email")
	pService.PullProfile(c, email)
}

// @Summary Puxar orientação
// @Description Puxa orientações do banco
// @Tags Pacient
// @Accept json
// @Produce json
// @Param Email header string true  "E-mail do paciente"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} pInter.OutputPacient "Questions input in db"
// @Failure 404 {object} errors.NotFound "Questions not inputing"
// @Router /api/pull-orientation [get]
func PullOrientation(c *gin.Context) {
	email := c.GetHeader("Email")
	pService.PullOrientation(c, email)
}

// @Summary Adicionar orientação
// @Description Adicionar orientações ao paciente
// @Tags Pacient
// @Accept json
// @Produce json
// @Param request body pInter.OrientationController true "Todas as questoes."
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} pInter.OrientationOutPut "Questions input in db"
// @Failure 404 {object} errors.NotFound "Questions not inputing"
// @Router /api/creat-orientation [post]
func CreatOrientation(c *gin.Context) {
	var patient pInter.CotrollerQuestionsMarker
	if err := c.BindJSON(&patient); err != nil {
		c.Set("Response", "Invalid parameters, need a JSON with email")
		c.Status(errors.StatusNotAcceptable)
		return
	}
}

// @Summary Pegar os dados do paciente
// @Description Pegar os dados do paciente atraves do nome completo
// @Tags Pacient
// @Accept json
// @Produce json
// @Param Name header string true  "Nome Completo Do Paciente"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} pInter.OutputPacient "Paciente recuperado com sucesso"
// @Failure 404 {object} errors.NotFound "Paciente não encontrado"
// @Router /api/pacient-name [get]
func GetByName(c *gin.Context) {
	email := c.GetHeader("Email")
	pService.PullByName(c, email)
}

// @Summary Criar Paciente
// @Description Criar um novo paciente
// @Tags Pacient
// @Accept json
// @Produce json
// @Param request body pInter.InputPacient true "Dados do paciente a ser criado"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} pInter.Pacient "token make:"
// @Router /api/add-pacient [post]
func AddPacient(c *gin.Context) {
	var input pInter.InputPacient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Set("Response", "Paramets is invalid, need a json")
		c.Status(errors.StatusNotAcceptable)
		return
	}
	pService.CreatePacient(c, input)
}
