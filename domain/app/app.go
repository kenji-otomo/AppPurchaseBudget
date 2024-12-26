package app

import (
	"time"
)

type App struct {
	id             *int64
	name           string
	lastPurchaseAt *time.Time
	createdAt      *time.Time
	updatedAt      *time.Time
}

func NewApp(id *int64, name string, lastPurchaseAt *time.Time) *App {
	return &App{
		id:             id,
		name:           name,
		lastPurchaseAt: lastPurchaseAt,
	}
}

func FromApp(id *int64, name string, lastPurchaseAt, createdAt, updatedAt *time.Time) *App {
	return &App{
		id:             id,
		name:           name,
		lastPurchaseAt: lastPurchaseAt,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}

type AppWithSum struct {
	app    *App
	amount int64
}

func NewAppWithSum(app *App, amount int64) *AppWithSum {
	return &AppWithSum{
		app:    app,
		amount: amount,
	}
}
