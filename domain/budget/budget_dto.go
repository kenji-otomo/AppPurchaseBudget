package budget

type BudgetDTO struct {
	ID     *int64     `json:"id"`
	Type   BudgetType `json:"type"`
	Amount int64      `json:"amount"`
}

func NewBudgetDTO(b *Budget) *BudgetDTO {
	return &BudgetDTO{
		ID:     b.id,
		Type:   b.budgetType,
		Amount: b.amount,
	}
}
