package dbs

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
	"time"
)

const (
	dbName = "algotrivia"
)

type MongoDB struct {
	Client *mongo.Client
}

// create a connection to mongodb
func (db *MongoDB) Connect() error {

	url := db.GetDbUrl()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	logrus.Println("db connected")
	db.Client = client

	return nil
}

// create a connection to the test db
func (db *MongoDB) ConnectTest() error {

	url := db.GetDbUrl()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	logrus.Println("db connected")
	db.Client = client

	return nil
}

// create a connection to mongodb
func (db *MongoDB) TestTearDown(coll string) error {
	c := db.Client.Database(dbName).Collection(coll)

	return c.Drop(nil)
}

func (db *MongoDB) GetDbUrl() string {
	e := godotenv.Load()

	if e != nil {
		fmt.Println(e)
	}
	env := os.Getenv("ENV")
	var dbUrl string

	switch env {
	case "dev":
		dbUrl = os.Getenv("MONGO_DEV_URL")
	case "prod":
		dbUrl = os.Getenv("MONGO_PROD_URL")
	default:
		dbUrl = os.Getenv("MONGO_LOCAL_URL")
	}
	return dbUrl
}



func (db MongoDB) CreateTTLIndex(coll *mongo.Collection, expir int32) error {
	index := mongo.IndexModel{
		Keys:    bson.D{{Key: "created_at", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(expir),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := coll.Indexes().CreateOne(ctx, index)

	if err != nil {
		return err
	}

	return nil
}

// returns flag for duplicate error

func (db MongoDB) IsMongoDuplicateError(err error) bool {

	if strings.Contains(err.Error(), "E11000") {
		return true
	}
	return false
}


// returns an instance of the questions collection
func (db MongoDB) QuestionCollection() *mongo.Collection {
	return db.Client.Database(dbName).Collection("questions")

}
