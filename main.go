package main

import (
	"XETH/config"
	"XETH/dbDriver"
	xethServer "XETH/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") // 请求头部
		if origin != "" {
			// 接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			// 服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// 允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			// 设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			// 允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		// 允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
			// 放行所有OPTIONS方法
			// c.AbortWithStatus(http.StatusNoContent)
		}
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()
		c.Next()
	}
}

func main() {
	// 此处使用py脚本，注意首先配置好py的工作目录为./pkg

	// 跨域
	r := gin.Default()
	r.Use(Cors())

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
	server.StartAndListen(":88")
}
