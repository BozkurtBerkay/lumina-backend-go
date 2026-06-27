package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type QuestionType string

const (
	MultipleChoice QuestionType = "MULTIPLE_CHOICE"
	TrueFalse      QuestionType = "TRUE_FALSE"
	OpenEnded      QuestionType = "OPEN_ENDED"
)

type Grade struct {
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	Name        string    `gorm:"not null;column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	OrderIndex  int       `gorm:"default:0;column:orderIndex" json:"orderIndex"`
	CreatedAt   time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	Courses     []Course  `gorm:"foreignKey:GradeID;constraint:OnDelete:CASCADE" json:"courses,omitempty"`
}

type Course struct {
	ID          string    `gorm:"primaryKey;column:id" json:"id"`
	Title       string    `gorm:"not null;column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	GradeID     string    `gorm:"not null;column:gradeId" json:"gradeId"`
	CreatedAt   time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	Grade       Grade     `gorm:"foreignKey:GradeID" json:"grade,omitempty"`
	Units       []Unit    `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE" json:"units,omitempty"`
}

type Unit struct {
	ID          string     `gorm:"primaryKey;column:id" json:"id"`
	Title       string     `gorm:"not null;column:title" json:"title"`
	Description string     `gorm:"column:description" json:"description"`
	OrderIndex  int        `gorm:"default:0;column:orderIndex" json:"orderIndex"`
	CourseID    string     `gorm:"not null;column:courseId" json:"courseId"`
	CreatedAt   time.Time  `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"column:updatedAt" json:"updatedAt"`
	Course      Course     `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	Questions   []Question `gorm:"foreignKey:UnitID;constraint:OnDelete:CASCADE" json:"questions,omitempty"`
}

type Question struct {
	ID            string         `gorm:"primaryKey;column:id" json:"id"`
	Content       string         `gorm:"not null;column:content" json:"content"`
	Type          QuestionType   `gorm:"type:text;default:'MULTIPLE_CHOICE';column:type" json:"type"`
	Options       datatypes.JSON `gorm:"column:options" json:"options"`
	CorrectAnswer string         `gorm:"column:correctAnswer" json:"correctAnswer"`
	ImageURL      string         `gorm:"column:imageUrl" json:"imageUrl"`
	ImageAlt      string         `gorm:"column:imageAlt" json:"imageAlt"`
	OrderIndex    int            `gorm:"default:0;column:orderIndex" json:"orderIndex"`
	UnitID        string         `gorm:"not null;column:unitId" json:"unitId"`
	CreatedAt     time.Time      `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"column:updatedAt" json:"updatedAt"`
	Unit          Unit           `gorm:"foreignKey:UnitID" json:"unit,omitempty"`
}



func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == "" {
		q.ID = uuid.New().String()
	}
	return
}

func (q *Question) TableName() string {
	return "Question"
}

func (g *Grade) BeforeCreate(tx *gorm.DB) (err error) {
	if g.ID == "" {
		g.ID = uuid.New().String()
	}
	return
}

func (g *Grade) TableName() string {
	return "Grade"
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}

func (c *Course) TableName() string {
	return "Course"
}

func (u *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}

func (u *Unit) TableName() string {
	return "Unit"
}


