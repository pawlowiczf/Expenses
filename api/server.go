package api

import (
	"expenses/config"
	db "expenses/db/sqlc"
	"expenses/token"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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

	maker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server.maker = maker
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	server.SetupRoutes(router)
	server.router = router



	return server, nil
}

func (server *Server) SetupRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // lub "*" aby pozwolić na dostęp ze wszystkich domen
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // Cache ustawień CORS na 12 godzin
	}))
	router.GET("/", server.Index)
	router.GET("/v1/check_auth", server.CheckAuth)

	router.POST("/v1/create_user", server.CreateUser)
	router.POST("/v1/login_user", server.LoginUser)

	authRoutes := router.Group("/").Use(authMiddlewareCookie(server.maker))
	authRoutes.POST("/v1/create_expense", server.CreateExpense)
	authRoutes.POST("/v1/get_expenses", server.GetExpenses)
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
