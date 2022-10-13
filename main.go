package main

import (
	server "XETH/server"
)

func main() {
	// 此处使用py脚本，注意首先配置好py的工作目录为./pkg

	server.RouterInit()
}
