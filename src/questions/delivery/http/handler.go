package questiondelivery

import (
	questionmodel "github.com/ifere/algotrivia/models"
	"github.com/ifere/algotrivia/src/questions"
	httplib "github.com/ifere/algotrivia/utils/http"
	"net/http"
	"strings"
	"time"
	"github.com/derekstavis/go-qs"
	"github.com/ifere/algotrivia/config/responses"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type QuestionDelivery struct {
	questionUseCase questions.QuestionUseCaseI
}

func NewQuestion(usecase questions.QuestionUseCaseI) QuestionDelivery {
	return QuestionDelivery{usecase}
}
func prepareFilter(filter bson.M) bson.M {


	dateCreated, ok := filter["created_at"]

	if ok {

		date := strings.Split(dateCreated.(string), "|")
		from, _ := time.Parse(time.RFC3339, date[0])
		to, _ := time.Parse(time.RFC3339, date[1])
		filter["created_at"] = bson.M{"$gte": from, "$lte": to}
	}

	return filter

}

func (t QuestionDelivery) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	c := httplib.C{W: w, R: r}

	var payload questionmodel.Question

	c.BindJSON(&payload)

	question, err := t.questionUseCase.CreateQuestion(payload)

	if err != nil {
		resp := responses.GeneralResponse{
			Success: false,
			Data:    nil,
			Message: errors.Cause(err).Error(),
		}

		c.Response(resp)
	}

	resp := responses.GeneralResponse{
		Success: true,
		Data:    question,
		Message: "question created",
	}

	c.Response(resp)

}

func (t QuestionDelivery) FetchQuestions(w http.ResponseWriter, r *http.Request) {

	c := httplib.C{W: w, R: r}

	filterBy := bson.M{}

	query := strings.SplitAfter(r.RequestURI, "?")

	if len(query) >= 2 {
		filter, _ := qs.Unmarshal(query[1])
		filterBy = filter

	}

	questions, err := t.questionUseCase.FetchQuestions(prepareFilter(filterBy))

	if err != nil {
		resp := responses.GeneralResponse{
			Success: false,
			Data:    nil,
			Message: errors.Cause(err).Error(),
		}

		c.Response(resp)
	} else {

		resp := responses.GeneralResponse{
			Success: true,
			Data:    questions,
			Message: "questions",
		}

		c.Response(resp)
	}
}


func (t QuestionDelivery) GetQuestion(w http.ResponseWriter, r *http.Request) {
	c := httplib.C{W: w, R: r}

	questionID := c.Params("questionID")

	question, err := t.questionUseCase.GetQuestion(questionID)

	if err != nil {
		resp := responses.GeneralResponse{
			Success: false,
			Data:    nil,
			Message: errors.Cause(err).Error(),
		}

		c.Response(resp)
	} else {

		resp := responses.GeneralResponse{
			Success: true,
			Data:    question,
			Message: "question retrieved",
		}

		c.Response(resp)
	}
}

func (t QuestionDelivery) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	c := httplib.C{W: w, R: r}

	questionID := c.Params("questionID")

	var payload questionmodel.Question

	c.BindJSON(&payload)

	sharing, err := t.questionUseCase.UpdateQuestion(questionID, payload)

	if err != nil {
		resp := responses.GeneralResponse{
			Success: false,
			Data:    nil,
			Message: errors.Cause(err).Error(),
		}

		c.Response(resp)
	}

	resp := responses.GeneralResponse{
		Success: true,
		Data:    sharing,
		Message: "question updated",
	}

	c.Response(resp)
}

func (t QuestionDelivery) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	c := httplib.C{W: w, R: r}

	questionID := c.Params("questionID")

	questions, err := t.questionUseCase.DeleteQuestion(questionID)

	if err != nil {
		resp := responses.GeneralResponse{
			Success: false,
			Data:    nil,
			Message: errors.Cause(err).Error(),
		}

		c.Response(resp)
	}

	resp := responses.GeneralResponse{
		Success: true,
		Data:    questions,
		Message: "question deleted",
	}

	c.Response(resp)
}
