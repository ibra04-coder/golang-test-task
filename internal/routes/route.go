package routes

import (
	"golang-test-task/internal/handlers"
	"net/http"
)

func NewRouter(h *handlers.NumberHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/numbers", h.HandleAdd)
	return mux
}