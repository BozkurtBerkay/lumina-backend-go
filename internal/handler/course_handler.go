package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/macbook/lumina-backend-go/internal/service"
)

type CourseHandler struct {
	service *service.CourseService
}

func NewCourseHandler(s *service.CourseService) *CourseHandler {
	return &CourseHandler{service: s}
}

func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	gradeIDStr := c.Query("gradeId")
	var gradeID *string

	if gradeIDStr != "" {
		gradeID = &gradeIDStr
	}

	courses, err := h.service.GetAllCourses(gradeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id := c.Param("id")

	course, err := h.service.GetCourseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if course == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var input struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		GradeID     string `json:"gradeId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.service.CreateCourse(input.Title, input.Description, input.GradeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, course)
}

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.service.UpdateCourse(id, input.Title, input.Description)
	if err != nil {
		if err.Error() == "course not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteCourse(id)
	if err != nil {
		if err.Error() == "course not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
