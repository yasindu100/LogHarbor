package storage

import (
    "logharbor/models"
    "sync"
)

var (
    logs []models.LogEntry
    mu   sync.Mutex
    id   = 1
)

func AddLog(log models.LogEntry) {
    mu.Lock()
    log.ID = id
    id++
    logs = append(logs, log)
    mu.Unlock()
}

func GetLogs() []models.LogEntry {
    mu.Lock()
    defer mu.Unlock()
    return logs
}
