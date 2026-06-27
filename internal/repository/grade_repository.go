package repository

import (
	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"github.com/macbook/lumina-backend-go/internal/domain/repositories"
	"gorm.io/gorm"
)

type gormGradeRepository struct {
	db *gorm.DB
}

func NewGradeRepository(db *gorm.DB) repositories.IGradeRepository {
	return &gormGradeRepository{db: db}
}

func (r *gormGradeRepository) FindAll() ([]entities.Grade, error) {
	var grades []entities.Grade
	err := r.db.Order("\"orderIndex\" asc").Find(&grades).Error
	return grades, err
}

func (r *gormGradeRepository) FindByID(id string) (*entities.Grade, error) {
	var grade entities.Grade
	err := r.db.First(&grade, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &grade, nil
}

func (r *gormGradeRepository) Create(grade *entities.Grade) error {
	return r.db.Create(grade).Error
}

func (r *gormGradeRepository) Update(grade *entities.Grade) error {
	return r.db.Save(grade).Error
}

func (r *gormGradeRepository) Delete(id string) error {
	return r.db.Delete(&entities.Grade{}, "id = ?", id).Error
}
