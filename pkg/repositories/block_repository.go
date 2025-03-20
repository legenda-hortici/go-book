package repositories

import (
	"context"
	"go-book/pkg/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlockRepository struct {
	collection *mongo.Collection
}

func NewBlockRepository(collection *mongo.Collection) *BlockRepository {
	return &BlockRepository{collection: collection}
}

func (b *BlockRepository) InsertBlock(block models.Block) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := b.collection.InsertOne(ctx, block)
	if err != nil {
		log.Printf("Error inserting block: %v", err)
		return err
	}

	return nil
}

func (r *BlockRepository) GetBlocks(topicID primitive.ObjectID) ([]models.Block, error) {
	var blocks []models.Block
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"topic_id": topicID})
	if err != nil {
		log.Printf("Error finding blocks: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var block models.Block
		if err := cursor.Decode(&block); err != nil {
			log.Printf("Error decoding block: %v", err)
			return nil, err
		}
		blocks = append(blocks, block)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blocks, nil
}

func (r *BlockRepository) DeleteAllBlocks(topicID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteMany(ctx, bson.M{"topic_id": topicID})
	if err != nil {
		return err
	}
	return nil
}

func (r *BlockRepository) DeleteBlock(blockID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": blockID})
	if err != nil {
		return err
	}
	return nil
}

func (r *BlockRepository) UpdateBlock(blockID primitive.ObjectID, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": blockID}, bson.M{"$set": bson.M{"content": content}})
	if err != nil {
		return err
	}

	return nil
}
