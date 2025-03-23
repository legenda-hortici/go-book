package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Инициализация базы данных
func InitDB() error {
	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	// Получаем URI и имя базы данных из переменных окружения
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	if mongoURI == "" || dbName == "" {
		return fmt.Errorf("mongo URI or DB Name is not set in .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	if err = Client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("error pinging MongoDB: %w", err)
	}

	log.Println("Connected to MongoDB!")
	return nil
}

func GetDB(collectionName string) (*mongo.Collection, error) {
	if Client == nil {
		return nil, fmt.Errorf("MongoDB client is not initialized")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("DB_NAME environment variable is not set")
	}

	return Client.Database(dbName).Collection(collectionName), nil
}

func ExtractObjectID(input string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(input)
}

func CloseDB() {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := Client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
		log.Println("Disconnected from MongoDB.")
	}
}
