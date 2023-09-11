package server

import (
	"fmt"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"nothing/internal/pkg/database"
	"nothing/pkg/jwt"
	"nothing/pkg/response"
	"nothing/pkg/util"
	"strings"
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
		if f != nil && len(f) != 0 {
			s.Engine.Use(f...)
		}
	}
}

func WithDB(dsn string) Option {
	return func(s *Server) {
		database.NewDataBase(dsn)
	}
}

func CorsMiddlewareFilter() gin.HandlerFunc {
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

//type Router struct {
//	Path    string
//	Handler func(c *gin.Context)
//	Mark    bool
//}

var whites = []string{
	"/login",
	"/blog/api/post/batch/query",
	"/blog/api/setting/batch/query",
	"/blog/api/post/attachment/content/query",
	"/blog/api/post/partition/batch/query",
	"/blog/api/post/query",
	"/feed",
	"/index.xml",
	"/blog/index.xml",
	"/blog/feed",
}

func AuthHandleFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//if ctx.Request.Method == http.MethodGet || strings.Contains(ctx.Request.RequestURI, "login") {
		//	ctx.Next()
		//	return
		//}
		uri := strings.Split(ctx.Request.RequestURI, "?")[0]
		if util.ArraysStringContain(whites, uri) {
			ctx.Next()
			return
		}
		authHeader := ctx.GetHeader("Authorization")
		tokenString := ""
		if strings.HasPrefix(authHeader, "Bearer ") {
			// 提取令牌部分
			tokenString = authHeader[7:]
		}
		result, err := jwt.Verify(tokenString)
		if err != nil || result == false {
			ctx.JSON(http.StatusOK, response.ErrJwtToken.BuildData(err.(*jwt2.ValidationError).Error()))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
