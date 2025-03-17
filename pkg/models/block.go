package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Block struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	TopicID primitive.ObjectID `bson:"topic_id"`
	Type    string             `bson:"type"`
	Content string             `bson:"content"`
}
