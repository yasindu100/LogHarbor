package models

import "time"

type LogEntry struct {
    ID        int       `json:"id"`
    Level     string    `json:"level"`
    Message   string    `json:"message"`
    Source    string    `json:"source"`
    Timestamp time.Time `json:"timestamp"`
}
