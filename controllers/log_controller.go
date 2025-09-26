package controllers

import (
    "logharbor/models"
    "logharbor/services"
    "net/http"
    "github.com/gin-gonic/gin"
)

func PostLog(c *gin.Context) {
    var log models.LogEntry
    if err := c.ShouldBindJSON(&log); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    services.SaveLog(log)
    c.JSON(http.StatusCreated, gin.H{"status": "log saved"})
}

func GetLogs(c *gin.Context) {
    logs := services.GetAllLogs()
    c.JSON(http.StatusOK, logs)
}
