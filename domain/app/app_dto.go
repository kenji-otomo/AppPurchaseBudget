package app

import (
	"time"

	"github.com/kenji-otomo/AppPurchaseBudget/repository"
)

type AppDTO struct {
	ID             *int64
	Name           string
	LastPurchaseAt time.Time
}

type CheckDuplicateAppResult struct {
	IsDuplicate bool
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
