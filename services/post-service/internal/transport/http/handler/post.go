package handler

import (
	"encoding/json"
	"net/http"
	"posts/internal/transport/http/dto"
	"posts/internal/usecases/post"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type PostHandler struct {
	createUC *post.CreateUseCase
	getUC    *post.GetUseCase
	updateUC *post.UpdateUseCase
	deleteUC *post.DeleteUseCase

	listUC *post.ListUseCase
}

func NewPostHandler(
	createUC *post.CreateUseCase,
	getUC *post.GetUseCase,
	updateUC *post.UpdateUseCase,
	deleteUC *post.DeleteUseCase,

	listUC *post.ListUseCase,
) *PostHandler {
	return &PostHandler{
		createUC: createUC,
		getUC:    getUC,
		updateUC: updateUC,
		deleteUC: deleteUC,

		listUC: listUC,
	}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePostRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	authorID, err := uuid.Parse(req.AuthorID)
	if err != nil {
		http.Error(w, "invalid author id", http.StatusBadRequest)
		return
	}

	post, err := h.createUC.Execute(r.Context(), post.CreateCommand{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: authorID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.PostResponse{
		ID:       post.ID.String(),
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *PostHandler) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}

	post, err := h.getUC.Execute(r.Context(), post.GetCommand{
		ID: id,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.PostResponse{
		ID:       post.ID.String(),
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}

	var req dto.UpdatePostRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	post, err := h.updateUC.Execute(r.Context(), post.UpdateCommand{
		ID:      id,
		Title:   &req.Title,
		Content: &req.Content,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.PostResponse{
		ID:       post.ID.String(),
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}

	post, err := h.deleteUC.Execute(r.Context(), post.DeleteCommand{
		ID: id,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.PostResponse{
		ID:       post.ID.String(),
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	postList, err := h.listUC.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.ListPostsResponse{
		Posts: make([]dto.PostResponse, len(postList)),
	}

	for i, post := range postList {
		resp.Posts[i] = dto.PostResponse{
			ID:       post.ID.String(),
			Title:    post.Title,
			Content:  post.Content,
			AuthorID: post.AuthorID.String(),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
