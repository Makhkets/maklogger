package main

import (
	"errors"
	"time"

	"github.com/makhkets/maklogger"
)

func main() {
	// Create a new logger instance
	logger := maklogger.NewLogger()

	// Advanced logging with structured fields
	logger.Info("User authentication started",
		maklogger.Field{Key: "user_id", Value: 12345},
		maklogger.Field{Key: "username", Value: "john_doe"},
		maklogger.Field{Key: "ip_address", Value: "192.168.1.100"},
		maklogger.Field{Key: "user_agent", Value: "Mozilla/5.0 (Windows NT 10.0; Win64; x64)"},
	)

	logger.Success("Database transaction completed",
		maklogger.Field{Key: "transaction_id", Value: "tx_789123"},
		maklogger.Field{Key: "affected_rows", Value: 5},
		maklogger.Field{Key: "duration_ms", Value: 234},
		maklogger.Field{Key: "table", Value: "users"},
	)

	logger.Warn("Performance warning",
		maklogger.Field{Key: "memory_usage", Value: "85%"},
		maklogger.Field{Key: "cpu_usage", Value: "78%"},
		maklogger.Field{Key: "response_time_ms", Value: 1500},
		maklogger.Field{Key: "threshold_ms", Value: 1000},
	)

	logger.Error("API request failed",
		maklogger.Field{Key: "error", Value: errors.New("connection timeout")},
		maklogger.Field{Key: "endpoint", Value: "/api/v1/users"},
		maklogger.Field{Key: "method", Value: "GET"},
		maklogger.Field{Key: "status_code", Value: 503},
		maklogger.Field{Key: "retry_count", Value: 3},
		maklogger.Field{Key: "timestamp", Value: time.Now().UTC()},
	)

	// Complex nested data
	userData := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"preferences": map[string]interface{}{
			"theme":    "dark",
			"language": "en",
			"notifications": map[string]bool{
				"email": true,
				"push":  false,
			},
		},
	}

	logger.Debug("User data processed",
		maklogger.Field{Key: "user_data", Value: userData},
		maklogger.Field{Key: "processing_time_ms", Value: 45},
		maklogger.Field{Key: "cache_hit", Value: true},
	)

	logger.Critical("System critical error",
		maklogger.Field{Key: "error_code", Value: "SYS_001"},
		maklogger.Field{Key: "subsystem", Value: "database"},
		maklogger.Field{Key: "available_connections", Value: 0},
		maklogger.Field{Key: "max_connections", Value: 100},
		maklogger.Field{Key: "uptime_seconds", Value: 86400},
	)
}
