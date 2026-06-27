package repository

import (
	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"github.com/macbook/lumina-backend-go/internal/domain/repositories"
	"gorm.io/gorm"
)

type gormQuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) repositories.IQuestionRepository {
	return &gormQuestionRepository{db: db}
}

func (r *gormQuestionRepository) FindAll(unitID *string) ([]entities.Question, error) {
	var questions []entities.Question
	query := r.db.Order("\"orderIndex\" asc")
	if unitID != nil {
		query = query.Where("\"unitId\" = ?", *unitID)
	}
	err := query.Find(&questions).Error
	return questions, err
}

func (r *gormQuestionRepository) FindByID(id string) (*entities.Question, error) {
	var question entities.Question
	err := r.db.First(&question, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &question, nil
}

func (r *gormQuestionRepository) Create(question *entities.Question) error {
	return r.db.Create(question).Error
}

func (r *gormQuestionRepository) Update(question *entities.Question) error {
	return r.db.Save(question).Error
}

func (r *gormQuestionRepository) Delete(id string) error {
	return r.db.Delete(&entities.Question{}, "id = ?", id).Error
}
