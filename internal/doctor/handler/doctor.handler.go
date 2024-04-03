package doctor

import (
	_ "millieMind/db"
	inter "millieMind/internal/doctor/interface"
	service "millieMind/internal/doctor/service"
	"net/http"

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
// @Router /api/create-doctor [post]
func CreateDoctor(c *gin.Context) {
	var doctor inter.ControllerDoctor
	if err := c.BindJSON(&doctor); err != nil {
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
	var input inter.SearchDoctors
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.GetDoctor(c, input)

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.ModifyDoctor(c, input)
}

func AddPacient(c *gin.Context)   {}
func ModifySkills(c *gin.Context) {}
