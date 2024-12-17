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

func NewPurchaseHistory(h *history.History) *PurchaseHistory {
	_, appID, amount, purchaseDate := h.Get()
	purchaseDate = purchaseDate.Add(9 * time.Hour) // dateで保存する関係上、日本時間の日付とするために9時間をプラスする

	return &PurchaseHistory{
		AppID:        appID,
		Amount:       amount,
		PurchaseDate: purchaseDate,
	}
}

func ListPurchaseHistories() ([]*PurchaseHistory, error) {
	results := []*PurchaseHistory{}
	err := db.Find(&results).Error
	return results, err
}

func (p *PurchaseHistory) Create(tx *gorm.DB) error {
	return tx.Create(&p).Error
}
