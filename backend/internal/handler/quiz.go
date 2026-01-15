package handler

import (
	"github.com/gofiber/fiber/v2"
	"quiz.com/quiz/internal/service"
)

type QuizHandler struct {
	quizService *service.QuizService
}

func NewQuizHandler(quizService *service.QuizService) QuizHandler {
	return QuizHandler{
		quizService: quizService,
	}
}

func (h QuizHandler) GetQuizzes(ctx *fiber.Ctx) error {
	quizzes, err := h.quizService.GetQuizzes()
	if err != nil {
		return err
	}

	return ctx.JSON(quizzes)
}
