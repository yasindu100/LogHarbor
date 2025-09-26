package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
)

func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Printf("Incoming %s request to %s", c.Request.Method, c.Request.URL.Path)
        c.Next()
    }
}
