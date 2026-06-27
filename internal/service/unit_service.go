package service

import (
	"errors"
	"strings"

	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"github.com/macbook/lumina-backend-go/internal/domain/repositories"
)

type UnitService struct {
	repo repositories.IUnitRepository
}

func NewUnitService(repo repositories.IUnitRepository) *UnitService {
	return &UnitService{repo: repo}
}

func (s *UnitService) GetAllUnits(courseID *string) ([]entities.Unit, error) {
	return s.repo.FindAll(courseID)
}

func (s *UnitService) GetUnitByID(id string) (*entities.Unit, error) {
	return s.repo.FindByID(id)
}

func (s *UnitService) CreateUnit(title, description string, orderIndex int, courseID string) (*entities.Unit, error) {
	if strings.TrimSpace(title) == "" {
		return nil, errors.New("unit title is required")
	}
	if courseID == "" {
		return nil, errors.New("course ID is required")
	}

	unit := &entities.Unit{
		Title:       title,
		Description: description,
		OrderIndex:  orderIndex,
		CourseID:    courseID,
	}

	err := s.repo.Create(unit)
	return unit, err
}

func (s *UnitService) UpdateUnit(id string, title, description string, orderIndex int) (*entities.Unit, error) {
	existingUnit, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existingUnit == nil {
		return nil, errors.New("unit not found")
	}

	if title != "" {
		existingUnit.Title = title
	}
	existingUnit.Description = description
	existingUnit.OrderIndex = orderIndex

	err = s.repo.Update(existingUnit)
	return existingUnit, err
}

func (s *UnitService) DeleteUnit(id string) error {
	existingUnit, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existingUnit == nil {
		return errors.New("unit not found")
	}

	return s.repo.Delete(id)
}
