package app

import (
	"time"
)

type app struct {
	id             *int64
	name           string
	lastPurchaseAt *time.Time
	createdAt      *time.Time
	updatedAt      *time.Time
}

func NewApp(name string) *app {
	return &app{
		name: name,
	}
}

func FromApp(id *int64, name string, lastPurchaseAt, createdAt, updatedAt *time.Time) *app {
	return &app{
		id:             id,
		name:           name,
		lastPurchaseAt: lastPurchaseAt,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}
