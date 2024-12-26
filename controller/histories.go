package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kenji-otomo/AppPurchaseBudget/domain/errorArg"
	"github.com/kenji-otomo/AppPurchaseBudget/domain/history"
	"github.com/kenji-otomo/AppPurchaseBudget/usecase"
)

func getHistories(w http.ResponseWriter, r *http.Request) {
	results, err := usecase.GetHistories()
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorArg.NewError(err.Error()))
	}

	dtos := []*history.HistoryWithNameDTO{}
	for _, result := range results {
		dtos = append(dtos, history.NewHistoryWithNameDTO(result))
	}

	writeResponse(w, dtos)
}

func createHitory(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var req history.HistoryRequest
	json.Unmarshal(body, &req)

	if err := usecase.CreateHitory(req); err != nil {
		writeError(w, http.StatusInternalServerError, errorArg.NewError(err.Error()))
	}

	writeStatus(w, http.StatusCreated)
}
