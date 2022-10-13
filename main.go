package main

import (
	"XETH/config"
	"XETH/dbDriver"
	xethServer "XETH/server"
	"fmt"
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

	databaseLogin := config.DatabaseGenerator
	driver, err := dbDriver.NewMysqlDriver(databaseLogin().UsrName, databaseLogin().Psw, databaseLogin().Addr, databaseLogin().DbName, &User{})
	if err != nil {
		fmt.Println(err)
	}

	// testings
	u := &User{Name: "333"}
	if err = driver.Create(u); err != nil {
		fmt.Println(err)
	}
	var u2 []User
	con := make([]dbDriver.QueryConstraint, 0)

	con = append(con, dbDriver.QueryConstraint {
		FieldName: "name",
		Operator:  ">",
		Value:     "1",
	})
	con = append(con, dbDriver.QueryConstraint {
		FieldName: "name",
		Operator:  "!=",
		Value:     "test",
	})
	driver.Query(&u2, con)
	fmt.Println(len(u2))
	driver.CloseDB()
	server.StartAndListen(":88")
}
