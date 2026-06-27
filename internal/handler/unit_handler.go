package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/macbook/lumina-backend-go/internal/service"
)

type UnitHandler struct {
	service *service.UnitService
}

func NewUnitHandler(s *service.UnitService) *UnitHandler {
	return &UnitHandler{service: s}
}

func (h *UnitHandler) GetAllUnits(c *gin.Context) {
	courseIDStr := c.Query("courseId")
	var courseID *string

	if courseIDStr != "" {
		courseID = &courseIDStr
	}

	units, err := h.service.GetAllUnits(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, units)
}

func (h *UnitHandler) GetUnitByID(c *gin.Context) {
	id := c.Param("id")

	unit, err := h.service.GetUnitByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if unit == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "unit not found"})
		return
	}

	c.JSON(http.StatusOK, unit)
}

func (h *UnitHandler) CreateUnit(c *gin.Context) {
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		OrderIndex  int    `json:"orderIndex"`
		CourseID    string `json:"courseId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unit, err := h.service.CreateUnit(input.Title, input.Description, input.OrderIndex, input.CourseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, unit)
}

func (h *UnitHandler) UpdateUnit(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		OrderIndex  int    `json:"orderIndex"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unit, err := h.service.UpdateUnit(id, input.Title, input.Description, input.OrderIndex)
	if err != nil {
		if err.Error() == "unit not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, unit)
}

func (h *UnitHandler) DeleteUnit(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteUnit(id)
	if err != nil {
		if err.Error() == "unit not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
