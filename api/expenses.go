package api

import (
	db "expenses/db/sqlc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateExpenseRequest struct {
	UserID      int64      `json:"user_id" binding:"required,gte=1"`
	CategoryID  int64      `json:"category_id" binding:"required,gte=1"`
	Amount      float32    `json:"amount" binding:"required,gte=0"`
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
