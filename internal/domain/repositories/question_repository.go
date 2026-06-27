package repositories

import (
	"github.com/macbook/lumina-backend-go/internal/domain/entities"
)

type IQuestionRepository interface {
	FindAll(unitID *string) ([]entities.Question, error)
	FindByID(id string) (*entities.Question, error)
	Create(question *entities.Question) error
	Update(question *entities.Question) error
	Delete(id string) error
}
