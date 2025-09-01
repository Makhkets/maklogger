package maklogger

import (
	"path/filepath"
	"runtime"
)

// getCallerInfo retrieves the file name, line number, and function name
// of the caller at the specified skip level in the call stack.
// This is used internally to provide source location information in logs.
func getCallerInfo(skip int) (file string, line int, function string) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "???", 0, "???"
	}
	fn := runtime.FuncForPC(pc)
	funcName := "???"
	if fn != nil {
		funcName = fn.Name()
	}
	return filepath.Base(file), line, funcName
}
