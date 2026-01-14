package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"quiz.com/quiz/internal/entity"
)

type QuizRepository struct {
	collection *mongo.Collection
}

func NewQuizRepository(collection *mongo.Collection) *QuizRepository {
	return &QuizRepository{
		collection: collection,
	}
}

func (r QuizRepository) InsertQuiz(quiz entity.Quiz) error {
	_, err := r.collection.InsertOne(context.Background(), quiz)
	return err
}

func (r QuizRepository) GetQuizzes() ([]entity.Quiz, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var quiz []entity.Quiz
	err = cursor.All(context.Background(), &quiz)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func (r QuizRepository) GetQuizById(id primitive.ObjectID) (*entity.Quiz, error) {
	result := r.collection.FindOne(context.Background(), bson.M{"_id": id})

	var quiz entity.Quiz
	err := result.Decode(&quiz)
	if err != nil {
		return nil, err
	}

	return &quiz, nil
}
