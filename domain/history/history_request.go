package history

import "time"

type HistoryRequest struct {
	AppID        int64     `json:"app_id"`
	Amount       int64     `json:"amount"`
	PurchaseDate time.Time `json:"purchase_date"`
}
