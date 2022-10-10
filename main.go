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
	u := &User{Name: "test"}
	if err = driver.Create(u); err != nil {
		fmt.Println(err)
	}
	u2 := &User{}
	driver.QuerySingle(u2, "name=?", "test")
	fmt.Println(u2.Name)
	driver.CloseDB()
	server.StartAndListen(":80")
}
