package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kenji-otomo/AppPurchaseBudget/domain/budget"
	"github.com/kenji-otomo/AppPurchaseBudget/usecase"
)

func fetchBudgetByType(w http.ResponseWriter, r *http.Request) {

	t := chi.URLParam(r, "type")

	tint, err := strconv.Atoi(t)
	if err != nil {
		log.Fatal(err)
	}

	b, err := usecase.FetchBudgetByType(int64(tint))
	if err != nil {
		log.Fatal(err)
	}

	writeResponse(w, budget.NewBudgetDTO(b))
}
