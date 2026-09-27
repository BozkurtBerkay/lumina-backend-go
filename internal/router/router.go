package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/macbook/lumina-backend-go/internal/handler"
)

func SetupRouter(
	gradeHandler *handler.GradeHandler,
	courseHandler *handler.CourseHandler,
	unitHandler *handler.UnitHandler,
	questionHandler *handler.QuestionHandler,
) *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(cors.Default()) // Basic CORS

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Lumina Go API is running smoothly."})
	})

	// System status check endpoint
	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service": "lumina-backend-go",
			"status":  "active",
			"version": "1.0.0",
		})
	})

	api := r.Group("/api")
	{
		// Grades
		api.GET("/grades", gradeHandler.GetAllGrades)
		api.GET("/grades/:id", gradeHandler.GetGradeByID)
		api.POST("/grades", gradeHandler.CreateGrade)
		api.PUT("/grades/:id", gradeHandler.UpdateGrade)
		api.DELETE("/grades/:id", gradeHandler.DeleteGrade)

		// Courses
		api.GET("/courses", courseHandler.GetAllCourses)
		api.GET("/courses/:id", courseHandler.GetCourseByID)
		api.POST("/courses", courseHandler.CreateCourse)
		api.PUT("/courses/:id", courseHandler.UpdateCourse)
		api.DELETE("/courses/:id", courseHandler.DeleteCourse)

		// Units
		api.GET("/units", unitHandler.GetAllUnits)
		api.GET("/units/:id", unitHandler.GetUnitByID)
		api.POST("/units", unitHandler.CreateUnit)
		api.PUT("/units/:id", unitHandler.UpdateUnit)
		api.DELETE("/units/:id", unitHandler.DeleteUnit)

		// Questions
		api.GET("/questions", questionHandler.GetAllQuestions)
		api.GET("/questions/:id", questionHandler.GetQuestionByID)
		api.POST("/questions", questionHandler.CreateQuestion)
		api.PUT("/questions/:id", questionHandler.UpdateQuestion)
		api.DELETE("/questions/:id", questionHandler.DeleteQuestion)
	}

	return r
}
