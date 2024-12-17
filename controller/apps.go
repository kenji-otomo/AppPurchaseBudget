package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/kenji-otomo/AppPurchaseBudget/domain/app"
	"github.com/kenji-otomo/AppPurchaseBudget/usecase"
)

func getApps(w http.ResponseWriter, r *http.Request) {

	results, err := usecase.GetApps()
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, results)
}

func checkDuplicateApp(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var req app.AppRequest
	json.Unmarshal(body, &req)

	checkApp, err := usecase.CheckDuplicateApp(req.Name)
	if err != nil {
		log.Fatal(err)
	}

	result := app.NewCheckDuplicateAppResult(checkApp)

	writeResponse(w, result)
}

func createApp(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var req app.AppRequest
	json.Unmarshal(body, &req)

	if err := usecase.CreateApp(req.Name); err != nil {
		log.Fatal(err)
	}
}
