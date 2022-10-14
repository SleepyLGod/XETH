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
		BlockNum:         createItnTxn.BlockNum,
		Timestamp:        createItnTxn.Timestamp,
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

func GetInternalTransactionsService() (itnTxnList []*model.InternalTransaction) {
	core.GetDB().Find(&itnTxnList)
	return
}

func GetInternalTransactionsWithConstraintsService(constraints []core.QueryConstraint) (itnTxnList []*model.InternalTransaction) {
	core.QueryWithDb(core.GetDB(), &itnTxnList, constraints)
	return
}

func DeleteInternalTransactionByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.InternalTransaction{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.InternalTransaction{}, id).Error
	return err, true
}

func GetInternalTransactionByIdService(id int64) (itnTxn []*model.InternalTransaction) {
	core.GetDB().First(&itnTxn, id)
	return
}
