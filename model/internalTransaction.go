package model

import "gorm.io/gorm"

type InternalTransaction struct {
	gorm.Model
	Id               int64  `gorm:"column:id;NOT NULL"`
	BlockNum         int64  `gorm:"column:block_num;NOT NULL"`
	Timestamp        int64  `gorm:"column:created_at;NOT NULL"`
	TransactionHash  string `gorm:"column:transaction_hash;NOT NULL"`
	TypeTraceAddress string `gorm:"column:type_trace_address;NOT NULL"`
	From             string `gorm:"column:from;NOT NULL"`
	To               string `gorm:"column:to"`
	FromIsContract   int8   `gorm:"column:from_is_contract"`
	ToIsContract     int8   `gorm:"column:to_is_contract"`
	Value            int64  `gorm:"column:value"`
	CallingFunction  string `gorm:"column:calling_function"`
	IsError          string `gorm:"column:is_error"`
}
