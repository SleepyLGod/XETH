package model

import "gorm.io/gorm"

type TokenInfo struct {
	gorm.Model
	Id          uint32 `gorm:"column:id;AUTO_INCREMENT"`
	Address     string `gorm:"column:address"`
	Name        string `gorm:"column:name"`
	Symbol      string `gorm:"column:symbol"`
	TotalSupply string `gorm:"column:total_supply"`
	Decimal     int32  `gorm:"column:decimal"`
}
