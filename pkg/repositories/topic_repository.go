package repositories

import (
	"context"
	"go-book/pkg/db"
	"go-book/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopicRepository struct {
	collection *mongo.Collection
}

func NewTopicRepository(collection *mongo.Collection) *TopicRepository {
	return &TopicRepository{collection: collection}
}

func (r *TopicRepository) GetTopics(ctx context.Context) ([]models.Topic, error) {
	var topics []models.Topic

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var topic models.Topic
		if err := cursor.Decode(&topic); err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return topics, nil
}

func (t *TopicRepository) InsertTopic(ctx context.Context, topic models.Topic) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	_, err := t.collection.InsertOne(ctx, topic)
	return err
}

func (t *TopicRepository) DeleteTopic(ctx context.Context, topic models.Topic) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := t.collection.DeleteOne(ctx, bson.M{"_id": topic.ID})
	if err != nil {
		return err
	}
	return nil
}

func GetTopicInfo(ctx context.Context, id primitive.ObjectID) (models.Topic, error) {
	var topic models.Topic
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err := db.Client.Database("go_book").Collection("topics").FindOne(ctx, bson.M{"_id": id}).Decode(&topic)
	return topic, err
}
