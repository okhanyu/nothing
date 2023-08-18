package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nothing/internal/pkg/database"
)

type Option func(server *Server)

type Server struct {
	Name   string
	Engine *gin.Engine
}

func (s *Server) Start(port int) {
	p := fmt.Sprintf(":%d", port)
	err := s.Engine.Run(p)
	if err != nil {
		return
	}
}

func NewServer(name string, options ...Option) (*Server, error) {
	server := &Server{Name: name, Engine: gin.Default()}
	for _, option := range options {
		option(server)
	}
	return server, nil
}

func WithRouter(router ...func(r *gin.Engine)) Option {
	return func(s *Server) {
		for _, f := range router {
			f(s.Engine)
		}
	}
}

func WithFilter(f ...gin.HandlerFunc) Option {
	return func(s *Server) {
		FilterRegister(s.Engine, f...)
	}
}

func WithDB(dsn string) Option {
	return func(s *Server) {
		database.NewDataBase(dsn)
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
			"Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func FilterRegister(e *gin.Engine, f ...gin.HandlerFunc) {
	if f != nil && len(f) != 0 {
		e.Use(f...)
	}
}

//type Router struct {
//	Path    string
//	Handler func(c *gin.Context)
//	Mark    bool
//}
