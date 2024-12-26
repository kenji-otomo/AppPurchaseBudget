package repository

import (
	"time"

	"github.com/kenji-otomo/AppPurchaseBudget/domain/budget"
)

type Budget struct {
	ID        *int64
	Type      budget.BudgetType
	Amount    int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func FetchBudgetByType(t budget.BudgetType) (*Budget, error) {
	budget := &Budget{}
	err := db.Where("type = ?", t).First(budget).Error
	return budget, err
}
