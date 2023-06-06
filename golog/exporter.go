package golog

// all exported log function

// Info log with info level
func Info(msg any) {
}

// Infow log with info level, with key-value pair as option
func Infow(msg string, kv ...any) {
}

// Warn log with warn level
func Warn(msg any) {
}

// Warnw log with warn level, with key-value pair as option
func Warnw(msg string, kv ...any) {
}

// Error log with error level
func Error(err any) {
}

// Errorw log with error level, with key-value pair as option
func Errorw(msg string, kv ...interface{}) {
}

// Raw log without level
func Raw(body []byte) {
}
