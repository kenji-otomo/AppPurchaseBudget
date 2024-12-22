package budget

import "github.com/kenji-otomo/AppPurchaseBudget/repository"

type BudgetDTO struct {
	ID     *int64                `json:"id"`
	Type   repository.BudgetType `json:"type"`
	Amount int64                 `json:"amount"`
}

func NewBudgetDTO(b *repository.Budget) *BudgetDTO {
	return &BudgetDTO{
		ID:     b.ID,
		Type:   b.Type,
		Amount: b.Amount,
	}
}
