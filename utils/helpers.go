package utils

import "time"

func ParseTimestamp(ts string) (time.Time, error) {
    return time.Parse(time.RFC3339, ts)
}
