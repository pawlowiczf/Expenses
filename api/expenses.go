package api

import (
	"context"
	db "expenses/db/sqlc"
	"expenses/token"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateExpenseRequest struct {
	UserID      int64      `json:"user_id" binding:"required,gte=1"`
	CategoryID  int64      `json:"category_id" binding:"required,gte=1"`
	Amount      float32    `json:"amount" binding:"required"`
	Description string     `json:"description" binding:"required,ascii"`
	Date        *time.Time `json:"date"`
}
type CreateExpenseResponse struct {
	db.Expense
}

func (server *Server) CreateExpense(ctx *gin.Context) {
	var req CreateExpenseRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payloadAny, exists := ctx.Get(authorizationPayload)
	payload := payloadAny.(*token.Payload)

	if !exists {
		err = fmt.Errorf("no authorization payload provided")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return	
	}

	if payload.UserID != req.UserID {
		err = fmt.Errorf("non authorized user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.CreateExpenseParams{
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
		Amount:      req.Amount,
		Description: req.Description,
	}
	if req.Date == nil {
		arg.Date = time.Now()
	} else {
		arg.Date = *req.Date
	}

	expense, err := server.store.CreateExpense(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, CreateExpenseResponse{
		Expense: expense,
	})
}

type GetExpensesRequest struct {
	UserID int64 `json:"user_id"`
}

type GetExpensesResponse struct {
	Expenses []TransformedExpense `json:"expenses"`
}

type TransformedExpense struct {
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

func (server *Server) GetExpenses(ctx *gin.Context) {
	var req GetExpensesRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	payloadAny, exists := ctx.Get(authorizationPayload)
	payload := payloadAny.(*token.Payload)

	fmt.Println(payload.UserID, req.UserID)

	if !exists {
		err = fmt.Errorf("no authorization payload provided")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return	
	}

	if payload.UserID != req.UserID {
		err = fmt.Errorf("non authorized user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	expenses, err := server.store.GetUserExpenses(ctx, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var response GetExpensesResponse
	for _, expense := range expenses {
		response.Expenses = append(response.Expenses, server.transformExpense(expense))
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getCategoryName(categoryID int64) string {
	name, err := server.store.GetCategoryName(context.Background(), categoryID)
	if err != nil {
		return ""
	}
	return name
}

func (server *Server) transformExpense(expense db.Expense) TransformedExpense {
	return TransformedExpense{
		Amount:      expense.Amount,
		Description: expense.Description,
		Category:    server.getCategoryName(expense.CategoryID),
	}
}