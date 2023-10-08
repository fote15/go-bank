package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "gobank/db/sqlc"
	"net/http"
)

type UpdateAccountRequest struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	var req UpdateAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, errorMessage(err))
		return
	}
	args := db.UpdateAccountParams{
		ID:      req.ID,
		Balance: req.Balance,
	}
	acc, err := server.store.UpdateAccount(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorMessage(err))
			return
		}
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, acc)
	return
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req GetAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, errorMessage(err))
		return
	}
	acc, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorMessage(err))
			return
		}
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, acc)
	return
}

type ListAccountRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	LIMIT  int32 `form:"limit" binding:"required,min=1,max=100"`
}

func (server *Server) listAccount(ctx *gin.Context) {
	var req ListAccountRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, errorMessage(err))
		return
	}
	arg := db.ListAccountsParams{
		Offset: (req.PageID - 1) * req.LIMIT,
		Limit:  req.LIMIT,
	}
	acc, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorMessage(err))
			return
		}
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}
	if len(acc) < 1 {
		ctx.JSON(http.StatusInternalServerError, "no account here")
		return
	}
	ctx.JSON(http.StatusOK, acc)
	return
}

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=KZT USD"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMessage(err))
		return
	}
	args := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMessage(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}
