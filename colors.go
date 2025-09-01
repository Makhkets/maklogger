package maklogger

import "fmt"

// Color represents an ANSI color code.
type Color string

// Level represents the severity level of a log message.
type Level int

// Log levels in order of severity.
const (
	LevelInfo Level = iota
	LevelSuccess
	LevelDebug
	LevelCritical
	LevelError
	LevelWarn
)

// ANSI color codes for text formatting.
const (
	Reset         Color = "\033[0m"
	Bold          Color = "\033[1m"
	Dim           Color = "\033[2m"
	Italic        Color = "\033[3m"
	Underline     Color = "\033[4m"
	Blink         Color = "\033[5m"
	Inverse       Color = "\033[7m"
	Hidden        Color = "\033[8m"
	Strikethrough Color = "\033[9m"

	BoldWhite Color = "\033[1;97m"
	Gray      Color = "\033[90m"   // Gray color for JSON fields
	DarkGray  Color = "\033[2;37m" // Dark gray

	// Text colors
	Black   Color = "\033[30m"
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Cyan    Color = "\033[36m"
	White   Color = "\033[37m"

	// Bright text colors
	BrightBlack   Color = "\033[90m"
	BrightRed     Color = "\033[91m"
	BrightGreen   Color = "\033[92m"
	BrightYellow  Color = "\033[93m"
	BrightBlue    Color = "\033[94m"
	BrightMagenta Color = "\033[95m"
	BrightCyan    Color = "\033[96m"
	BrightWhite   Color = "\033[97m"

	// Background colors
	BgBlack   Color = "\033[40m"
	BgRed     Color = "\033[41m"
	BgGreen   Color = "\033[42m"
	BgYellow  Color = "\033[43m"
	BgBlue    Color = "\033[44m"
	BgMagenta Color = "\033[45m"
	BgCyan    Color = "\033[46m"
	BgWhite   Color = "\033[47m"

	// Bright background colors
	BgBrightBlack   Color = "\033[100m"
	BgBrightRed     Color = "\033[101m"
	BgBrightGreen   Color = "\033[102m"
	BgBrightYellow  Color = "\033[103m"
	BgBrightBlue    Color = "\033[104m"
	BgBrightMagenta Color = "\033[105m"
	BgBrightCyan    Color = "\033[106m"
	BgBrightWhite   Color = "\033[107m"
)

// Colorize applies ANSI color codes to text with optional background color.
func Colorize(text string, fg Color, bg ...Color) string {
	if len(bg) > 0 {
		return fmt.Sprintf("%s%s%s%s", fg, bg[0], text, Reset)
	}
	return fmt.Sprintf("%s%s%s", fg, text, Reset)
}

// ColorizeIfEnabled applies colors only if they are enabled.
// This function is used internally to respect the color settings.
func ColorizeIfEnabled(text string, enabled bool, fg Color, bg ...Color) string {
	if !enabled {
		return text
	}
	return Colorize(text, fg, bg...)
}
