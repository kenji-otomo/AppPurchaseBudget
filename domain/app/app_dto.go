package app

import (
	"time"
)

type AppDTO struct {
	ID             *int64     `json:"id"`
	Name           string     `json:"name"`
	LastPurchaseAt *time.Time `json:"last_purchase_at"`
}

func NewAppDTO(app *App) *AppDTO {
	return &AppDTO{
		ID:             app.id,
		Name:           app.name,
		LastPurchaseAt: app.lastPurchaseAt,
	}
}

type CheckDuplicateAppResult struct {
	IsDuplicate bool `json:"is_duplicate"`
	*AppDTO
}

func NewCheckDuplicateAppResult(app *App) *CheckDuplicateAppResult {

	var appDTO *AppDTO
	if app != nil {
		appDTO = &AppDTO{
			ID:             app.id,
			Name:           app.name,
			LastPurchaseAt: app.lastPurchaseAt,
		}
	}

	return &CheckDuplicateAppResult{
		IsDuplicate: app != nil,
		AppDTO:      appDTO,
	}
}

type AppWithSumDTO struct {
	AppDTO
	Amount int64 `json:"amount"`
}

func NewAppWithSumDTO(a *AppWithSum) *AppWithSumDTO {
	return &AppWithSumDTO{
		AppDTO: AppDTO{
			ID:             a.app.id,
			Name:           a.app.name,
			LastPurchaseAt: a.app.lastPurchaseAt,
		},
		Amount: a.amount,
	}
}
