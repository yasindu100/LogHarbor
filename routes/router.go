package routes

import (
    "logharbor/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/logs", controllers.PostLog)
    r.GET("/logs", controllers.GetLogs)
}
