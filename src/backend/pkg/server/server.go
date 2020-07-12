package server

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *gin.Engine
}

func Get() *Server {
	return &Server{}
}

func (s *Server) ConnectRouter(router *gin.Engine) *Server {
	s.server = router
	return s
}

func (s *Server) Start(serverPort string) error {
	if serverPort == "" {
		return errors.New("Server missing port")
	}

	if len(s.server.Routes()) <= 0 {
		return errors.New("Server missing address")
	}

	return s.server.Run(serverPort)
}
