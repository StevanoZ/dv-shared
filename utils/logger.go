package shrd_utils

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"

	timeZone := "Asia/Jakarta"
	defaultTimeZone := os.Getenv("DEFAULT_TIME_ZONE")

	if defaultTimeZone != "" {
		timeZone = defaultTimeZone
	}

	loc, _ := time.LoadLocation(timeZone)
	_, hours := time.Now().In(loc).Zone()
	differentHours := hours / 3600

	encoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Add(time.Duration(differentHours) * time.Hour).UTC().Format("2006-01-02T15:04:05Z0700"))
	})
	encoderConfig.StacktraceKey = os.Getenv("SERVICE_NAME")
	config.EncoderConfig = encoderConfig

	logger, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		logger.Fatal("Error when init logger: " + err.Error())
	}
}

func LogInfo(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func LogFatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}

func LogPanic(message string, fields ...zap.Field) {
	logger.Panic(message, fields...)
}

func LogDebug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func LogError(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}
