package model

import (
	core "XETH/dbDriver"
	"time"

	"gorm.io/gorm"
)

// 数据类型对应参考；https://segmentfault.com/a/1190000011016366
type Block struct {
	gorm.Model
	BlockNum                int64     `gorm:"column:id;NOT NULL"`
	Timestamp               time.Time `gorm:"column:created_at;NOT NULL"`
	TransactionCount        int32     `gorm:"column:transaction_count;NOT NULL"`
	InternalTransctionCount int32     `gorm:"column:internal_transaction_count"`
	MinerAddress            string    `gorm:"column:miner_address"`
	BlockReward             int64     `gorm:"column:block_reward"`
	UnclesReward            string    `gorm:"column:uncles_reward"` // TODO: 待检测string类型是否正确
	Difficulty              string    `gorm:"column:difficulty"`
	TotalDifficulty         string    `gorm:"column:total_difficulty"`
	Size                    int32     `gorm:"column:size"`
	GasUsed                 int32     `gorm:"column:gas_used"`
	GasLimit                int32     `gorm:"column:gas_limit"`
	BaseFeePerGas           int64     `gorm:"column:base_fee_per_gas"`
	BurntFees               int64     `gorm:"column:burnt_fees"`
	ExtraData               string    `gorm:"column:extra_data"`
	Hash                    string    `gorm:"column:hash"`
	ParentHash              string    `gorm:"column:parent_hash"`
	Sha3Uncles              string    `gorm:"column:sha3_uncles"`
	StateRoot               string    `gorm:"column:state_root"`
	Nounce                  string    `gorm:"column:nounce"`
}

// Block
func (block *Block) Block() string {
	return "block"
}

func init() {
	db := core.GetDB()
	db.AutoMigrate(&Block{})
}
