package service

import (
	"errors"
	"strings"

	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"github.com/macbook/lumina-backend-go/internal/domain/repositories"
)

type CourseService struct {
	repo repositories.ICourseRepository
}

func NewCourseService(repo repositories.ICourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) GetAllCourses(gradeID *string) ([]entities.Course, error) {
	return s.repo.FindAll(gradeID)
}

func (s *CourseService) GetCourseByID(id string) (*entities.Course, error) {
	return s.repo.FindByID(id)
}

func (s *CourseService) CreateCourse(title, description string, gradeID string) (*entities.Course, error) {
	if strings.TrimSpace(title) == "" {
		return nil, errors.New("course title is required")
	}
	if gradeID == "" {
		return nil, errors.New("grade ID is required")
	}

	course := &entities.Course{
		Title:       title,
		Description: description,
		GradeID:     gradeID,
	}

	err := s.repo.Create(course)
	return course, err
}

func (s *CourseService) UpdateCourse(id string, title, description string) (*entities.Course, error) {
	existingCourse, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existingCourse == nil {
		return nil, errors.New("course not found")
	}

	if title != "" {
		existingCourse.Title = title
	}
	existingCourse.Description = description

	err = s.repo.Update(existingCourse)
	return existingCourse, err
}

func (s *CourseService) DeleteCourse(id string) error {
	existingCourse, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existingCourse == nil {
		return errors.New("course not found")
	}

	return s.repo.Delete(id)
}
