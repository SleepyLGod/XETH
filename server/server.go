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
		blockGroup.POST("/create-block", controller.CreateBlockWithDTO)
		blockGroup.GET("/get-blocks", controller.GetBlocks)
		/*
			删除后在数据库中仍存在，只是deleted_at被设置为删除时间，从而用gorm查询不到。
			但是再次插入相同id的数据时，会插入失败。
		*/
		blockGroup.POST("/del-block-by-id", controller.DeleteBlockById)
		blockGroup.POST("/get-block-by-id", controller.GetBlockById)
		blockGroup.POST("/update-block-by-id", controller.UpdateBlockById)
		blockGroup.POST("/disable-block-by-id", controller.DisableBlockById)
		blockGroup.POST("/enable-block-by-id", controller.EnableBlockById)
		blockGroup.POST("/get-blocks-with-constraints", controller.GetBlocksWithConstraints)
	}

	// transaction 路由组
	txnGroup := server.engine.Group("/api/blockTxn/v1")
	txnGroup.Use(Cors())
	{
		txnGroup.POST("/create-txn", controller.CreateBlockTransactionWithDTO)
		txnGroup.GET("/get-txns", controller.GetBlockTransactions)
		txnGroup.POST("/del-txn-by-id", controller.DeleteBlockTransactionById)
		txnGroup.POST("/get-txn-by-id", controller.GetBlockTransactionById) // 当ID不存在时，会返回"0"Txn
		txnGroup.POST("/get-txns-with-constraints", controller.GetBlockTransactionsWithConstraints)
	}

	// internal_transaction 路由组
	itnTxnGroup := server.engine.Group("/api/itnTxn/v1")
	itnTxnGroup.Use(Cors())
	{
		itnTxnGroup.POST("/create-txn", controller.CreateInternalTransactionWithDTO)
		itnTxnGroup.GET("/get-txns", controller.GetInternalTransactions)
		itnTxnGroup.POST("/del-txn-by-id", controller.DeleteInternalTransactionById)
		itnTxnGroup.POST("/get-txn-by-id", controller.GetInternalTransactionById) // 当ID不存在时，会返回"0"Txn
		itnTxnGroup.POST("/get-txns-with-constraints", controller.GetInternalTransactionsWithConstraints)
	}

	// contract_info 路由组
	cttInfoGroup := server.engine.Group("/api/cttInfo/v1")
	cttInfoGroup.Use(Cors())
	{
		cttInfoGroup.POST("/create-ctt", controller.CreateContractInfoWithDTO)
		cttInfoGroup.GET("/get-ctts", controller.GetContractInfos)
		cttInfoGroup.POST("/del-ctt-by-id", controller.DeleteContractInfoById)
		cttInfoGroup.POST("/get-ctt-by-id", controller.GetContractInfoById) // 当ID不存在时，会返回"0"Ctt
		cttInfoGroup.POST("/get-ctts-with-constraints", controller.GetContractInfoWithConstraints)
	}

	// erc20 路由组
	erc20Group := server.engine.Group("/api/erc20/v1")
	erc20Group.Use(Cors())
	{
		erc20Group.POST("/create-erc20", controller.CreateERC20TransactionWithDTO)
		erc20Group.GET("/get-erc20s", controller.GetERC20Transactions)
		erc20Group.POST("/del-erc20-by-id", controller.DeleteERC20TransactionById)
		erc20Group.POST("/get-erc20-by-id", controller.GetERC20TransactionById) // 当ID不存在时，会返回"0"ERC20
		erc20Group.POST("/get-erc20s-with-constraints", controller.GetERC20TransactionsWithConstraints)
	}

	// erc721 路由组
	erc721Group := server.engine.Group("/api/erc721/v1")
	erc721Group.Use(Cors())
	{
		erc721Group.POST("/create-erc721", controller.CreateERC721TransactionWithDTO)
		erc721Group.GET("/get-erc721s", controller.GetERC721Transactions)
		erc721Group.POST("/del-erc721-by-id", controller.DeleteERC721TransactionById)
		erc721Group.POST("/get-erc721-by-id", controller.GetERC721TransactionById)
		erc721Group.POST("/get-erc721s-with-constraints", controller.GetERC721TransactionsWithConstraints)
	}

	// token_info 路由组
	tokenGroup := server.engine.Group("/api/token/v1")
	tokenGroup.Use(Cors())
	{
		tokenGroup.POST("/create-token", controller.CreateTokenInfoWithDTO)
		tokenGroup.GET("/get-tokens", controller.GetTokenInfos)
		tokenGroup.POST("/del-token-by-id", controller.DeleteTokenInfoById)
		tokenGroup.POST("/get-token-by-id", controller.GetTokenInfoById)
		tokenGroup.POST("/get-tokens-with-constraints", controller.GetTokenInfosWithConstraints)
	}

	// run
	server.StartAndListen(":88")
}
