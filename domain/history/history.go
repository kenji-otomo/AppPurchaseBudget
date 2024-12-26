package history

import "time"

type History struct {
	id           *int64
	appID        int64
	appName      *string
	amount       int64
	purchaseDate time.Time
}

func NewHistory(id *int64, appID, amount int64, purchaseDate time.Time, appName *string) *History {
	return &History{
		id:           id,
		appID:        appID,
		amount:       amount,
		purchaseDate: purchaseDate,
		appName:      appName,
	}
}

func (h *History) Get() (id *int64, appID, amount int64, purchaseDate time.Time) {
	id = h.id
	appID = h.appID
	amount = h.amount
	purchaseDate = h.purchaseDate
	return
}
