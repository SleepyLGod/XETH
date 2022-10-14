package main

import (
	server "XETH/server"
	"XETH/sql"
)

func main() {
	// 此处使用py脚本，注意首先配置好py的工作目录为./pkg

	err := sql.InsertBlockFromFile("/Users/franky/Downloads/1000000to1999999_Block_Info.csv", "/Users/franky/Downloads/1000000to1999999_Block_MinerReward.csv")

	if err != nil {
		panic(err)
	}
	server.RouterInit()
}
