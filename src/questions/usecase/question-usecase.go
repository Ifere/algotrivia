package questionusecase

import (
	"github.com/pkg/errors"
	apperrors "github.com/ifere/algotrivia/config/errors"
	questionmodel "github.com/ifere/algotrivia/models"
	"github.com/ifere/algotrivia/src/questions"
)

type QuestionUseCase struct {
	QuestionMongoRepo questions.QuestionMongoRepoI

}

func NewQuestionUseCase(questionMongoRepo questions.QuestionMongoRepoI) QuestionUseCase {
	return QuestionUseCase{questionMongoRepo}
}

func (c QuestionUseCase) CreateQuestion(question questionmodel.Question) (questionmodel.Question, error) {

	question, err := c.QuestionMongoRepo.CreateQuestion(question)

	if err != nil {
		return questionmodel.Question{}, errors.Wrap(err, apperrors.NotCreated{Resource: "question"}.Error())
	}

	return question, nil
}

func (c QuestionUseCase) FetchQuestions(filter interface{}) ([]questionmodel.Question, error) {
	qns, err := c.QuestionMongoRepo.FetchQuestionsFromDb(filter)

	if err != nil {
		return []questionmodel.Question{}, errors.Wrap(err, apperrors.ErrorGetting{Resource: "questions"}.Error())
	}

	return qns, nil
}

func (c QuestionUseCase) GetQuestion(questionID string) (questionmodel.Question, error) {
	question, err := c.QuestionMongoRepo.GetQuestion(questionID)

	if err != nil {
		return questionmodel.Question{}, errors.Wrap(err, apperrors.ErrorGetting{Resource: "question"}.Error())
	}

	return question, nil
}

func (c QuestionUseCase) UpdateQuestion(questionID string, update questionmodel.Question) (questionmodel.Question, error) {
	question, err := c.QuestionMongoRepo.UpdateQuestion(questionID, update)

	if err != nil {
		return questionmodel.Question{}, errors.Wrap(err, apperrors.ErrorUpdating{Resource: "question"}.Error())
	}

	return question, nil
}

func (c QuestionUseCase) DeleteQuestion(questionID string) (questionmodel.Question, error) {
	question, err := c.QuestionMongoRepo.DeleteQuestion(questionID)

	if err != nil {
		return questionmodel.Question{}, errors.Wrap(err, apperrors.ErrorDeleting{Resource: "question"}.Error())
	}

	return question, nil
}

func (c QuestionUseCase) CheckDuplicate(ref string) bool {
	return c.QuestionMongoRepo.CheckDuplicate(ref)
}
