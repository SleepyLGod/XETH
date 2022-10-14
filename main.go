package main

import (
	server "XETH/server"
	"XETH/sql"
	"fmt"
)

func main() {
	// 此处使用py脚本，注意首先配置好py的工作目录为./pkg
	err := sql.InsertBlockInfoFromFile("/Users/franky/Downloads/1000000to1999999_Block_Info.csv")
	if err != nil {
		fmt.Println(err)
	}
	server.RouterInit()
}
