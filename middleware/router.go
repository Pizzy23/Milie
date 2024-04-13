package middleware

import (
	_ "millieMind/docs"
	consults "millieMind/internal/consults/handler"
	doctor "millieMind/internal/doctor/handler"
	pacient "millieMind/internal/pacient/handler"
	questions "millieMind/internal/questions/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Milie Mind
// @version 1.0
// @description API Milie Mind
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/add-questions", questions.AddQuestions)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(ResponseHandler())
	//Use response, but not Token
	r.GET("/token", generateTokenHandler)

	auth := r.Group("/api")
	auth.Use(authMiddleware)
	//Response and token service
	auth.PUT("/modify-doctor", doctor.ModifyDoctor)
	auth.PUT("/modify-skill", doctor.ModifySkills)
	auth.PUT("/questions", pacient.QuestionsMark)

	auth.POST("/create-orientation", pacient.CreatOrientation)
	auth.POST("/questions", pacient.QuestionsPulledAge)
	auth.POST("/consults", consults.AddConsults)
	auth.POST("/test-token", testeToken)
	auth.POST("/create-doctor", doctor.CreateDoctor)
	auth.POST("/add-pacient", doctor.AddPacient)

	auth.GET("/activity", pacient.PullActivity)
	auth.GET("/orientation", pacient.PullOrientation)
	auth.GET("/pacient", pacient.PullPacient)
	auth.GET("/profile", pacient.PullProfile)
	auth.GET("/search-doctor", doctor.SearchDoctor)
	auth.GET("/consults", consults.CheckConsults)
	auth.GET("/pull-questions", questions.PullQuestions)

	auth.DELETE("/consults", consults.RemoveConsults)

	return r
}
