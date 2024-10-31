package questions

import (
	questionmodel "github.com/ifere/algotrivia/models"
)

//go:generate mockgen -destination=../../mocks/mock_transaction_repo.go -package=mocks nebula-backend/pkg/transactions TransactionRepoI

type QuestionMongoRepoI interface {
	CreateQuestion(question questionmodel.Question) (questionmodel.Question, error)
	FetchQuestionsFromDb(filter interface{}) ([]questionmodel.Question, error)
	CheckDuplicate(ref string) bool
	GetQuestion(questionID string) (questionmodel.Question, error)
	UpdateQuestion(questionID string, update questionmodel.Question) (questionmodel.Question, error)
	DeleteQuestion(questionID string) (questionmodel.Question, error)
}

