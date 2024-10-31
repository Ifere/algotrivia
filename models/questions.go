package questions

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Question struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name            string             `bson:"name,omitempty" json:"name,omitempty"`
	CreatedAt       time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	Description     string             `bson:"description,omitempty" json:"description,omitempty"`
	Type        string             `bson:"type,omitempty" json:"type,omitempty"`

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
