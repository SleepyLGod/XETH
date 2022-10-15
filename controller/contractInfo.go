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

func CreateContractInfoWithDTO(c *gin.Context) {
	var requestBody DTO.CreateContractInfoDTO
	if err := c.BindJSON(&requestBody); err != nil {
		config.Error(c, int(config.ApiCode.CREATEDFAILED), config.ApiCode.GetMessage(config.ApiCode.CREATEDFAILED))
		return
	}
	ok := service.CreateContractInfoServiceWithDTO(requestBody)
	if !ok {
		config.Error(c, int(config.ApiCode.CREATEDFAILED), config.ApiCode.GetMessage(config.ApiCode.CREATEDFAILED))
		return
	}
	config.Success(c, nil)
}

func GetContractInfos(c *gin.Context) {
	contractInfoList := service.GetContractInfosService()
	fmt.Println(contractInfoList)
	config.Success(c, contractInfoList)
}

func DeleteContractInfoById(c *gin.Context) {
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
	err, ok = service.DeleteContractInfoByIdService(newId)
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

func GetContractInfoById(c *gin.Context) {
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
	contractInfo := service.GetContractInfoByIdService(newId)
	config.Success(c, contractInfo)
}

func GetContractInfoWithConstraints(c *gin.Context) {
	var cons []dbDriver.QueryConstraint
	if err := c.BindJSON(&cons); err != nil {
		config.Error(c, int(config.ApiCode.INVALIDPARAMS), config.ApiCode.GetMessage(config.ApiCode.INVALIDPARAMS))
		return
	}
	contractInfoList := service.GetContractInfosWithConstraintsService(cons)
	config.Success(c, contractInfoList)
}
