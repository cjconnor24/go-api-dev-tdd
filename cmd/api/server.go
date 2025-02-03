package main

import "github.com/gin-gonic/gin"

type server struct {
	router *gin.Engine
}

func newServer() *server {
	return &server{}
}
func (s *server) routes() *gin.Engine {
	if s.router == nil {
		s.setupRoutes()
	}

	return s.router
}

func (s *server) run(addr string) error {
	return s.router.Run(addr)
}
