package main

import (
	"fmt"
	"go-book/internal/handlers"
	"go-book/internal/services"
	"go-book/pkg/db"
	"go-book/pkg/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() (*mux.Router, error) {
	// Инизиализация репозитория и сервиса для Topic
	topicsCollection, err := db.GetDB("topics")
	if err != nil {
		return nil, fmt.Errorf("failed to get topics collection: %w", err)
	}
	topicRepo := repositories.NewTopicRepository(topicsCollection)
	topicService := services.NewTopicService(topicRepo)
	topicHandlers := handlers.NewTopicHandler(topicService)

	// Инизиализация репозитория и сервиса для Block
	blocksCollection, err := db.GetDB("blocks")
	if err != nil {
		return nil, fmt.Errorf("failed to get blocks collection: %w", err)
	}
	blockRepo := repositories.NewBlockRepository(blocksCollection)
	blockService := services.NewBlockService(blockRepo)
	blockHandlers := handlers.NewBlockHandler(blockService)

	mux := mux.NewRouter()
	mux.HandleFunc("/", topicHandlers.MainHandler)
	mux.HandleFunc("/add_topic", topicHandlers.CreateTopicHandler).Methods("POST")
	mux.HandleFunc("/delete_topic/{id}", topicHandlers.DeleteTopicHandler).Methods("POST")
	mux.HandleFunc("/topic/{id}", topicHandlers.TopicHandler)

	mux.HandleFunc("/topic/{id}/blocks", blockHandlers.ShowBlockHandler)
	mux.HandleFunc("/topic/{id}/add_block", blockHandlers.AddBlockHandler).Methods("POST")
	mux.HandleFunc("/blocks/delete/{id}", blockHandlers.DeleteBlockHandler).Methods("POST")
	mux.HandleFunc("/blocks/edit/{id}", blockHandlers.UpdateBlockHandler).Methods("POST")

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	return mux, nil
}
