package repository

import (
	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"github.com/macbook/lumina-backend-go/internal/domain/repositories"
	"gorm.io/gorm"
)

type gormUnitRepository struct {
	db *gorm.DB
}

func NewUnitRepository(db *gorm.DB) repositories.IUnitRepository {
	return &gormUnitRepository{db: db}
}

func (r *gormUnitRepository) FindAll(courseID *string) ([]entities.Unit, error) {
	var units []entities.Unit
	query := r.db.Order("\"orderIndex\" asc")
	if courseID != nil {
		query = query.Where("\"courseId\" = ?", *courseID)
	}
	err := query.Find(&units).Error
	return units, err
}

func (r *gormUnitRepository) FindByID(id string) (*entities.Unit, error) {
	var unit entities.Unit
	err := r.db.First(&unit, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &unit, nil
}

func (r *gormUnitRepository) Create(unit *entities.Unit) error {
	return r.db.Create(unit).Error
}

func (r *gormUnitRepository) Update(unit *entities.Unit) error {
	return r.db.Save(unit).Error
}

func (r *gormUnitRepository) Delete(id string) error {
	return r.db.Delete(&entities.Unit{}, "id = ?", id).Error
}
