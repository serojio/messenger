package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"users/internal/transport/http/dto"
	"users/internal/usecases/user"
)

type UserHandler struct {
	createUC *user.CreateUseCase
	getUC    *user.GetUseCase
	updateUC *user.UpdateUseCase
	deleteUC *user.DeleteUseCase
}

func NewUserHandler(
	createUC *user.CreateUseCase,
	getUC *user.GetUseCase,
	updateUC *user.UpdateUseCase,
	deleteUC *user.DeleteUseCase,
) *UserHandler {
	return &UserHandler{
		createUC: createUC,
		getUC:    getUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, err := h.createUC.Execute(r.Context(), user.CreateCommand{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.UserResponse{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
