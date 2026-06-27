package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/macbook/lumina-backend-go/internal/handler"
	"github.com/macbook/lumina-backend-go/internal/repository"
	"github.com/macbook/lumina-backend-go/internal/router"
	"github.com/macbook/lumina-backend-go/internal/service"
	"github.com/macbook/lumina-backend-go/pkg/database"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Connect to database
	database.ConnectDB()

	// Initialize Repositories
	gradeRepo := repository.NewGradeRepository(database.DB)
	courseRepo := repository.NewCourseRepository(database.DB)
	unitRepo := repository.NewUnitRepository(database.DB)
	questionRepo := repository.NewQuestionRepository(database.DB)

	// Initialize Services
	gradeService := service.NewGradeService(gradeRepo)
	courseService := service.NewCourseService(courseRepo)
	unitService := service.NewUnitService(unitRepo)
	questionService := service.NewQuestionService(questionRepo)

	// Initialize Handlers
	gradeHandler := handler.NewGradeHandler(gradeService)
	courseHandler := handler.NewCourseHandler(courseService)
	unitHandler := handler.NewUnitHandler(unitService)
	questionHandler := handler.NewQuestionHandler(questionService)

	// Setup Router
	r := router.SetupRouter(gradeHandler, courseHandler, unitHandler, questionHandler)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server is starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
