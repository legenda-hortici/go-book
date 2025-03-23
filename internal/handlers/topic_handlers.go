package handlers

import (
	"go-book/internal/services"
	"go-book/pkg/db"
	"go-book/pkg/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("web/templates/*.html"))

type TopicHandler struct {
	service *services.TopicService
}

func NewTopicHandler(service *services.TopicService) *TopicHandler {
	return &TopicHandler{service: service}
}

func (h *TopicHandler) MainHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	topics, err := h.service.GetTopics(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Topics": topics,
	}

	err = templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Printf("Error executing template:%v", err)
	}
}

func (h *TopicHandler) CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	var topic models.Topic
	title := r.FormValue("topicTitle")
	description := r.FormValue("topicDescription")

	topic = models.Topic{
		Title:       title,
		Description: description,
	}

	if err := h.service.CreateTopic(ctx, topic); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *TopicHandler) DeleteTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := db.ExtractObjectID(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTopic(ctx, models.Topic{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *TopicHandler) TopicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["id"]
	http.Redirect(w, r, "/topic/"+vars+"/blocks", http.StatusSeeOther)
}
