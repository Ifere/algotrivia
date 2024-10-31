package questions

import (
	questionmodel "github.com/ifere/algotrivia/models"
)

type QuestionUseCaseI interface {
	CreateQuestion(question questionmodel.Question) (questionmodel.Question, error)
	FetchQuestions(filter interface{}) ([]questionmodel.Question, error)
	CheckDuplicate(ref string) bool
	GetQuestion(questionID string) (questionmodel.Question, error)
	UpdateQuestion(questionID string, update questionmodel.Question) (questionmodel.Question, error)
	DeleteQuestion(questionID string) (questionmodel.Question, error)
}
