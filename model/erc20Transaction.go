package model

import "gorm.io/gorm"

type ERC20Transaction struct {
	gorm.Model
	Id              uint32 `gorm:"column:id;AUTO_INCREMENT"`
	TimeStamp       int64  `gorm:"column:created_at;NOT NULL"`
	BlockNum        int64  `gorm:"column:block_num"`
	TransactionHash string `gorm:"column:transaction_hash"`
	TokenAddress    string `gorm:"column:token_address"`
	From            string `gorm:"column:from"`
	To              string `gorm:"column:to"`
	FromIsContract  int8   `gorm:"column:from_is_contract"`
	ToIsContract    int8   `gorm:"column:to_is_contract"`
	Value           int64  `gorm:"column:value"`
}
