package main

import (
	"github.com/makhkets/maklogger"
)

func main() {
	// Create a new logger instance
	logger := maklogger.NewLogger()

	// Disable colors for plain text output
	logger.SetColorsEnabled(false)

	// Log messages without colors - useful for log files or systems that don't support colors
	logger.Info("Application started (no colors)")
	logger.Success("Operation completed successfully (no colors)")
	logger.Debug("Debug information (no colors)")
	logger.Warn("Warning message (no colors)")
	logger.Error("Error occurred (no colors)")
	logger.Critical("Critical error (no colors)")

	// With structured fields (also without colors)
	logger.Info("User action logged (no colors)",
		maklogger.Field{Key: "user_id", Value: 456},
		maklogger.Field{Key: "action", Value: "login"},
		maklogger.Field{Key: "timestamp", Value: "2025-09-02T15:30:45Z"},
	)

	// Check if colors are enabled
	if !logger.ColorsEnabled() {
		logger.Info("Colors are disabled - perfect for log files!")
	}
}
