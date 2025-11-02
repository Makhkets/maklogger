# MakLogger 🚀

A lightweight, colorful, and beautiful logging library for Go applications with emoji support and structured field logging.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-blue.svg)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen)

## ✨ Features

- 🎨 **Beautiful colored output** with emoji icons
- 📊 **Structured field logging** with JSON formatting
- 🖥️ **Cross-platform support** (Windows, Linux)
- ⚡ **High performance** with minimal overhead
- 🔧 **Simple API** - easy to integrate
- 🎯 **Multiple log levels** (Info, Success, Debug, Warning, Error, Critical)
- 🌈 **Automatic ANSI color detection** and Windows CMD support

## 📦 Installation

```bash
go get github.com/makhkets/maklogger
```

## 🚀 Quick Start

```go
package main

import (
    "errors"
    "github.com/makhkets/maklogger"
)

func main() {
    // Create a new logger instance
    logger := maklogger.NewLogger()
    
    // Simple logging
    logger.Info("Application started successfully")
    logger.Success("Database connection established")
    logger.Warn("High memory usage detected")
    logger.Error("Failed to process request")
    logger.Debug("Processing user data")
    logger.Critical("System overload detected")
}
```

## 📋 Structured Logging with Fields

Add structured data to your logs using fields:

```go
logger.Info("User logged in",
    maklogger.Field{Key: "user_id", Value: 12345},
    maklogger.Field{Key: "username", Value: "john_doe"},
    maklogger.Field{Key: "login_time", Value: "2025-09-02T10:30:45Z"},
)

logger.Error("Database query failed",
    maklogger.Field{Key: "error", Value: errors.New("connection timeout")},
    maklogger.Field{Key: "query", Value: "SELECT * FROM users"},
    maklogger.Field{Key: "duration_ms", Value: 5000},
)
```

## 🎨 Log Levels and Colors

| Level | Icon | Color | Description |
|-------|------|-------|-------------|
| `Info` | 📝 | Blue | General information |
| `Success` | ✅ | Green | Success operations |
| `Debug` | 🐛 | Magenta | Debug information |
| `Warning` | ⚠️ | Yellow | Warning messages |
| `Error` | ❌ | Red | Error messages |
| `Critical` | 🛑 | Bright Red | Critical errors |

## ⚙️ Configuration

### Disable Colors

```go
logger := maklogger.NewLogger()
logger.SetColorsEnabled(false) // Disable colors

// Check if colors are enabled
if logger.ColorsEnabled() {
    fmt.Println("Colors are enabled")
}
```

## 📁 Output Format

The logger produces beautiful, structured output:

```
🕒 2025-09-02 15:30:45.123 │ 📝 INFO     │ 📁 main.go:15 ⚡ main │ 💬 Application started successfully
🕒 2025-09-02 15:30:45.124 │ ✅ SUCCESS  │ 📁 main.go:16 ⚡ main │ 💬 Database connection established
🕒 2025-09-02 15:30:45.125 │ 📝 INFO     │ 📁 main.go:20 ⚡ main │ 💬 User logged in
📊 Fields:
  {
    "login_time": "2025-09-02T10:30:45Z",
    "user_id": 12345,
    "username": "john_doe"
  }
```

## 🏗️ API Reference

### Types

```go
// MakLogger is the main logger instance
type MakLogger struct {
    // Contains unexported fields
}

// Field represents a key-value pair for structured logging
type Field struct {
    Key   string
    Value any
}
```

### Methods

```go
// Create a new logger instance
func NewLogger() *MakLogger

// Log level methods
func (mk *MakLogger) Info(msg string, fields ...Field)
func (mk *MakLogger) Success(msg string, fields ...Field) 
func (mk *MakLogger) Debug(msg string, fields ...Field)
func (mk *MakLogger) Warn(msg string, fields ...Field)
func (mk *MakLogger) Error(msg string, fields ...Field)
func (mk *MakLogger) Critical(msg string, fields ...Field)

// Configuration methods
func (mk *MakLogger) ColorsEnabled() bool
func (mk *MakLogger) SetColorsEnabled(enabled bool)
```

## 🖥️ Platform Support

- ✅ **Windows** - Automatic ANSI color support for CMD and PowerShell
- ✅ **Linux** - Native ANSI color support

## 🔧 Requirements

- Go 1.21 or higher
- No external dependencies (uses only Go standard library)

## 📚 Examples

Check out the `examples/` directory for more usage examples:

- [Basic Usage](examples/basic/main.go) - Simple logging examples
- [Advanced Usage](examples/advanced/main.go) - Complex logging with fields
- [Without Colors](examples/no-colors/main.go) - Logging without colors

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by modern logging libraries
- Built with ❤️ for the Go community
- Special thanks to all contributors

## 📞 Support

If you have any questions or need help, please:

- Open an issue on GitHub
- Check the documentation
- Look at the examples

---

Made with ❤️ by [makhkets](https://github.com/makhkets)
