package maklogger

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// captureOutput captures stdout for testing log output
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	if logger == nil {
		t.Fatal("NewLogger() returned nil")
	}

	// By default, colors should be enabled
	if !logger.ColorsEnabled() {
		t.Error("Colors should be enabled by default")
	}
}

func TestSetColorsEnabled(t *testing.T) {
	logger := NewLogger()

	// Test disabling colors
	logger.SetColorsEnabled(false)
	if logger.ColorsEnabled() {
		t.Error("Colors should be disabled")
	}

	// Test enabling colors
	logger.SetColorsEnabled(true)
	if !logger.ColorsEnabled() {
		t.Error("Colors should be enabled")
	}
}

func TestLogLevels(t *testing.T) {
	logger := NewLogger()
	logger.SetColorsEnabled(false) // Disable colors for easier testing

	tests := []struct {
		name     string
		logFunc  func(string, ...Field)
		message  string
		expected string
	}{
		{"Info", logger.Info, "test info message", "INFO"},
		{"Success", logger.Success, "test success message", "SUCCESS"},
		{"Debug", logger.Debug, "test debug message", "DEBUG"},
		{"Warn", logger.Warn, "test warn message", "WARNING"},
		{"Error", logger.Error, "test error message", "ERROR"},
		{"Critical", logger.Critical, "test critical message", "CRITICAL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				tt.logFunc(tt.message)
			})

			if !strings.Contains(output, tt.expected) {
				t.Errorf("Expected output to contain '%s', got: %s", tt.expected, output)
			}

			if !strings.Contains(output, tt.message) {
				t.Errorf("Expected output to contain message '%s', got: %s", tt.message, output)
			}
		})
	}
}

func TestLogWithFields(t *testing.T) {
	logger := NewLogger()
	logger.SetColorsEnabled(false) // Disable colors for easier testing

	fields := []Field{
		{Key: "user_id", Value: 123},
		{Key: "username", Value: "testuser"},
		{Key: "active", Value: true},
	}

	output := captureOutput(func() {
		logger.Info("test message with fields", fields...)
	})

	// Check that the main message is present
	if !strings.Contains(output, "test message with fields") {
		t.Error("Expected output to contain the main message")
	}

	// Check that Fields: header is present
	if !strings.Contains(output, "Fields:") {
		t.Error("Expected output to contain 'Fields:' header")
	}

	// Check that field values are present in JSON format
	if !strings.Contains(output, "user_id") {
		t.Error("Expected output to contain 'user_id' field")
	}

	if !strings.Contains(output, "testuser") {
		t.Error("Expected output to contain 'testuser' value")
	}
}

func TestFieldTypes(t *testing.T) {
	logger := NewLogger()
	logger.SetColorsEnabled(false)

	// Test various field types
	fields := []Field{
		{Key: "string_field", Value: "test string"},
		{Key: "int_field", Value: 42},
		{Key: "float_field", Value: 3.14},
		{Key: "bool_field", Value: true},
		{Key: "nil_field", Value: nil},
	}

	output := captureOutput(func() {
		logger.Info("testing field types", fields...)
	})

	// Verify different types are handled correctly
	expectedValues := []string{"test string", "42", "3.14", "true", "null"}
	for _, expected := range expectedValues {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain '%s', got: %s", expected, output)
		}
	}
}

func TestColorize(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		fg       Color
		bg       []Color
		expected string
	}{
		{
			name:     "foreground only",
			text:     "test",
			fg:       Red,
			expected: "\033[31mtest\033[0m",
		},
		{
			name:     "foreground and background",
			text:     "test",
			fg:       Red,
			bg:       []Color{BgBlue},
			expected: "\033[31m\033[44mtest\033[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Colorize(tt.text, tt.fg, tt.bg...)
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestColorizeIfEnabled(t *testing.T) {
	text := "test"
	colored := ColorizeIfEnabled(text, true, Red)
	uncolored := ColorizeIfEnabled(text, false, Red)

	// When enabled, should return colored text
	if colored == text {
		t.Error("Expected colored text when colors are enabled")
	}

	// When disabled, should return plain text
	if uncolored != text {
		t.Error("Expected plain text when colors are disabled")
	}
}

func TestFormatFieldsAsJSON(t *testing.T) {
	logger := NewLogger()

	// Test empty fields
	result := logger.formatFieldsAsJSON([]Field{})
	if result != "" {
		t.Error("Expected empty string for no fields")
	}

	// Test single field
	fields := []Field{
		{Key: "test_key", Value: "test_value"},
	}
	result = logger.formatFieldsAsJSON(fields)

	if !strings.Contains(result, "test_key") {
		t.Error("Expected result to contain field key")
	}

	if !strings.Contains(result, "test_value") {
		t.Error("Expected result to contain field value")
	}

	// Check JSON formatting with proper indentation
	if !strings.Contains(result, "  {") {
		t.Error("Expected proper JSON indentation")
	}
}

func TestGetCallerInfo(t *testing.T) {
	file, line, function := getCallerInfo(0)

	// Should not return default values for valid caller
	if file == "???" || line == 0 || function == "???" {
		t.Error("getCallerInfo should return valid caller information")
	}

	// File should be the test file
	if !strings.Contains(file, "test.go") {
		t.Errorf("Expected file to contain 'test.go', got: %s", file)
	}

	// Function should contain test function name
	if !strings.Contains(function, "Test") {
		t.Errorf("Expected function to contain 'Test', got: %s", function)
	}

	// Line should be positive
	if line <= 0 {
		t.Errorf("Expected positive line number, got: %d", line)
	}
}

func TestLogTimestamp(t *testing.T) {
	logger := NewLogger()
	logger.SetColorsEnabled(false)

	output := captureOutput(func() {
		logger.Info("timestamp test")
	})

	// Check that output contains timestamp-like format (YYYY-MM-DD HH:MM:SS.mmm)
	if !strings.Contains(output, "2025-") && !strings.Contains(output, "2024-") {
		t.Error("Expected output to contain year")
	}

	// Check for time separator
	if !strings.Contains(output, ":") {
		t.Error("Expected output to contain time separators")
	}
}

func TestComplexFieldValues(t *testing.T) {
	logger := NewLogger()
	logger.SetColorsEnabled(false)

	// Test complex nested structures
	complexValue := map[string]interface{}{
		"nested": map[string]interface{}{
			"key":    "value",
			"number": 42,
		},
		"array": []string{"item1", "item2"},
	}

	fields := []Field{
		{Key: "complex", Value: complexValue},
	}

	output := captureOutput(func() {
		logger.Info("complex field test", fields...)
	})

	// Should contain nested structure elements
	if !strings.Contains(output, "nested") {
		t.Error("Expected output to contain nested structure")
	}

	if !strings.Contains(output, "array") {
		t.Error("Expected output to contain array field")
	}
}

// Benchmark tests
func BenchmarkLogger_Info(b *testing.B) {
	logger := NewLogger()
	logger.SetColorsEnabled(false) // Disable colors for consistent benchmarking

	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("benchmark test message")
	}
}

func BenchmarkLogger_InfoWithFields(b *testing.B) {
	logger := NewLogger()
	logger.SetColorsEnabled(false)

	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()

	fields := []Field{
		{Key: "user_id", Value: 123},
		{Key: "action", Value: "test"},
		{Key: "timestamp", Value: "2025-09-02T15:30:45Z"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("benchmark test with fields", fields...)
	}
}
