package app

import (
	"time"

	"github.com/kenji-otomo/AppPurchaseBudget/repository"
)

type AppDTO struct {
	ID             *int64    `json:"id"`
	Name           string    `json:"name"`
	LastPurchaseAt time.Time `json:"last_purchase_at"`
}

func NewAppDTO(appa *repository.App) *AppDTO {
	return &AppDTO{
		ID:             appa.ID,
		Name:           appa.Name,
		LastPurchaseAt: appa.LastPurchaseAt,
	}
}

type CheckDuplicateAppResult struct {
	IsDuplicate bool `json:"is_duplicate"`
	*AppDTO
}

func NewCheckDuplicateAppResult(appa *repository.App) *CheckDuplicateAppResult {

	var appDTO *AppDTO
	if appa != nil {
		appDTO = &AppDTO{
			ID:             appa.ID,
			Name:           appa.Name,
			LastPurchaseAt: appa.LastPurchaseAt,
		}
	}

	return &CheckDuplicateAppResult{
		IsDuplicate: appa != nil,
		AppDTO:      appDTO,
	}
}

type AppWithSumDTO struct {
	AppDTO
	Amount int64 `json:"amount"`
}

func NewAppWithSumDTO(r *repository.AppWithSum) *AppWithSumDTO {
	return &AppWithSumDTO{
		AppDTO: AppDTO{
			ID:             r.App.ID,
			Name:           r.App.Name,
			LastPurchaseAt: r.App.LastPurchaseAt,
		},
		Amount: r.Amount,
	}
}
