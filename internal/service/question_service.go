package service

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"github.com/macbook/lumina-backend-go/internal/domain/repositories"
	"gorm.io/datatypes"
)


type QuestionService struct {
	repo repositories.IQuestionRepository
}

func NewQuestionService(repo repositories.IQuestionRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) GetAllQuestions(unitID *string) ([]entities.Question, error) {
	return s.repo.FindAll(unitID)
}

func (s *QuestionService) GetQuestionByID(id string) (*entities.Question, error) {
	return s.repo.FindByID(id)
}

func (s *QuestionService) CreateQuestion(content string, qType entities.QuestionType, options datatypes.JSON, correctAnswer, imageUrl, imageAlt string, orderIndex int, unitID string) (*entities.Question, error) {
	if strings.TrimSpace(content) == "" {
		return nil, errors.New("question content is required")
	}
	if unitID == "" {
		return nil, errors.New("unit ID is required")
	}

	// Format validation for multi-choice questions
	if qType == entities.MultipleChoice {
		var optsMap map[string]string
		if err := json.Unmarshal(options, &optsMap); err != nil {
			return nil, errors.New("options must be a valid JSON object (e.g. {\"A\": \"...\", \"B\": \"...\"})")
		}

		validAnswers := map[string]bool{"A": true, "B": true, "C": true, "D": true}
		if !validAnswers[correctAnswer] {
			return nil, errors.New("correctAnswer must be A, B, C, or D")
		}

		if _, ok := optsMap[correctAnswer]; !ok {
			return nil, errors.New("correctAnswer key must exist in options")
		}
	}

	question := &entities.Question{
		Content:       content,
		Type:          qType,
		Options:       options,
		CorrectAnswer: correctAnswer,
		ImageURL:      imageUrl,
		ImageAlt:      imageAlt,
		OrderIndex:    orderIndex,
		UnitID:        unitID,
	}

	err := s.repo.Create(question)
	return question, err
}


func (s *QuestionService) UpdateQuestion(id string, content string, qType entities.QuestionType, options datatypes.JSON, correctAnswer, imageUrl, imageAlt string, orderIndex int) (*entities.Question, error) {
	existingQuestion, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existingQuestion == nil {
		return nil, errors.New("question not found")
	}

	if content != "" {
		existingQuestion.Content = content
	}

	newType := existingQuestion.Type
	if qType != "" {
		newType = qType
	}

	newOptions := existingQuestion.Options
	if options != nil {
		newOptions = options
	}

	newCorrectAnswer := existingQuestion.CorrectAnswer
	if correctAnswer != "" {
		newCorrectAnswer = correctAnswer
	}

	// Validation for updated values
	if newType == entities.MultipleChoice {
		var optsMap map[string]string
		if err := json.Unmarshal(newOptions, &optsMap); err != nil {
			return nil, errors.New("options must be a valid JSON object")
		}

		validAnswers := map[string]bool{"A": true, "B": true, "C": true, "D": true}
		if !validAnswers[newCorrectAnswer] {
			return nil, errors.New("correctAnswer must be A, B, C, or D")
		}

		if _, ok := optsMap[newCorrectAnswer]; !ok {
			return nil, errors.New("correctAnswer key must exist in options")
		}
	}

	existingQuestion.Type = newType
	existingQuestion.Options = newOptions
	existingQuestion.CorrectAnswer = newCorrectAnswer
	existingQuestion.ImageURL = imageUrl
	existingQuestion.ImageAlt = imageAlt
	existingQuestion.OrderIndex = orderIndex

	err = s.repo.Update(existingQuestion)
	return existingQuestion, err
}


func (s *QuestionService) DeleteQuestion(id string) error {
	existingQuestion, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if existingQuestion == nil {
		return errors.New("question not found")
	}

	return s.repo.Delete(id)
}
