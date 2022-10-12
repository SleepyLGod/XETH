package main

import (
	"XETH/dbDriver"
	xethServer "XETH/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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
	driver, err := dbDriver.NewMysqlDriver("root", "uniquefranky", "localhost", "xeth", &User{})
	if err != nil {
		fmt.Println(err)
	}

	//testings
	u := &User{Name: "333"}
	if err = driver.Create(u); err != nil {
		fmt.Println(err)
	}
	var u2 []User
	con := make([]dbDriver.QueryConstraint, 0)

	con = append(con, dbDriver.QueryConstraint{
		FieldName: "name",
		Operator:  ">",
		Value:     "1",
	})
	con = append(con, dbDriver.QueryConstraint{
		FieldName: "name",
		Operator:  "!=",
		Value:     "test",
	})
	driver.Query(&u2, con)
	fmt.Println(len(u2))
	driver.CloseDB()
	server.StartAndListen(":80")
}
