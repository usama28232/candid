package logging

import (
	"authexample/shared"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogProfile string

const (
	info  LogProfile = "info"
	debug LogProfile = "debug"
)

var (
	appLogger    *zap.SugaredLogger = nil
	accessLogger *zap.SugaredLogger = nil
)

func GetAccessLogger() *zap.SugaredLogger {
	if accessLogger == nil {
		accessLogger = getLogger(zap.InfoLevel, shared.ACCESS_LOG_FILE)
	}
	return accessLogger
}

func GetLogger() *zap.SugaredLogger {
	if appLogger == nil {
		prof, err := shared.GetConfigByKey(shared.LOG_PROFILE_KEY)
		if err == nil {
			logProfileStr := strings.ToLower(prof.(string))
			switch logProfileStr {
			case string(debug):
				appLogger = getLogger(zap.DebugLevel, shared.APP_LOG_FILE)
			default:
				appLogger = getLogger(zap.InfoLevel, shared.APP_LOG_FILE)
			}
		} else {
			defaultLogger()
		}
	}
	return appLogger
}

func defaultLogger() {
	appLogger = getLogger(zap.InfoLevel, shared.APP_LOG_FILE)
}

// Gets Zap logger with specified instance
//
// returns a SugaredLogger instance
func getLogger(level zapcore.Level, filename string) *zap.SugaredLogger {
	// Configure logger options
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(shared.DATETIME_FMT))
	}

	logFile, errLogFile := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if errLogFile != nil {
		defer logFile.Close()
		panic("Failed to open log file " + filename)
	}

	var err error
	_logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(config),
		zapcore.AddSync(logFile), level)).Sugar()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	return _logger
}
