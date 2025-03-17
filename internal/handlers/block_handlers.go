package handlers

import (
	"fmt"
	"go-book/internal/services"
	"go-book/pkg/models"
	"net/http"

	"go-book/pkg/repositories"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlockHandler struct {
	service *services.BlockService
}

func NewBlockHandler(service *services.BlockService) *BlockHandler {
	return &BlockHandler{service: service}
}

func (h *BlockHandler) AddBlockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var block models.Block
	content := r.FormValue("content")
	blockType := r.FormValue("blockType")

	block = models.Block{
		ID:      primitive.NewObjectID(),
		TopicID: id,
		Type:    blockType,
		Content: content,
	}

	if err := h.service.AddBlock(block); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/topic/%s", vars), http.StatusSeeOther)
}

var TopicID string

func (h *BlockHandler) ShowBlockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["id"]
	TopicID = vars

	id, err := primitive.ObjectIDFromHex(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blocks, err := h.service.GetBlocks(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	topic, err := repositories.GetTopicInfo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Blocks": blocks,
		"Topic":  topic,
	}
	templates.ExecuteTemplate(w, "blocks", data)
}

func (h *BlockHandler) DeleteBlockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteBlock(models.Block{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/topic/"+TopicID+"/blocks", http.StatusSeeOther)
}
