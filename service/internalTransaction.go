package service

import (
	dto "XETH/DTO"
	core "XETH/dbDriver"
	"XETH/model"
)

// CreateInternalTransactionServiceWithDTO 推荐
func CreateInternalTransactionServiceWithDTO(createItnTxn dto.CreateInternalTransactionDTO) bool {
	db := core.GetDB()
	itnTxn := model.InternalTransaction{
		Id:               createItnTxn.Id,
		BlockNum:         createItnTxn.BlockNum,
		TimeStamp:        createItnTxn.TimeStamp,
		TransactionHash:  createItnTxn.TransactionHash,
		TypeTraceAddress: createItnTxn.TypeTraceAddress,
		From:             createItnTxn.From,
		To:               createItnTxn.To,
		FromIsContract:   createItnTxn.FromIsContract,
		ToIsContract:     createItnTxn.ToIsContract,
		Value:            createItnTxn.Value,
		CallingFunction:  createItnTxn.CallingFunction,
		IsError:          createItnTxn.IsError,
	}
	tx := db.Create(&itnTxn)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func GetInternalTransactionsService() []dto.CreateInternalTransactionDTO {
	var schema []*model.InternalTransaction
	core.GetDB().Find(&schema)
	return dto.InternalTransactionDTOS(schema)
}

func GetInternalTransactionsWithConstraintsService(constraints []core.QueryConstraint) []dto.CreateInternalTransactionDTO {
	var schema []*model.InternalTransaction
	core.QueryWithDb(core.GetDB(), &schema, constraints)
	return dto.InternalTransactionDTOS(schema)
}

func DeleteInternalTransactionByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.InternalTransaction{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.InternalTransaction{}, id).Error
	return err, true
}

func GetInternalTransactionByIdService(id int64) dto.CreateInternalTransactionDTO {
	var schema *model.InternalTransaction
	core.GetDB().First(&schema, id)
	return dto.InternalTransactionDTOFromGorm(schema)
}
