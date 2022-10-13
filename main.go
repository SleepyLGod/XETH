package main

import (
	xethServer "XETH/server"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	// 此处使用py脚本，注意首先配置好py的工作目录为./pkg

	server := xethServer.NewServer()
	server.Route("GET", "/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello XETH!")
		return
	})

	server.StartAndListen(":88")
}
