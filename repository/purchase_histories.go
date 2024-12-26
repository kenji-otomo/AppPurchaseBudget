package repository

import (
	"time"

	"github.com/kenji-otomo/AppPurchaseBudget/domain/history"
	"gorm.io/gorm"
)

type PurchaseHistory struct {
	ID           *int64
	AppID        int64
	Amount       int64
	PurchaseDate time.Time
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type PurchaseHistoryWithName struct {
	PurchaseHistory
	Name string
}

func NewPurchaseHistory(h *history.History) *PurchaseHistory {
	id, appID, amount, purchaseDate := h.Get()
	purchaseDate = purchaseDate.Add(9 * time.Hour) // dateで保存する関係上、日本時間の日付とするために9時間をプラスする

	return &PurchaseHistory{
		ID:           id,
		AppID:        appID,
		Amount:       amount,
		PurchaseDate: purchaseDate,
	}
}

func ListPurchaseHistories(start, end time.Time) ([]*PurchaseHistoryWithName, error) {
	results := []*PurchaseHistoryWithName{}
	err := db.Select("p.*, a.name").
		Table("purchase_histories p").
		Joins("INNER JOIN apps a ON p.app_id = a.id").
		Where("p.created_at BETWEEN ? AND ?", start, end).Scan(&results).Error
	return results, err
}

func FetchTotalPurchaseAmount(start, end time.Time) (*int64, error) {
	var totalAmount *int64
	err := db.Select("SUM(amount)").
		Table("purchase_histories").
		Where("created_at BETWEEN ? AND ?", start, end).Scan(&totalAmount).Error

	return totalAmount, err
}

func (p *PurchaseHistory) Create(tx *gorm.DB) error {
	return tx.Create(&p).Error
}
