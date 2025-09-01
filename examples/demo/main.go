package main

import (
	"errors"

	"github.com/makhkets/maklogger"
)

func main() {
	l := maklogger.NewLogger()

	// Примеры с различными типами сообщений
	l.Info("Application started successfully")
	l.Success("Database connection established")
	l.Debug("Processing user request")

	// Примеры с полями (будут отображаться как JSON)
	l.Info("User logged in",
		maklogger.Field{Key: "user_id", Value: 12345},
		maklogger.Field{Key: "username", Value: "john_doe"},
		maklogger.Field{Key: "login_time", Value: "2025-09-02T10:30:45Z"},
	)

	l.Warn("High memory usage detected",
		maklogger.Field{Key: "memory_usage", Value: "85%"},
		maklogger.Field{Key: "threshold", Value: "80%"},
		maklogger.Field{Key: "action", Value: "alert_sent"},
	)

	l.Error("Database query failed",
		maklogger.Field{Key: "error", Value: errors.New("connection timeout")},
		maklogger.Field{Key: "query", Value: "SELECT * FROM users WHERE active = true"},
		maklogger.Field{Key: "duration_ms", Value: 5000},
	)

	l.Critical("System overload detected",
		maklogger.Field{Key: "cpu_usage", Value: "98%"},
		maklogger.Field{Key: "load_average", Value: 15.7},
		maklogger.Field{Key: "available_memory", Value: "2GB"},
	)
}
