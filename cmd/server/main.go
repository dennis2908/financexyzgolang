
package main

import (
    "database/sql"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"

    "xyz-finance/internal/handler"
    "xyz-finance/internal/middleware"
    "xyz-finance/internal/repository"
    "xyz-finance/internal/usecase"
)

func main() {
    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        dsn = "root:root@tcp(db:3306)/xyz?parseTime=true"
    }

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    repo := repository.NewLimitRepository(db)
    uc := usecase.NewTransactionUsecase(repo)
    h := handler.NewTransactionHandler(uc)

    r := gin.New()
    r.Use(gin.Logger(), gin.Recovery())
    r.Use(middleware.SecurityHeaders())
    r.Use(middleware.RateLimit())
    r.Use(middleware.JWTAuth())

    r.POST("/transactions", h.Create)

    r.Run(":8080")
}
