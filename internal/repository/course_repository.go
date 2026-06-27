package repository

import (
	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"github.com/macbook/lumina-backend-go/internal/domain/repositories"
	"gorm.io/gorm"
)

type gormCourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) repositories.ICourseRepository {
	return &gormCourseRepository{db: db}
}

func (r *gormCourseRepository) FindAll(gradeID *string) ([]entities.Course, error) {
	var courses []entities.Course
	query := r.db.Order("\"createdAt\" desc")
	if gradeID != nil {
		query = query.Where("\"gradeId\" = ?", *gradeID)
	}
	err := query.Find(&courses).Error
	return courses, err
}

func (r *gormCourseRepository) FindByID(id string) (*entities.Course, error) {
	var course entities.Course
	err := r.db.First(&course, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &course, nil
}


func (r *gormCourseRepository) Create(course *entities.Course) error {
	return r.db.Create(course).Error
}

func (r *gormCourseRepository) Update(course *entities.Course) error {
	return r.db.Save(course).Error
}

func (r *gormCourseRepository) Delete(id string) error {
	return r.db.Delete(&entities.Course{}, "id = ?", id).Error
}

