// Package maklogger provides a lightweight, colored logging library for Go applications.
// It supports multiple log levels with beautiful colored output and structured field logging.
package maklogger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

// MakLogger represents the main logger instance with configurable color support.
type MakLogger struct {
	colorsEnabled bool
}

// Field represents a key-value pair for structured logging.
// Fields are displayed as formatted JSON in the log output.
type Field struct {
	Key   string
	Value any
}

var buf bytes.Buffer

// NewLogger creates a new MakLogger instance with colors enabled by default.
// On Windows, it automatically enables ANSI color support for CMD.
// On Unix systems (Linux/macOS), ANSI colors are supported by default.
func NewLogger() *MakLogger {
	logger := &MakLogger{colorsEnabled: true}

	// Enable ANSI colors for Windows CMD
	if runtime.GOOS == "windows" {
		logger.enableWindowsANSI()
	}
	// On Unix systems (Linux/macOS) ANSI colors are supported by default

	return logger
}

// enableWindowsANSI enables ANSI escape sequence support in Windows CMD.
func (mk *MakLogger) enableWindowsANSI() {
	if runtime.GOOS != "windows" {
		return // Do nothing on non-Windows systems
	}

	defer func() {
		if r := recover(); r != nil {
			// If we couldn't enable ANSI, disable colors
			mk.colorsEnabled = false
		}
	}()

	// Windows-specific constants
	const (
		STD_OUTPUT_HANDLE                  = ^uintptr(10) // -11 as uintptr
		ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004
	)

	// Load Windows API functions
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode := kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode := kernel32.NewProc("SetConsoleMode")
	procGetStdHandle := kernel32.NewProc("GetStdHandle")

	handle, _, _ := procGetStdHandle.Call(STD_OUTPUT_HANDLE)
	var mode uint32

	// Get current console mode
	ret, _, _ := procGetConsoleMode.Call(handle, uintptr(unsafe.Pointer(&mode)))
	if ret == 0 {
		mk.colorsEnabled = false
		return
	}

	// Enable virtual terminal
	mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING
	ret, _, _ = procSetConsoleMode.Call(handle, uintptr(mode))
	if ret == 0 {
		mk.colorsEnabled = false
	}
}

// ColorsEnabled returns whether colors are currently enabled.
func (mk *MakLogger) ColorsEnabled() bool {
	return mk.colorsEnabled
}

// SetColorsEnabled sets whether colors should be used in log output.
func (mk *MakLogger) SetColorsEnabled(enabled bool) {
	mk.colorsEnabled = enabled
}

// log is the core logging method that formats and outputs log messages.
func (mk *MakLogger) log(level Level, color Color, msg string, fields ...Field) {
	file, line, fn := getCallerInfo(3)

	// Get detailed information
	now := time.Now()
	timestamp := now.Format("2006-01-02 15:04:05.000")

	// Format module and function
	moduleParts := strings.Split(fn, ".")
	shortFn := fn
	if len(moduleParts) > 0 {
		shortFn = moduleParts[len(moduleParts)-1]
	}

	// Create beautiful module with icons
	module := fmt.Sprintf("%s %s:%s %s %s",
		ColorizeIfEnabled("📁", mk.colorsEnabled, BrightBlue),
		ColorizeIfEnabled(file, mk.colorsEnabled, Cyan),
		ColorizeIfEnabled(strconv.Itoa(line), mk.colorsEnabled, BrightCyan),
		ColorizeIfEnabled("⚡", mk.colorsEnabled, BrightYellow),
		ColorizeIfEnabled(shortFn, mk.colorsEnabled, Magenta),
	)

	// Main message without PID (according to specification)
	message := fmt.Sprintf("%s %s │ %s │ %s │ %s %s",
		ColorizeIfEnabled("🕒 ", mk.colorsEnabled, BrightGreen),
		ColorizeIfEnabled(timestamp, mk.colorsEnabled, Green),
		mk.getColoredLevel(level),
		module,
		ColorizeIfEnabled("💬 ", mk.colorsEnabled, BrightWhite),
		mk.getColoredMessage(level, msg),
	)

	fmt.Println(message)

	// Process fields if they exist - display on next line (according to specification)
	if len(fields) > 0 {
		fieldStr := mk.formatFieldsAsJSON(fields)
		fmt.Printf("%s %s\n%s\n",
			ColorizeIfEnabled("📊 ", mk.colorsEnabled, BrightMagenta),
			ColorizeIfEnabled("Fields:", mk.colorsEnabled, BrightWhite),
			ColorizeIfEnabled(fieldStr, mk.colorsEnabled, BrightBlack), // gray color for JSON
		)
	}
}

