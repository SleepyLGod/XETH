package controller

import (
	"XETH/config"
	"XETH/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateBlock(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}
	timestamp, ok := c.GetPostForm("timestamp")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}
	transactionCount := c.PostForm("transactionCount")
	BlockNum, _ := strconv.ParseInt(id, 10, 64)
	Timestamp, _ := strconv.ParseInt(timestamp, 10, 64)
	//TransactionCount, _ := strconv.Atoi(transactionCount)
	TransactionCount, _ := strconv.ParseInt(transactionCount, 10, 32)
	ok = service.CreateBlockService(BlockNum, Timestamp, int32(TransactionCount))
	if !ok {
		config.Error(c, int(config.ApiCode.CREATEUSERFAILED), config.ApiCode.GetMessage(config.ApiCode.CREATEUSERFAILED))
		return
	}
	config.Success(c, nil)

}

func GetBlocks(c *gin.Context) {
	blockList := service.GetBlocksService()
	fmt.Println(blockList)
	config.Success(c, blockList)
}

func DeleteBlockById(c *gin.Context) {
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
	err, ok = service.DeleteBlockByIdService(newId)
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

func GetBlockById(c *gin.Context) {
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
	user := service.GetBlockByIdService(newId)
	config.Success(c, user)
}

func UpdateBlockById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}
	BlockNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		config.Error(c, int(config.ApiCode.CONVERTFAILED), config.ApiCode.GetMessage(config.ApiCode.CONVERTFAILED))
		return
	}
	timestamp, ok := c.GetPostForm("timestamp")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}
	transactionCount, ok := c.GetPostForm("transactionCount")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}
	Timestamp, _ := strconv.ParseInt(timestamp, 10, 64)
	TransactionCount, _ := strconv.ParseInt(transactionCount, 10, 32)
	_, ok, err = service.UpdateBlockByIdService(BlockNum, Timestamp, int32(TransactionCount))
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

func DisableBlockById(c *gin.Context) {
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
	_, err = service.DisableBlockByIdService(newId)
	if err != nil {
		config.Error(c, int(config.ApiCode.NOSUCHID), config.ApiCode.GetMessage(config.ApiCode.NOSUCHID))
		return
	}
	config.Success(c, nil)
}

func EnableBlockById(c *gin.Context) {
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
	_, err = service.EnableBlockByIdService(newId)
	if err != nil {
		config.Error(c, int(config.ApiCode.NOSUCHID), config.ApiCode.GetMessage(config.ApiCode.NOSUCHID))
		return
	}
	config.Success(c, nil)
}
