package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Id int `gorm:"primarey_key" sql:"AUTO_INCREMENT"`
	//BalanceID int
	//Balance   Balance
	Email string
}