// Info logs an informational message with optional structured fields.
func (mk *MakLogger) Info(msg string, fields ...Field) {
	mk.log(LevelInfo, Yellow, msg, fields...)
}

// Warn logs a warning message with optional structured fields.
func (mk *MakLogger) Warn(msg string, fields ...Field) {
	mk.log(LevelWarn, Yellow, msg, fields...)
}

// Error logs an error message with optional structured fields.
func (mk *MakLogger) Error(msg string, fields ...Field) {
	mk.log(LevelError, Red, msg, fields...)
}

// Success logs a success message with optional structured fields.
func (mk *MakLogger) Success(msg string, fields ...Field) {
	mk.log(LevelSuccess, Red, msg, fields...)
}

// Debug logs a debug message with optional structured fields.
func (mk *MakLogger) Debug(msg string, fields ...Field) {
	mk.log(LevelDebug, Red, msg, fields...)
}

// Critical logs a critical message with optional structured fields.
func (mk *MakLogger) Critical(msg string, fields ...Field) {
	mk.log(LevelCritical, Red, msg, fields...)
}

// formatFieldsAsJSON formats fields into a beautiful JSON string (according to specification with 2-space indentation).
func (mk *MakLogger) formatFieldsAsJSON(fields []Field) string {
	if len(fields) == 0 {
		return ""
	}

	// Create map for JSON serialization
	fieldMap := make(map[string]interface{})
	for _, field := range fields {
		fieldMap[field.Key] = field.Value
	}

	// Serialize to beautiful JSON with indentation (json.MarshalIndent with 2-space indentation)
	jsonBytes, err := json.MarshalIndent(fieldMap, "  ", "  ")
	if err != nil {
		return fmt.Sprintf(`  {
    "error": "failed to marshal fields: %v"
  }`, err)
	}

	// Add indentation to each JSON line for beautiful output
	lines := strings.Split(string(jsonBytes), "\n")
	for i, line := range lines {
		lines[i] = "  " + line
	}

	return strings.Join(lines, "\n")
}

// getColoredLevel returns a formatted log level with color settings.
func (mk *MakLogger) getColoredLevel(level Level) string {
	switch level {
	case LevelInfo:
		return fmt.Sprintf("%s %s",
			ColorizeIfEnabled("📝 ", mk.colorsEnabled, BrightBlue),
			ColorizeIfEnabled("INFO    ", mk.colorsEnabled, BoldWhite, BgBlue))
	case LevelSuccess:
		return fmt.Sprintf("%s %s",
			ColorizeIfEnabled("✅ ", mk.colorsEnabled, BrightGreen),
			ColorizeIfEnabled("SUCCESS ", mk.colorsEnabled, BoldWhite, BgGreen))
	case LevelDebug:
		return fmt.Sprintf("%s %s",
			ColorizeIfEnabled("🐛 ", mk.colorsEnabled, BrightMagenta),
			ColorizeIfEnabled("DEBUG   ", mk.colorsEnabled, BoldWhite, BgMagenta))
	case LevelCritical:
		return fmt.Sprintf("%s %s",
			ColorizeIfEnabled("🛑 ", mk.colorsEnabled, BrightRed),
			ColorizeIfEnabled("CRITICAL", mk.colorsEnabled, BoldWhite, BgBrightRed))
	case LevelError:
		return fmt.Sprintf("%s %s",
			ColorizeIfEnabled("❌ ", mk.colorsEnabled, BrightRed),
			ColorizeIfEnabled("ERROR   ", mk.colorsEnabled, BoldWhite, BgRed))
	case LevelWarn:
		return fmt.Sprintf("%s %s",
			ColorizeIfEnabled("⚠️ ", mk.colorsEnabled, BrightYellow),
			ColorizeIfEnabled("WARNING ", mk.colorsEnabled, Bold, BgYellow))
	}

	return "UNDEFINED"
}

// getColoredMessage returns a formatted message with color settings.
func (mk *MakLogger) getColoredMessage(level Level, message string) string {
	switch level {
	case LevelInfo:
		return ColorizeIfEnabled(message, mk.colorsEnabled, BrightWhite)
	case LevelSuccess:
		return ColorizeIfEnabled(message, mk.colorsEnabled, BrightGreen)
	case LevelDebug:
		return ColorizeIfEnabled(message, mk.colorsEnabled, BrightMagenta)
	case LevelCritical:
		return ColorizeIfEnabled(message, mk.colorsEnabled, BrightRed, BgBlack)
	case LevelError:
		return ColorizeIfEnabled(message, mk.colorsEnabled, BrightRed)
	case LevelWarn:
		return ColorizeIfEnabled(message, mk.colorsEnabled, BrightYellow)
	}

	return "UNDEFINED"
}
