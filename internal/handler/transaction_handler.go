package handler

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
)

type UC interface {
    Create(ctx context.Context, userID int64, tenor int, amount float64) error
}

type TransactionHandler struct {
    uc UC
}

func NewTransactionHandler(uc UC) *TransactionHandler {
    return &TransactionHandler{uc: uc}
}

func (h *TransactionHandler) Create(c *gin.Context) {
    var req struct {
        UserID int64   `json:"user_id" binding:"required"`
        Tenor  int     `json:"tenor" binding:"required"`
        Amount float64 `json:"amount" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, err)
        return
    }

    err := h.uc.Create(c.Request.Context(), req.UserID, req.Tenor, req.Amount)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "limit exceeded"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "success"})
}
