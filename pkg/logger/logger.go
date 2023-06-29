package logger

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger = zap.SugaredLogger

func New(logPath string, logLevel string) (*Logger, func(), error) {
	if err := os.MkdirAll(logPath, 0755); err != nil {
		return nil, nil, err
	}

	now := time.Now()
	logfile := path.Join(logPath, fmt.Sprintf("%s.log", now.Format("2006-01-02")))

	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, err
	}

	var level zapcore.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		level = zapcore.InfoLevel
	}

	// create config encoder
	config := zap.NewProductionEncoderConfig()

	// modify encode config for both file/console encoder
	config.TimeKey = "timestamp"
	config.LevelKey = "severity"
	config.MessageKey = "message"
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	// modify encode config for file encoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	// new encode for file
	fileEncoder := zapcore.NewJSONEncoder(config)

	// modify encode config for console encoder
	config.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + level.CapitalString() + "]")
	}

	// new encode for console
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	// Set PID once
	pid := strconv.Itoa(os.Getpid())

	// create the logger
	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.WarnLevel),
		zap.Fields(zap.Field{Key: "pid", Type: zapcore.StringType, String: pid}),
	)

	sugar := logger.Sugar()

	close := func() {
		logger.Sync()
		file.Close()
	}

	return sugar, close, nil
}
