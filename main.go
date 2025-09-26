package main

import (
    "logharbor/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":5000")
}
