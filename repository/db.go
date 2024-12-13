package repository

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func NewDB(createdDB *gorm.DB) {
	db = createdDB
}

func BeginTransaction() *gorm.DB {
	return db.Begin()
}
