package main

import (
	"go-book/internal/handlers"
	"go-book/internal/services"
	"go-book/pkg/db"
	"go-book/pkg/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {

	// Инизиализация репозитория и сервиса для Topic
	topicRepo := repositories.NewTopicRepository(db.GetDB("topics"))
	topicService := services.NewTopicService(topicRepo)
	topicHandlers := handlers.NewTopicHandler(topicService)

	// Инизиализация репозитория и сервиса для Block
	blockRepo := repositories.NewBlockRepository(db.GetDB("blocks"))
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

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	return mux
}
