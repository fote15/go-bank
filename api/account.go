package api

import (
	"github.com/gin-gonic/gin"
	db "gobank/db/sqlc"
	"net/http"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding: "required"`
	Currency string `json:"currency" binding: "required, oneof: KZT USD"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMessage(err))
	}
	args := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMessage(err))
	}
	ctx.JSON(http.StatusOK, account)
}
