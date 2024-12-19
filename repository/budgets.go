package repository

import (
	"time"
)

type Budget struct {
	ID        *int64
	Type      BudgetType
	Amount    int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type BudgetType int64

var (
	BudgetTypeDay   BudgetType = 1
	BudgetTypeMonth BudgetType = 2
	BudgetTypeYear  BudgetType = 3
)

func FetchBudgetByType(t BudgetType) (*Budget, error) {
	budget := &Budget{}
	err := db.Where("type = ?", t).First(budget).Error
	return budget, err
}
