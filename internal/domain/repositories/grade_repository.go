package repositories

import (
	"github.com/macbook/lumina-backend-go/internal/domain/entities"
)

type IGradeRepository interface {
	FindAll() ([]entities.Grade, error)
	FindByID(id string) (*entities.Grade, error)
	Create(grade *entities.Grade) error
	Update(grade *entities.Grade) error
	Delete(id string) error
}
