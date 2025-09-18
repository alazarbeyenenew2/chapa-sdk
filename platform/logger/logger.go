package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Debug(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
	Sync() error
}

type ZapLogger struct {
	logger *zap.Logger
}

func NewLogger(isProduction bool) (*ZapLogger, error) {
	var logger *zap.Logger
	var err error

	if isProduction {

		config := zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger, err = config.Build()
	} else {

		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger, err = config.Build()
	}

	if err != nil {
		return nil, err
	}

	return &ZapLogger{logger: logger}, nil
}

func (l *ZapLogger) Info(msg string, fields ...interface{}) {
	zapFields := convertFields(fields...)
	l.logger.Info(msg, zapFields...)
}

func (l *ZapLogger) Error(msg string, fields ...interface{}) {
	zapFields := convertFields(fields...)
	l.logger.Error(msg, zapFields...)
}

func (l *ZapLogger) Debug(msg string, fields ...interface{}) {
	zapFields := convertFields(fields...)
	l.logger.Debug(msg, zapFields...)
}

func (l *ZapLogger) Warn(msg string, fields ...interface{}) {
	zapFields := convertFields(fields...)
	l.logger.Warn(msg, zapFields...)
}

func (l *ZapLogger) Fatal(msg string, fields ...interface{}) {
	zapFields := convertFields(fields...)
	l.logger.Fatal(msg, zapFields...)
}

func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

func convertFields(fields ...interface{}) []zapcore.Field {
	zapFields := make([]zapcore.Field, 0, len(fields)/2)
	for i := 0; i < len(fields)-1; i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		zapFields = append(zapFields, zap.Any(key, fields[i+1]))
	}
	return zapFields
}
