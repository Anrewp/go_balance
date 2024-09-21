package model

import (
	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	ID          int `gorm:"primarey_key"`
	AmountCents int64
}
