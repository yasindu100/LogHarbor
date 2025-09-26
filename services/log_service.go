package services

import (
    "logharbor/models"
    "logharbor/storage"
)

func SaveLog(log models.LogEntry) {
    storage.AddLog(log)
}

func GetAllLogs() []models.LogEntry {
    return storage.GetLogs()
}
