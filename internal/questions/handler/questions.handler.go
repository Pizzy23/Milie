package questions

import (
	inter "millieMind/internal/questions/interface"
	"millieMind/internal/questions/service"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.InputQuestionsDB(c, input)
}

func PullQuestions(c *gin.Context) {}

func AnswerQuestions(c *gin.Context) {}
