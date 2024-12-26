package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kenji-otomo/AppPurchaseBudget/domain/app"
	"github.com/kenji-otomo/AppPurchaseBudget/domain/errorArg"
	"github.com/kenji-otomo/AppPurchaseBudget/usecase"
)

func getApps(w http.ResponseWriter, r *http.Request) {

	results, err := usecase.GetApps()
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorArg.NewError(err.Error()))
	}

	dtos := []*app.AppDTO{}
	for _, r := range results {
		dtos = append(dtos, app.NewAppDTO(r))
	}

	writeResponse(w, dtos)
}

func checkDuplicateApp(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var req app.AppRequest
	json.Unmarshal(body, &req)

	checkApp, err := usecase.CheckDuplicateApp(req.Name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorArg.NewError(err.Error()))
	}

	writeResponse(w, app.NewCheckDuplicateAppResult(checkApp))
}

func createApp(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var req app.AppRequest
	json.Unmarshal(body, &req)

	rApp, err := usecase.CreateApp(req.Name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorArg.NewError(err.Error()))
	}

	writeResponse(w, app.NewAppDTO(rApp))
}

func updateApp(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var reqs []*app.UpdateAppRequest
	json.Unmarshal(body, &reqs)

	if err := usecase.UpdateAppName(reqs); err != nil {
		writeError(w, http.StatusInternalServerError, errorArg.NewError(err.Error()))
	}

	writeStatus(w, http.StatusOK)
}

func fetchAppHistory(w http.ResponseWriter, r *http.Request) {

	data, err := usecase.FetchHistoryData()
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorArg.NewError(err.Error()))
	}

	results := []*app.AppWithSumDTO{}
	for _, d := range data {
		results = append(results, app.NewAppWithSumDTO(d))
	}

	writeResponse(w, results)
}
