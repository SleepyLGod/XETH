package service

import (
	core "XETH/dbDriver"
	model "XETH/model"
	constants "XETH/utils"
	"fmt"
	"time"
)

// gorm.Model
// blockNum                int64     `gorm:"column:id;NOT NULL"`
// timestamp               time.Time `gorm:"column:created_at;NOT NULL"`
// transactionCount        int32     `gorm:"column:transaction_count;NOT NULL"`
// internalTransctionCount int32     `gorm:"column:internal_transaction_count"`
// minerAddress            string    `gorm:"column:miner_address"`
// blockReward             int64     `gorm:"column:block_reward"`
// unclesReward            string    `gorm:"column:uncles_reward"` // TODO: 待检测string类型是否正确
// difficulty              string    `gorm:"column:difficulty"`
// totalDifficulty         string    `gorm:"column:total_difficulty"`
// size                    int32     `gorm:"column:size"`
// gasUsed                 int32     `gorm:"column:gas_used"`
// gasLimit                int32     `gorm:"column:gas_limit"`
// baseFeePerGas           int64     `gorm:"column:base_fee_per_gas"`
// burntFees               int64     `gorm:"column:burnt_fees"`
// extraData               string    `gorm:"column:extra_data"`
// hash                    string    `gorm:"column:hash"`
// parentHash              string    `gorm:"column:parent_hash"`
// sha3Uncles              string    `gorm:"column:sha3_uncles"`
// stateRoot               string    `gorm:"column:state_root"`
// nounce                  string    `gorm:"column:nounce"`

// CreateBlockService
func CreateBlockService(blockNum int64, timestamp time.Time, transactionCount int32) bool {
	db := core.GetDB()
	block := model.Block{
		BlockNum:         blockNum,
		Timestamp:        timestamp,
		TransactionCount: transactionCount,
	}
	tx := db.Create(&block)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// GetBlocksService
func GetBlocksService() (blockList []*model.Block) {
	core.GetDB().Find(&blockList)
	return blockList
}

// DeleteBlockByIdService
func DeleteBlockByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.Block{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.Block{}, id).Error
	return err, true
}

// GetBlockByIdService
func GetBlockByIdService(id int64) (block []*model.Block) {
	core.GetDB().First(&block, id)
	return
}

// UpdateBlockByIdService
func UpdateBlockByIdService(id int64, username, timestamp time.Time, transactionCount int32) (block model.Block, ok bool, err error) {
	err = core.GetDB().First(&block, id).Error
	if err != nil {
		return block, false, err
	}
	block.BlockNum = id
	block.Timestamp = timestamp
	block.TransactionCount = transactionCount
	err = core.GetDB().Save(&block).Error
	if err != nil {
		return block, false, err
	}
	return block, true, nil
}

// 考虑到逻辑删除和恢复，这里的禁用操作是将state字段设置为0
// DisableBlockByIdService
func DisableBlockByIdService(id int64) (block model.Block, err error) {
	err = core.GetDB().First(&block, id).Error
	fmt.Println(block)
	if err != nil {
		return block, err
	}
	core.GetDB().Model(&block).Update("state", constants.DISABLE)
	return block, nil
}

// EnableBlockByIdService
func EnableBlockByIdService(id int) (block model.Block, err error) {
	err = core.GetDB().First(&block, id).Error
	fmt.Println(block)
	if err != nil {
		return block, err
	}
	core.GetDB().Model(&block).Update("state", constants.ENABLE)
	return block, nil
}
