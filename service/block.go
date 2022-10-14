package service

import (
	dto "XETH/DTO"
	core "XETH/dbDriver"
	model "XETH/model"
	constants "XETH/utils"
	"fmt"
)

// CreateBlockService 不推荐，只是适合参数特别少的传参方式，但是为了统一格式建议统一使用下面一种方法
func CreateBlockService(blockNum int64, timestamp int64, transactionCount int32) bool {
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

// CreateBlockServiceWithDTO 推荐
func CreateBlockServiceWithDTO(CreateBlockDTO dto.CreateBlockDTO) bool {
	db := core.GetDB()
	block := model.Block{
		BlockNum:                CreateBlockDTO.BlockNum,
		Timestamp:               CreateBlockDTO.Timestamp,
		TransactionCount:        CreateBlockDTO.TransactionCount,
		InternalTransctionCount: CreateBlockDTO.InternalTransctionCount,
		MinerAddress:            CreateBlockDTO.MinerAddress,
		BlockReward:             CreateBlockDTO.BlockReward,
		UnclesReward:            CreateBlockDTO.UnclesReward,
		Difficulty:              CreateBlockDTO.Difficulty,
		TotalDifficulty:         CreateBlockDTO.TotalDifficulty,
		Size:                    CreateBlockDTO.Size,
		GasUsed:                 CreateBlockDTO.GasUsed,
		GasLimit:                CreateBlockDTO.GasLimit,
		BaseFeePerGas:           CreateBlockDTO.BaseFeePerGas,
		BurntFees:               CreateBlockDTO.BurntFees,
		ExtraData:               CreateBlockDTO.ExtraData,
		Hash:                    CreateBlockDTO.Hash,
		ParentHash:              CreateBlockDTO.ParentHash,
		Sha3Uncles:              CreateBlockDTO.Sha3Uncles,
		StateRoot:               CreateBlockDTO.StateRoot,
		Nounce:                  CreateBlockDTO.Nounce,
	}
	tx := db.Create(&block)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func GetBlocksService() (blockList []*model.Block) {
	core.GetDB().Find(&blockList)
	return blockList
}

func GetBlocksWithConstraintsService(constraints []core.QueryConstraint) (blockList []*model.Block) {
	core.QueryWithDb(core.GetDB(), &blockList, constraints)
	return blockList
}

func DeleteBlockByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.Block{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.Block{}, id).Error
	return err, true
}

func GetBlockByIdService(id int64) (block []*model.Block) {
	core.GetDB().First(&block, id)
	return
}

func UpdateBlockByIdService(id int64, timestamp int64, transactionCount int32) (block model.Block, ok bool, err error) {
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

// DisableBlockByIdService 考虑到逻辑删除和恢复，这里的禁用操作是将state字段设置为0
func DisableBlockByIdService(id int64) (block model.Block, err error) {
	err = core.GetDB().First(&block, id).Error
	fmt.Println(block)
	if err != nil {
		return block, err
	}
	core.GetDB().Model(&block).Update("state", constants.DISABLE)
	return block, nil
}

// EnableBlockByIdService  考虑到逻辑删除和恢复，这里的操作是将state字段设置为1
func EnableBlockByIdService(id int64) (block model.Block, err error) {
	err = core.GetDB().First(&block, id).Error
	fmt.Println(block)
	if err != nil {
		return block, err
	}
	core.GetDB().Model(&block).Update("state", constants.ENABLE)
	return block, nil
}
