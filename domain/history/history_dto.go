package history

import (
	"time"
)

type HistoryDTO struct {
	ID           *int64    `json:"id"`
	AppID        int64     `json:"app_id"`
	Amount       int64     `json:"amount"`
	PurchaseDate time.Time `json:"purchase_date"`
}

func NewHistoryDTO(h *History) *HistoryDTO {
	return &HistoryDTO{
		ID:           h.id,
		AppID:        h.appID,
		Amount:       h.amount,
		PurchaseDate: h.purchaseDate,
	}
}

type HistoryWithNameDTO struct {
	HistoryDTO
	Name string `json:"name"`
}

func NewHistoryWithNameDTO(h *History) *HistoryWithNameDTO {
	return &HistoryWithNameDTO{
		HistoryDTO: *NewHistoryDTO(h),
		Name:       *h.appName,
	}
}
