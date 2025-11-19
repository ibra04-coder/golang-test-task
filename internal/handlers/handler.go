package handlers

import (
	"encoding/json"
	"golang-test-task/internal/entity"
	"golang-test-task/internal/usecase"
	"net/http"
)

type NumberHandler struct {
	uc usecase.NumberUseCase
}

func NewNumberHandler(uc usecase.NumberUseCase) *NumberHandler {
	return &NumberHandler{uc: uc}
}

func (h *NumberHandler) HandleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input entity.NumberInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	sortedNumbers, err := h.uc.AddAndGetSorted(input.Number)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sortedNumbers)
}