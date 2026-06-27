package repositories

import (
	"github.com/macbook/lumina-backend-go/internal/domain/entities"
)

type ICourseRepository interface {
	FindAll(gradeID *string) ([]entities.Course, error)
	FindByID(id string) (*entities.Course, error)
	Create(course *entities.Course) error
	Update(course *entities.Course) error
	Delete(id string) error
}
