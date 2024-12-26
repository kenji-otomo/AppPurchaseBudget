package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kenji-otomo/AppPurchaseBudget/domain/budget"
	"github.com/kenji-otomo/AppPurchaseBudget/domain/errorArg"
	"github.com/kenji-otomo/AppPurchaseBudget/usecase"
)

func fetchBudgetByType(w http.ResponseWriter, r *http.Request) {

	t := chi.URLParam(r, "type")

	tint, err := strconv.Atoi(t)
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorArg.NewError(err.Error()))
	}

	b, err := usecase.FetchBudgetByType(int64(tint))
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorArg.NewError(err.Error()))
	}

	writeResponse(w, budget.NewBudgetDTO(b))
}
