package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kenji-otomo/AppPurchaseBudget/usecase"
)

// ルーティング
func Route(r *chi.Mux) {

	r.Get("/", usecase.HelloWorld)
	r.Get("/game", getApps)
	r.Post("/game", createApp)
}

func writeResponse(w http.ResponseWriter, arg any) {

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(arg); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
