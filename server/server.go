package xethServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	return &Server{gin.Default()}
}

func (s *Server) StartAndListen(port string) {
	err := s.engine.Run(port)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route(method string, path string, handlerFunc gin.HandlerFunc) {
	if method == "GET" {
		s.engine.GET(path, handlerFunc)
	} else if method == "POST" {
		s.engine.POST(path, handlerFunc)
	} else if method == "PUT" {
		s.engine.PUT(path, handlerFunc)
	} else if method == "DELETE" {
		s.engine.DELETE(path, handlerFunc)
	} else {
		panic(fmt.Sprintf("unsupported http method: %s", method))
	}
}
