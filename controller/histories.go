package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/kenji-otomo/AppPurchaseBudget/domain/history"
	"github.com/kenji-otomo/AppPurchaseBudget/usecase"
)

func getHistories(w http.ResponseWriter, r *http.Request) {
	results, err := usecase.GetHistories()
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, results)
}

func createHitory(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var req history.HistoryRequest
	json.Unmarshal(body, &req)

	if err := usecase.CreateHitory(req); err != nil {
		log.Fatal(err)
	}
}
