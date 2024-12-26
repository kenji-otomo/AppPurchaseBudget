package budget

type Budget struct {
	id         *int64
	budgetType BudgetType
	amount     int64
}

type BudgetType int64

var (
	BudgetTypeDay   BudgetType = 1
	BudgetTypeMonth BudgetType = 2
	BudgetTypeYear  BudgetType = 3
)

func NewBudget(id *int64, budgetType BudgetType, amount int64) *Budget {
	return &Budget{
		id:         id,
		budgetType: budgetType,
		amount:     amount,
	}
}
