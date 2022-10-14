package service

import (
	dto "XETH/DTO"
	core "XETH/dbDriver"
	"XETH/model"
)

// CreateERC721TransactionServiceWithDTO 推荐
func CreateERC721TransactionServiceWithDTO(createERC721Txn dto.CreateERC721TransactionDTO) bool {
	db := core.GetDB()
	erc721 := model.ERC721Transaction{
		Id:              createERC721Txn.Id,
		TimeStamp:       createERC721Txn.TimeStamp,
		BlockNum:        createERC721Txn.BlockNum,
		TransactionHash: createERC721Txn.TransactionHash,
		TokenAddress:    createERC721Txn.TokenAddress,
		From:            createERC721Txn.From,
		To:              createERC721Txn.To,
		FromIsContract:  createERC721Txn.FromIsContract,
		ToIsContract:    createERC721Txn.ToIsContract,
		TokenId:         createERC721Txn.TokenId,
	}
	tx := db.Create(&erc721)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func GetERC721TransactionsService() (erc721TxnList []*model.ERC721Transaction) {
	core.GetDB().Find(&erc721TxnList)
	return
}

func GetERC721TransactionsWithConstraintsService(constraints []core.QueryConstraint) (erc721TxnList []*model.ERC721Transaction) {
	core.QueryWithDb(core.GetDB(), &erc721TxnList, constraints)
	return
}

func DeleteERC721TransactionByIdService(id int64) (err error, ok bool) {
	err = core.GetDB().First(&model.ERC721Transaction{}, id).Error
	if err != nil {
		return err, false
	}
	err = core.GetDB().Delete(&model.ERC721Transaction{}, id).Error
	return err, true
}

func GetERC721TransactionByIdService(id int64) (erc20Txn []*model.ERC721Transaction) {
	core.GetDB().First(&erc20Txn, id)
	return
}
