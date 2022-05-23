package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bug struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	UserId      primitive.ObjectID `bson:"userId" json:"userId"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}
