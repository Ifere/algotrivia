package questionrepo

import (
	dbs "github.com/ifere/algotrivia/config/db"
	questionmodel "github.com/ifere/algotrivia/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

type MongodbRepo struct {
	dbs.MongoDB
}

func NewMongodbTQuestionRepo(mongoDB dbs.MongoDB) *MongodbRepo {
	return &MongodbRepo{MongoDB: mongoDB}
}

type graphData struct {
	Amount float64 `json:"amount" bson:"amount"`
	Date   string  `json:"date" bson:"date"`
}

func (c MongodbRepo) CreateQuestion(question questionmodel.Question) (questionmodel.Question, error) {
	coll := c.QuestionCollection()

	question.SetCreatedAt()

	result, err := coll.InsertOne(nil, question)

	if err != nil {
		return questionmodel.Question{}, err
	}

	question.ID = result.InsertedID.(primitive.ObjectID)

	return question, nil
}

func (c MongodbRepo) FetchQuestionsFromDb(filter interface{}) ([]questionmodel.Question, error) {
	coll := c.QuestionCollection()

	var questions []questionmodel.Question

	collation := options.Collation{Locale: "en_US", Strength: 2}
	option := options.Find().SetCollation(&collation)

	cur, err := coll.Find(nil, filter, option)
	if err != nil {
		return []questionmodel.Question{}, err
	}

	err = cur.All(nil, &questions)

	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (c MongodbRepo) GetQuestion(questionID string) (questionmodel.Question, error) {
	coll := c.QuestionCollection()

	var question questionmodel.Question

	id, _ := primitive.ObjectIDFromHex(questionID)

	result := coll.FindOne(nil, bson.D{{Key: "_id", Value: id}})

	err := result.Decode(&question)

	if err != nil {
		return questionmodel.Question{}, nil
	}

	return question, nil
}

func (c MongodbRepo) UpdateQuestion(questionID string, update questionmodel.Question) (questionmodel.Question, error) {
	coll := c.QuestionCollection()

	id, _ := primitive.ObjectIDFromHex(questionID)

	var question questionmodel.Question

	option := options.FindOneAndUpdate().SetReturnDocument(options.ReturnDocument(options.After))

	err := coll.FindOneAndUpdate(nil, bson.D{{Key: "_id", Value: id}}, bson.D{{Key: "$set", Value: update}}, option).Decode(&question)

	if err != nil {
		return questionmodel.Question{}, err
	}

	return question, nil
}

func (c MongodbRepo) DeleteQuestion(questionID string) (questionmodel.Question, error) {
	coll := c.QuestionCollection()

	var question questionmodel.Question

	id, _ := primitive.ObjectIDFromHex(questionID)

	err := coll.FindOneAndDelete(nil, bson.D{{Key: "_id", Value: id}}).Decode(&question)

	if err != nil {
		return questionmodel.Question{}, err
	}

	return question, nil

}

func (c MongodbRepo) CheckDuplicate(ref string) bool {
	coll := c.QuestionCollection()

	var question questionmodel.Question

	err := coll.FindOne(nil, bson.D{{Key: "reference", Value: ref}}).Decode(&question)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false
		}
	}

	return true

}

func prepareDateRange() (dateRange bson.M, days float64) {

	beginningOfMonth := Date(time.Now().Year(), time.Now().Month(), 1)

	endOfMonth := Date(time.Now().Year(), time.Now().AddDate(0, 1, 0).Month(), 0)

	days = float64(endOfMonth.Day())
	return bson.M{"$gte": beginningOfMonth, "$lte": endOfMonth}, days
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
