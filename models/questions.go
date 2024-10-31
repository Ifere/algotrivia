package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Question struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	CategoryID   primitive.ObjectID `bson:"categoryID,omitempty" json:"categoryID,omitempty"`
	QuestionText string             `bson:"questionText,omitempty" json:"questionText,omitempty"`
	QuestionType string             `bson:"questionType,omitempty" json:"questionType,omitempty"`
	Difficulty   string             `bson:"difficulty,omitempty" json:"difficulty,omitempty"`
	Data         interface{}        `bson:"data,omitempty" json:"data,omitempty"`
}

func (q *Question) SetQuestionId() {
	q.ID = primitive.NewObjectID()
}

func (q *Question) SetCreatedAt() {
	q.CreatedAt = time.Now()
}
func (q *Question) GetCreatedAt() time.Time {
	return q.CreatedAt
}
