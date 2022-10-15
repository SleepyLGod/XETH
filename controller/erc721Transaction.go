package controller

import (
	"XETH/DTO"
	"XETH/config"
	"XETH/dbDriver"
	"XETH/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateERC721TransactionWithDTO(c *gin.Context) {
	var requestBody DTO.CreateERC721TransactionDTO
	if err := c.BindJSON(&requestBody); err != nil {
		config.Error(c, int(config.ApiCode.CREATEDFAILED), config.ApiCode.GetMessage(config.ApiCode.CREATEDFAILED))
		return
	}
	ok := service.CreateERC721TransactionServiceWithDTO(requestBody)
	if !ok {
		config.Error(c, int(config.ApiCode.CREATEDFAILED), config.ApiCode.GetMessage(config.ApiCode.CREATEDFAILED))
		return
	}
	config.Success(c, nil)
}

func GetERC721Transactions(c *gin.Context) {
	erc721TxnList := service.GetERC721TransactionsService()
	config.Success(c, erc721TxnList)
}

func DeleteERC721TransactionById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}
	fmt.Println("-------" + id + "---------")
	newId, err := strconv.ParseInt(id, 10, 64)
	fmt.Printf("%T", newId)
	if err != nil {
		config.Error(c, int(config.ApiCode.CONVERTFAILED), config.ApiCode.GetMessage(config.ApiCode.CONVERTFAILED))
		return
	}
	err, ok = service.DeleteERC721TransactionByIdService(newId)
	if !ok {
		config.Error(c, int(config.ApiCode.NOSUCHID), config.ApiCode.GetMessage(config.ApiCode.NOSUCHID))
		return
	}
	if err != nil {
		config.Error(c, int(config.ApiCode.FAILED), config.ApiCode.GetMessage(config.ApiCode.FAILED))
		return
	}
	config.Success(c, nil)
}

func GetERC721TransactionById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		config.Error(c, int(config.ApiCode.CONVERTFAILED), config.ApiCode.GetMessage(config.ApiCode.CONVERTFAILED))
		return
	}
	erc721Txn := service.GetERC721TransactionByIdService(newId)
	config.Success(c, erc721Txn)
}

func GetERC721TransactionsWithConstraints(c *gin.Context) {
	var cons []dbDriver.QueryConstraint
	if err := c.BindJSON(&cons); err != nil {
		config.Error(c, int(config.ApiCode.INVALIDPARAMS), config.ApiCode.GetMessage(config.ApiCode.INVALIDPARAMS))
		return
	}
	erc721List := service.GetERC721TransactionsWithConstraintsService(cons)
	config.Success(c, erc721List)
}
