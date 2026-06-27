package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/macbook/lumina-backend-go/internal/service"
)

type GradeHandler struct {
	service *service.GradeService
}

func NewGradeHandler(s *service.GradeService) *GradeHandler {
	return &GradeHandler{service: s}
}

func (h *GradeHandler) GetAllGrades(c *gin.Context) {
	grades, err := h.service.GetAllGrades()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grades)
}

func (h *GradeHandler) GetGradeByID(c *gin.Context) {
	id := c.Param("id")

	grade, err := h.service.GetGradeByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if grade == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "grade not found"})
		return
	}

	c.JSON(http.StatusOK, grade)
}

func (h *GradeHandler) CreateGrade(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		OrderIndex  int    `json:"orderIndex"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade, err := h.service.CreateGrade(input.Name, input.Description, input.OrderIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grade)
}

func (h *GradeHandler) UpdateGrade(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		OrderIndex  int    `json:"orderIndex"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade, err := h.service.UpdateGrade(id, input.Name, input.Description, input.OrderIndex)
	if err != nil {
		if err.Error() == "grade not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, grade)
}

func (h *GradeHandler) DeleteGrade(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteGrade(id)
	if err != nil {
		if err.Error() == "grade not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
