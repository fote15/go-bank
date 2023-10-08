package api

import (
	"github.com/gin-gonic/gin"
	db "gobank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.POST("/account-update", server.updateAccount)
	router.GET("/accounts", server.listAccount)
	server.router = router
	return server
}

func errorMessage(err error) gin.H {
	return gin.H{"error": err.Error()}
}
