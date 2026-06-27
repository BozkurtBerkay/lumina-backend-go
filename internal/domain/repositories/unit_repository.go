package repositories

import (
	"github.com/macbook/lumina-backend-go/internal/domain/entities"
)

type IUnitRepository interface {
	FindAll(courseID *string) ([]entities.Unit, error)
	FindByID(id string) (*entities.Unit, error)
	Create(unit *entities.Unit) error
	Update(unit *entities.Unit) error
	Delete(id string) error
}
