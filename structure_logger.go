package superlog

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// structure log with zap
type zapLogger struct {
	*zap.SugaredLogger
}

// sugared logger instance
func newZapStructureLogger(writer io.Writer, zapEncoder zapcore.Encoder) *zapLogger {
	core := zapcore.NewCore(
		zapEncoder,
		zapcore.AddSync(writer),
		zapcore.InfoLevel, // level to enable
	)

	// >= WARN level => enable stacktrace
	zap := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))

	return &zapLogger{
		SugaredLogger: zap.Sugar(),
	}
}

type zapEncoderOption func(*zapcore.EncoderConfig)

func newZapEncoderConfig(opts ...zapEncoderOption) zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		// TimeKey:    "ts",
		// NameKey:    "logger",
		// CallerKey:  "caller",
		// FunctionKey:    "",
		// StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     "\n",
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(l.CapitalString())
		},
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: func(time.Duration, zapcore.PrimitiveArrayEncoder) {
		},
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(caller.TrimmedPath())
		},
		EncodeName: func(name string, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(name)
		},
		// NewReflectedEncoder: func(io.Writer) zapcore.ReflectedEncoder {
		// },
		ConsoleSeparator: "",
	}

	// customize encoder config
	for _, opt := range opts {
		opt(&encoderConfig)
	}

	return encoderConfig
}

func zapEncoderWithTimeKey(timeKey string) zapEncoderOption {
	return func(encoderConfig *zapcore.EncoderConfig) {
		encoderConfig.TimeKey = timeKey
	}
}

func zapEncoderWithStacktraceKey(stacktraceKey string) zapEncoderOption {
	return func(encoderConfig *zapcore.EncoderConfig) {
		encoderConfig.StacktraceKey = stacktraceKey
	}
}

func (zap *zapLogger) info(msg any) {
	zap.Info(msg)
}

func (zap *zapLogger) infow(msg string, kv ...interface{}) {
	zap.Infow(msg, kv...)
}

func (zap *zapLogger) warn(msg any) {
	zap.Warn(msg)
}

func (zap *zapLogger) warnw(msg string, kv ...interface{}) {
	zap.Warnw(msg, kv...)
}
func (zap *zapLogger) error(msg any) {
	zap.Error(msg)
}
func (zap *zapLogger) errorf(msg string, args ...any) {
	zap.Errorf(msg, args...)
}

func (zap *zapLogger) errorw(msg string, kv ...interface{}) {
	zap.Errorw(msg, kv...)
}
