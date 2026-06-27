package service

import (
	"errors"
	"strings"

	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"github.com/macbook/lumina-backend-go/internal/domain/repositories"
)

type GradeService struct {
	repo repositories.IGradeRepository
}

func NewGradeService(repo repositories.IGradeRepository) *GradeService {
	return &GradeService{repo: repo}
}

func (s *GradeService) GetAllGrades() ([]entities.Grade, error) {
	return s.repo.FindAll()
}

func (s *GradeService) GetGradeByID(id string) (*entities.Grade, error) {
	return s.repo.FindByID(id)
}

func (s *GradeService) CreateGrade(name, description string, orderIndex int) (*entities.Grade, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("grade name is required")
	}

	grade := &entities.Grade{
		Name:        name,
		Description: description,
		OrderIndex:  orderIndex,
	}

	err := s.repo.Create(grade)
	return grade, err
}

func (s *GradeService) UpdateGrade(id, name, description string, orderIndex int) (*entities.Grade, error) {
	existingGrade, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existingGrade == nil {
		return nil, errors.New("grade not found")
	}

	if name != "" {
		existingGrade.Name = name
	}
	existingGrade.Description = description
	existingGrade.OrderIndex = orderIndex

	err = s.repo.Update(existingGrade)
	return existingGrade, err
}

func (s *GradeService) DeleteGrade(id string) error {
	existingGrade, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existingGrade == nil {
		return errors.New("grade not found")
	}

	return s.repo.Delete(id)
}
