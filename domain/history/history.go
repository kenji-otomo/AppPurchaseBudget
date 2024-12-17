package history

import "time"

type History struct {
	id           *int64
	appID        int64
	amount       int64
	purchaseDate time.Time
}

func NewHistory(appID, amount int64, purchaseDate time.Time) *History {
	return &History{
		appID:        appID,
		amount:       amount,
		purchaseDate: purchaseDate,
	}
}

func (h *History) Get() (id *int64, appID, amount int64, purchaseDate time.Time) {
	id = h.id
	appID = h.appID
	amount = h.amount
	purchaseDate = h.purchaseDate
	return
}
