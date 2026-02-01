package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Endpoint simulasi sesuai input Anda
	r.POST("/transactions", func(c *gin.Context) {
		// Simulasi pengecekan Auth Bearer Token (OWASP Standard) 
		token := c.GetHeader("Authorization")
		if token != "Bearer testtoken" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		var jsonInput map[string]interface{}
		if err := c.ShouldBindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   jsonInput,
		})
	})

	t.Run("Success with valid token", func(t *testing.T) {
		// Data input sesuai permintaan Anda
		payload := map[string]interface{}{
			"user_id": 1,
			"tenor":   1,
			"amount":  500000,
		}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
		
		// Menambahkan Auth Bearer Token
		req.Header.Set("Authorization", "Bearer testtoken")
		req.Header.Set("Content-Type", "application/json")

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "success")
	})

	t.Run("Fail with missing token", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/transactions", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}