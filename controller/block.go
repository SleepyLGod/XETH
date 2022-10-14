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
	// TransactionCount, _ := strconv.Atoi(transactionCount)
	TransactionCount, _ := strconv.ParseInt(transactionCount, 10, 32)
	ok = service.CreateBlockService(BlockNum, Timestamp, int32(TransactionCount))
	if !ok {
		config.Error(c, int(config.ApiCode.CREATEDFAILED), config.ApiCode.GetMessage(config.ApiCode.CREATEDFAILED))
		return
	}
	config.Success(c, nil)
}

// CreateBlockWithDTO 建议使用json方式的请求
func CreateBlockWithDTO(c *gin.Context) {
	// // just for testings
	//fmt.Printf("c.Request.body: %v", string(data))
	//log.Printf("c.Request.Method: %v", c.Request.Method)
	//log.Printf("c.Request.ContentType: %v", c.ContentType())
	//log.Printf("c.Request.Body: %v", c.Request.Body)
	//err := c.Request.ParseForm()
	//if err != nil {
	//	return
	//}
	//log.Printf("c.Request.Form: %v", c.Request.PostForm)
	//for k, v := range c.Request.PostForm {
	//	log.Printf("k:%v\n", k)
	//	log.Printf("v:%v\n", v)
	//}
	//log.Printf("c.Request.ContentLength: %v", c.Request.ContentLength)
	//data, _ = ioutil.ReadAll(c.Request.Body)
	//log.Printf("c.Request.GetBody: %v", string(data))
	//
	//data, err = c.GetRawData()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Printf("data: %v\n", string(data))
	////把读过的字节流重新放到body
	//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	var requestBody DTO.CreateBlockDTO
	if err := c.BindJSON(&requestBody); err != nil {
		config.Error(c, int(config.ApiCode.CREATEDFAILED), config.ApiCode.GetMessage(config.ApiCode.CREATEDFAILED))
		return
	}
	ok := service.CreateBlockServiceWithDTO(requestBody)
	if !ok {
		config.Error(c, int(config.ApiCode.CREATEDFAILED), config.ApiCode.GetMessage(config.ApiCode.CREATEDFAILED))
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
	block := service.GetBlockByIdService(newId)
	config.Success(c, block)
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

func GetBlocksWithConstraints(c *gin.Context) {
	var cons []dbDriver.QueryConstraint
	if err := c.BindJSON(&cons); err != nil {
		config.Error(c, int(config.ApiCode.INVALIDPARAMS), config.ApiCode.GetMessage(config.ApiCode.INVALIDPARAMS))
		return
	}
	blockList := service.GetBlocksWithConstraintsService(cons)
	config.Success(c, blockList)
}
