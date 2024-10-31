package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	CategoryID primitive.ObjectID `bson:"categoryID,omitempty" json:"categoryID,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	Name       string             `bson:"name,omitempty" json:"name,omitempty"`
	Image      string             `bson:"image,omitempty" json:"image,omitempty"`
	Tags	   []string 		  `bson:"tags,omitempty" json:"tags,omitempty"`
}

func (c *Category) SetCategoryId() {
	c.CategoryID = primitive.NewObjectID()

}
func (c *Category) SetCreatedAt() {
	c.CreatedAt = time.Now()
}
func (c *Category) getCreatedAt() time.Time {
	return c.CreatedAt
}
