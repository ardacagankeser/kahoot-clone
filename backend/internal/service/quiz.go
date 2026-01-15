package service

import (
	"quiz.com/quiz/internal/entity"
	"quiz.com/quiz/internal/repository"
)

type QuizService struct {
	quizRepository *repository.QuizRepository
}

func NewQuizService(quizRepository *repository.QuizRepository) *QuizService {
	return &QuizService{
		quizRepository: quizRepository,
	}
}

func (s QuizService) GetQuizzes() ([]entity.Quiz, error) {
	return s.quizRepository.GetQuizzes()
}
