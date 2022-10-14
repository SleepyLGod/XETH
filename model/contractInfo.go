package model

import "gorm.io/gorm"

type ContractInfo struct {
	gorm.Model
	Id                     uint32 `gorm:"column:id;AUTO_INCREMENT"`
	CreatedBlock           int64  `gorm:"column:created_block"`
	CreatedTimeStamp       int64  `gorm:"column:created_at"`
	Address                string `gorm:"column:address"`
	CreatedTransactionHash string `gorm:"column:created_transaction_hash"`
	Creator                string `gorm:"column:creator"`
	CreatorIsContract      int8   `gorm:"column:creator_is_contract"`
	CreateValue            int64  `gorm:"column:create_value"`
	CreationCode           string `gorm:"column:creation_code"`
	ContractCode           string `gorm:"column:contract_code"`
}
