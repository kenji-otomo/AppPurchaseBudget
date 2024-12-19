package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// ルーティング
func Route(r *chi.Mux) {

	r.Get("/app", getApps)
	r.Post("/app", createApp)
	r.Post("/app/check", checkDuplicateApp)
	r.Patch("/app", updateApp)
	r.Get("/app/history", fetchAppHistory)

	r.Get("/history", getHistories)
	r.Post("/history", createHitory)

	r.Get("/budget/{type}", fetchBudgetByType)
}

func writeResponse(w http.ResponseWriter, arg any) {

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(arg); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
