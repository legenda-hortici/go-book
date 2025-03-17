package db

import (
	"context"
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
func InitDB() {
	// Загрузка переменных окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Получаем URI и имя базы данных из переменных окружения
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	if mongoURI == "" || dbName == "" {
		log.Fatalf("Mongo URI or DB Name is not set in .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	var err2 error
	Client, err2 = mongo.Connect(ctx, clientOptions)
	if err2 != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err2)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB!")
}

func GetDB(collectionName string) *mongo.Collection {
	if Client == nil {
		log.Fatal("MongoDB client is not initialized")
	}

	dbName := os.Getenv("DB_NAME")
	return Client.Database(dbName).Collection(collectionName)
}

func ExtractObjectID(input string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(input)
}

func CloseDB() {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := Client.Disconnect(ctx); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
		log.Println("Disconnected from MongoDB.")
	}
}
