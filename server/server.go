package xethServer

import (
	"XETH/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	engine *gin.Engine
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

func NewServer() *Server {
	r := gin.Default()
	// 跨域
	r.Use(Cors())
	return &Server{r}
}

func (s *Server) StartAndListen(port string) {
	err := s.engine.Run(port)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route(method string, path string, controllerFunc gin.HandlerFunc) {
	if method == "GET" {
		s.engine.GET(path, controllerFunc)
	} else if method == "POST" {
		s.engine.POST(path, controllerFunc)
	} else if method == "PUT" {
		s.engine.PUT(path, controllerFunc)
	} else if method == "DELETE" {
		s.engine.DELETE(path, controllerFunc)
	} else {
		panic(fmt.Sprintf("unsupported http method: %s", method))
	}
}

// RouterInit 路由生成函数！
func RouterInit() {
	server := NewServer()

	//server.Route("GET", "/", func(context *gin.Context) {
	//	context.String(http.StatusOK, "Hello XETH!")
	//	return
	//})

	// block 路由组
	blockGroup := server.engine.Group("/api/block/v1")
	blockGroup.Use(Cors())
	{
		blockGroup.POST("/create-block", controller.CreateBlock)
		blockGroup.GET("/get-blocks", controller.GetBlocks)
		blockGroup.POST("/del-block-by-id", controller.DeleteBlockById)
		blockGroup.POST("/get-block-by-id", controller.GetBlockById)
		blockGroup.POST("/update-block-by-id", controller.UpdateBlockById)
		blockGroup.POST("/disable-block-by-id", controller.DisableBlockById)
		blockGroup.POST("/enable-block-by-id", controller.EnableBlockById)
	}

	// transaction 路由组

	// internal_transaction 路由组

	// ...其余路由组

	// run
	server.StartAndListen(":88")
}
