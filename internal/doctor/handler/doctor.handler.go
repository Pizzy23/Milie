package doctor

import (
	_ "millieMind/db"
	inter "millieMind/internal/doctor/interface"
	service "millieMind/internal/doctor/service"
	pInter "millieMind/internal/pacient/interface"
	pService "millieMind/internal/pacient/service"
	sInter "millieMind/internal/skills/interface"
	sService "millieMind/internal/skills/service"
	errors "millieMind/middleware/interfaces/errors"

	"github.com/gin-gonic/gin"
)

// @Summary Cria um novo Doutor
// @Description Cria um novo Doutor no banco de dados
// @Tags Doctor
// @Accept json
// @Produce json
// @Param request body inter.ControllerDoctor true "Dados do Doutor a ser criado"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.OutputDoctor "Doutor criado com sucesso"
// @Failure 406 {object} errors.NotAcceptable "Error"
// @Router /api/create-doctor [post]
func CreateDoctor(c *gin.Context) {
	var doctor inter.ControllerDoctor
	if err := c.BindJSON(&doctor); err != nil {
		c.Set("Response", "Paramets is invalid, need a json")
		c.Status(errors.StatusNotAcceptable)
		return
	}
	service.AddDoctor(c, doctor)
}

// @Summary Buscar um doutor
// @Description Buscar um novo Doutor no banco de dados
// @Tags Doctor
// @Accept json
// @Produce json
// @Param Email header string true "Email do doutor"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.OutputDoctor "Doutor criado com sucesso"
// @Router /api/search-doctor [get]
func SearchDoctor(c *gin.Context) {
	email := c.GetHeader("Email")
	service.GetDoctor(c, email)

}

// @Summary Altera Doutor
// @Description Altera o doutor, tem que me passar oque precisa alterar
// @Tags Doctor
// @Accept json
// @Produce json
// @Param request body inter.ControllerDoctor true "Dados do Doutor a ser criado"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.OutputDoctor "Doutor criado com sucesso"
// @Router /api/modify-doctor [put]
func ModifyDoctor(c *gin.Context) {
	var input inter.ControllerDoctor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Set("Response", "Paramets is invalid, need a json")
		c.Status(errors.StatusNotAcceptable)
		return
	}
	service.ModifyDoctor(c, input)
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

// @Summary Editar Skills
// @Description Mudar o valor da habilidade do paciente
// @Tags Pacient
// @Accept json
// @Produce json
// @Param request body sInter.Skills true "Dados da habilidade"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.OutputDoctor "Habilidade editada com sucesso"
// @Failure 406 {object} errors.NotAcceptable "Error"
// @Router /api/modif-skill [put]
func ModifySkills(c *gin.Context) {
	var input sInter.Skills
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Set("Response", "Paramets is invalid, need a json")
		c.Status(errors.StatusNotAcceptable)
		return
	}
	sService.ModifSkills(c, input)
}
