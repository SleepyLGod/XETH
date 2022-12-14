package service

import (
	dto "XETH/DTO"
	core "XETH/dbDriver"
	"XETH/model"
)

// CreateERC20TransactionServiceWithDTO 推荐
func CreateERC20TransactionServiceWithDTO(createERC20Txn dto.CreateERC20TransactionDTO) bool {
	db := core.GetDB()
	erc20 := model.ERC20Transaction{
		Id:              createERC20Txn.Id,
		TimeStamp:       createERC20Txn.TimeStamp,
		BlockNum:        createERC20Txn.BlockNum,
		TransactionHash: createERC20Txn.TransactionHash,
		TokenAddress:    createERC20Txn.TokenAddress,
		From:            createERC20Txn.From,
		To:              createERC20Txn.To,
		FromIsContract:  createERC20Txn.FromIsContract,
		ToIsContract:    createERC20Txn.ToIsContract,
		Value:           createERC20Txn.Value,
	}
	tx := db.Create(&erc20)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func GetERC20TransactionsService() []dto.CreateERC20TransactionDTO {
	var schema []*model.ERC20Transaction
	core.GetDB().Find(&schema)
	return dto.ERC20TransactionDTOS(schema)
}

func GetERC20TransactionsWithConstraintsService(constraints []core.QueryConstraint) []dto.CreateERC20TransactionDTO {
	var schema []*model.ERC20Transaction
	core.QueryWithDb(core.GetDB(), &schema, constraints)
	return dto.ERC20TransactionDTOS(schema)
}

func DeleteERC20TransactionByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.ERC20Transaction{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.ERC20Transaction{}, id).Error
	return err, true
}

func GetERC20TransactionByIdService(id int64) dto.CreateERC20TransactionDTO {
	var schema *model.ERC20Transaction
	core.GetDB().First(&schema, id)
	return dto.ERC20TransactionDTOFromGorm(schema)
}
