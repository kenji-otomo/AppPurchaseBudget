package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kenji-otomo/AppPurchaseBudget/domain/errorArg"
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
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(arg); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeStatus(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

func writeError(w http.ResponseWriter, status int, err *errorArg.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(err); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
