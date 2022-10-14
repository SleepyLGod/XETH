package main

import (
	"XETH/DTO"
	server "XETH/server"
	"encoding/json"
	"os"
)

func main() {
	// 此处使用py脚本，注意首先配置好py的工作目录为./pkg
	test()
	server.RouterInit()
}

func test() {
	file, err := os.OpenFile("1.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	block := DTO.CreateBlockDTO{}
	json.NewEncoder(file).Encode(block)
}
