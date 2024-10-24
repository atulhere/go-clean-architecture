package handler

import (
	"encoding/json"
	"fmt"
	"go-clean-architecture/usecase"
	"net/http"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: uc}
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Call aa rahi hai")
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := h.UserUsecase.Login(request.Username, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
