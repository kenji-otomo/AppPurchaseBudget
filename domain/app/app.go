package domain

import "time"

type app struct {
	id        *int64
	name      string
	createdAt *time.Time
	updatedAt *time.Time
}

func NewApp(name string) *app {
	return &app{
		name: name,
	}
}

func FromApp(id *int64, name string, createdAt, updatedAt *time.Time) *app {
	return &app{
		id:        id,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
