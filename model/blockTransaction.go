package model

import (
	"gorm.io/gorm"
)

type BlockTransaction struct {
	gorm.Model
	Id                   int64  `gorm:"column:id;NOT NULL"`
	BlockNum             int64  `gorm:"column:block_num;NOT NULL"`
	Timestamp            int64  `gorm:"column:created_at;NOT NULL"`
	TransactionHash      string `gorm:"column:transaction_hash;NOT NULL"`
	From                 string `gorm:"column:from;NOT NULL"`
	To                   string `gorm:"column:to"`
	ToCreate             string `gorm:"column:to_create"`
	FromIsContract       int8   `gorm:"column:from_is_contract"`
	ToIsContract         int8   `gorm:"column:to_is_contract"`
	Value                int64  `gorm:"column:value"`
	GasLimit             int32  `gorm:"column:gas_limit"`
	GasPrice             int64  `gorm:"column:gas_price"`
	GasUsed              int32  `gorm:"column:gas_used"`
	CallingFunction      string `gorm:"column:calling_function"`
	IsError              string `gorm:"column:is_error"`
	Eip2718type          string `gorm:"column:eip_2718_type"`
	BaseFeePerGas        int64  `gorm:"column:base_fee_per_gas"`
	MaxFeePerGas         int64  `gorm:"column:max_fee_per_gas"`
	MaxPriorityFeePerGas int64  `gorm:"column:max_priority_per_gas"`
}
