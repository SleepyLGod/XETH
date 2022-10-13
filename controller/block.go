package controller

import (
	"XETH/config"
	"XETH/service"
	"XETH/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

//CreateBlock
func CreateBlock(c *gin.Context) {
	username, ok := c.GetPostForm("username")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}

	password, ok := c.GetPostForm("password")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}
	newPassword := utils.EncryMd5(password)

	phone := c.PostForm("phone")
	email := c.PostForm("email")
	state, _ := strconv.Atoi(c.PostForm("state"))

	ok = service.CreateBlockService(username, newPassword, phone, email, state)
	if !ok {
		config.Error(c, int(config.ApiCode.CREATEUSERFAILED), config.ApiCode.GetMessage(config.ApiCode.CREATEUSERFAILED))
		return
	}
	config.Success(c, nil)

}

//FetchBlocks
func GetBlocks(c *gin.Context) {
	userList := service.GetBlocksService()
	fmt.Println(userList)

	config.Success(c, userList)
}

//DeleteBlockById
func DeleteBlockById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}
	fmt.Println("-------" + id + "---------")

	newId, err := strconv.Atoi(id)
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

//GetBlockById
func GetBlockById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		config.Error(c, int(config.ApiCode.CONVERTFAILED), config.ApiCode.GetMessage(config.ApiCode.CONVERTFAILED))
		return
	}

	user := service.GetBlockByIdService(newId)

	config.Success(c, user)
}

//UpdateBlockById
func UpdateBlockById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		config.Error(c, int(config.ApiCode.CONVERTFAILED), config.ApiCode.GetMessage(config.ApiCode.CONVERTFAILED))
		return
	}

	username, ok := c.GetPostForm("username")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}

	password, ok := c.GetPostForm("password")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}

	newPassword := utils.EncryMd5(password)

	phone := c.PostForm("phone")
	email := c.PostForm("email")

	_, ok, err = service.UpdateBlockByIdService(newId, username, newPassword, phone, email)
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

//DisableBlockById
func DisableBlockById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id)
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

//EnableBlockById
func EnableBlockById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		config.Error(c, int(config.ApiCode.LCAKPARAMETERS), config.ApiCode.GetMessage(config.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id)
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
