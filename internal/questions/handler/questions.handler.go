package questions

import (
	inter "millieMind/internal/questions/interface"
	"millieMind/internal/questions/service"
	"millieMind/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Adicionar as questoes
// @Description Um novo metodo de adicionar as questoes no back
// @Tags Questions
// @Accept json
// @Produce json
// @Param request body inter.Resp true "Dados do Doutor a ser criado"
// @Success 200 {object} inter.OutputDoctor "Doutor criado com sucesso"
// @Router /add-questions [post]
func AddQuestions(c *gin.Context) {
	var input inter.Resp
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Set("Response", "Paramets is invalid, need a json")
		c.Status(http.StatusNotAcceptable)
		return
	}
	service.InputQuestionsDB(c, input)
}

// @Summary Puxar as questoes
// @Description Puxar as questoes no back
// @Tags Questions
// @Accept json
// @Produce json
// @Param EnumFormsName header string true "Formularios : " Enums(Portage)
// @Param EnumCategories header string true "Categorias : " Enums(Portage - Desenvolvimento Motor, Portage - Auto cuidados, Portage - Socializacao, Portage - Linguagem, Portage - Cognicao)
// @Param EnumAges header string true "Age Rate : " Enums(0 - 1, 1 - 2, 2 - 3, 3 - 4, 4 - 5, 5 - 6)
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.OutputDoctor "Doutor criado com sucesso"
// @Failure 406 {object} string "Error"
// @Router /api/pull-questions [get]
func PullQuestions(c *gin.Context) {

	formsName := c.GetHeader("EnumFormsName")
	categories := c.GetHeader("EnumCategories")
	age := c.GetHeader("EnumAges")
	valid := util.EnumQuestions(formsName, categories, age)
	if valid != "" {
		c.Set("Response", valid)
		c.Status(http.StatusNotAcceptable)
		return
	}
	data := inter.SearchQuestions{
		FormsName:  formsName,
		Categories: categories,
		Age:        age,
	}
	service.GetQuestions(c, data)

}

func AnswerQuestions(c *gin.Context) {}
