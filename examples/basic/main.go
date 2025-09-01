package main

import (
	"github.com/makhkets/maklogger"
)

func main() {
	// Create a new logger instance
	logger := maklogger.NewLogger()

	// Simple logging examples
	logger.Info("Application started successfully")
	logger.Success("Database connection established")
	logger.Debug("Processing user request")
	logger.Warn("High memory usage detected")
	logger.Error("Failed to load configuration")
	logger.Critical("System overload detected")
}
