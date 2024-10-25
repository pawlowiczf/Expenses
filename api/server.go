package api

import (
	"expenses/config"
	db "expenses/db/sqlc"
	"expenses/token"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	config config.Config
	router *gin.Engine
	maker  token.TokenMaker 
}

func NewServer(store *db.Store, config config.Config) (*Server, error) {
	server := &Server{
		store:  store,
		config: config,
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	server.SetupRoutes(router)
	server.router = router

	maker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err 
	}
	server.maker = maker 

	return server, nil 
}

func (server *Server) SetupRoutes(router *gin.Engine) {
	router.GET("/", server.Index)

	router.POST("/v1/create_user", server.CreateUser)
	router.POST("/v1/create_expense", server.CreateExpense)
	router.POST("/v1/login_user", server.LoginUser)
}

func (server *Server) RunServer(address string) (err error) {
	return server.router.Run(address)
}

func TestServer(store *db.Store, config config.Config) *Server {
	server := &Server{
		store:  store,
		config: config,
	}
	router := gin.Default()
	server.router = router
	server.SetupRoutes(router)

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
