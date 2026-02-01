
package middleware

import (
    "time"

    "github.com/gin-gonic/gin"
)

var bucket = make(chan struct{}, 50)

func init() {
    go func() {
        ticker := time.NewTicker(time.Second)
        for range ticker.C {
            for len(bucket) < cap(bucket) {
                bucket <- struct{}{}
            }
        }
    }()
}

func RateLimit() gin.HandlerFunc {
    return func(c *gin.Context) {
        select {
        case <-bucket:
            c.Next()
        default:
            c.AbortWithStatus(429)
        }
    }
}
