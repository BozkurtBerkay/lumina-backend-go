package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"github.com/macbook/lumina-backend-go/internal/service"
	"gorm.io/datatypes"
)

type QuestionHandler struct {
	service *service.QuestionService
}

func NewQuestionHandler(s *service.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: s}
}

func (h *QuestionHandler) GetAllQuestions(c *gin.Context) {
	unitIDStr := c.Query("unitId")
	var unitID *string

	if unitIDStr != "" {
		unitID = &unitIDStr
	}

	questions, err := h.service.GetAllQuestions(unitID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (h *QuestionHandler) GetQuestionByID(c *gin.Context) {
	id := c.Param("id")

	question, err := h.service.GetQuestionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if question == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}

	c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) CreateQuestion(c *gin.Context) {
	var input struct {
		Content       string                `json:"content" binding:"required"`
		Type          entities.QuestionType `json:"type" binding:"required"`
		Options       datatypes.JSON        `json:"options"`
		CorrectAnswer string                `json:"correctAnswer"`
		ImageURL      string                `json:"imageUrl"`
		ImageAlt      string                `json:"imageAlt"`
		OrderIndex    int                   `json:"orderIndex"`
		UnitID        string                `json:"unitId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question, err := h.service.CreateQuestion(input.Content, input.Type, input.Options, input.CorrectAnswer, input.ImageURL, input.ImageAlt, input.OrderIndex, input.UnitID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, question)
}

func (h *QuestionHandler) UpdateQuestion(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Content       string                `json:"content"`
		Type          entities.QuestionType `json:"type"`
		Options       datatypes.JSON        `json:"options"`
		CorrectAnswer string                `json:"correctAnswer"`
		ImageURL      string                `json:"imageUrl"`
		ImageAlt      string                `json:"imageAlt"`
		OrderIndex    int                   `json:"orderIndex"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question, err := h.service.UpdateQuestion(id, input.Content, input.Type, input.Options, input.CorrectAnswer, input.ImageURL, input.ImageAlt, input.OrderIndex)
	if err != nil {
		if err.Error() == "question not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteQuestion(id)
	if err != nil {
		if err.Error() == "question not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
