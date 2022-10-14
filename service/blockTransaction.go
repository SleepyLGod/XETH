package service

import (
	dto "XETH/DTO"
	core "XETH/dbDriver"
	"XETH/model"
)

// CreateBlockTransactionServiceWithDTO 推荐
func CreateBlockTransactionServiceWithDTO(createBlockTxn dto.CreateBlockTransactionDTO) bool {
	db := core.GetDB()
	blockTxn := model.BlockTransaction{
		Id:                   createBlockTxn.Id,
		BlockNum:             createBlockTxn.BlockNum,
		Timestamp:            createBlockTxn.Timestamp,
		TransactionHash:      createBlockTxn.TransactionHash,
		From:                 createBlockTxn.From,
		To:                   createBlockTxn.To,
		ToCreate:             createBlockTxn.ToCreate,
		FromIsContract:       createBlockTxn.FromIsContract,
		ToIsContract:         createBlockTxn.ToIsContract,
		Value:                createBlockTxn.Value,
		GasLimit:             createBlockTxn.GasLimit,
		GasPrice:             createBlockTxn.GasPrice,
		GasUsed:              createBlockTxn.GasUsed,
		CallingFunction:      createBlockTxn.CallingFunction,
		IsError:              createBlockTxn.IsError,
		Eip2718type:          createBlockTxn.Eip2718type,
		BaseFeePerGas:        createBlockTxn.BaseFeePerGas,
		MaxFeePerGas:         createBlockTxn.MaxFeePerGas,
		MaxPriorityFeePerGas: createBlockTxn.MaxPriorityFeePerGas,
	}
	tx := db.Create(&blockTxn)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func GetBlockTransactionsService() (blockTxnList []*model.BlockTransaction) {
	core.GetDB().Find(&blockTxnList)
	return blockTxnList
}

func GetBlockTransactionsWithConstraintsService(constraints []core.QueryConstraint) (blockTxnList []*model.BlockTransaction) {
	core.QueryWithDb(core.GetDB(), &blockTxnList, constraints)
	return blockTxnList
}

func DeleteBlockTransactionByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.BlockTransaction{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.BlockTransaction{}, id).Error
	return err, true
}

func GetBlockTransactionByIdService(id int64) (blockTxn []*model.BlockTransaction) {
	core.GetDB().First(&blockTxn, id)
	return
}
