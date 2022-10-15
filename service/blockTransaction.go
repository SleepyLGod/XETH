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
		TimeStamp:            createBlockTxn.TimeStamp,
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

func GetBlockTransactionsService() []dto.CreateBlockTransactionDTO {
	var schema []*model.BlockTransaction
	core.GetDB().Find(&schema)
	return dto.BlockTransactionDTOS(schema)
}

func GetBlockTransactionsWithConstraintsService(constraints []core.QueryConstraint) []dto.CreateBlockTransactionDTO {
	var schema []*model.BlockTransaction
	core.QueryWithDb(core.GetDB(), &schema, constraints)
	return dto.BlockTransactionDTOS(schema)
}

func DeleteBlockTransactionByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.BlockTransaction{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.BlockTransaction{}, id).Error
	return err, true
}

func GetBlockTransactionByIdService(id int64) dto.CreateBlockTransactionDTO {
	var schema *model.BlockTransaction
	core.GetDB().First(&schema, id)
	return dto.BlockTransactionDTOFromGorm(schema)
}
