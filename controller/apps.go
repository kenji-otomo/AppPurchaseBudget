package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/kenji-otomo/AppPurchaseBudget/usecase"
)

func getApps(w http.ResponseWriter, r *http.Request) {

	results, err := usecase.GetApps()
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, results)
}

func createApp(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var req GameRequest
	json.Unmarshal(body, &req)

	if err := usecase.CreateApp(req.Name); err != nil {
		log.Fatal(err)
	}
}

type GameRequest struct {
	Name string `json:"name"`
}
