package logger

import (
	"google-rtb/config"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *Logger
var printedWarning bool

type Logger struct {
	serviceName    string
	embeddedLogger *logrus.Logger
}

type LogParams struct {
	params map[string]interface{}
}

func NewLogParams() *LogParams {
	return &LogParams{}
}

func Init() {
	embeddedLogger := logrus.New()

	f, err := os.OpenFile(
		config.Cfg.Logger.FileLocation,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644,
	)

	if err != nil {
		panic(err)
	}

	embeddedLogger.SetOutput(f)

	logger = &Logger{}
	logger.serviceName = config.Cfg.ServerConfigurations.InstanceName
	logger.embeddedLogger = embeddedLogger
	logger.setFormatter()

	var logLevel string
	switch config.Cfg.Logger.Level {
	case "warn":
		logLevel = "WARN"
	case "info":
		logLevel = "INFO"
	default:
		logLevel = "DEBUG"
	}

	SetLogLevel(logLevel)
}

func SetLogLevel(level string) {
	if logger == nil {
		panic("logger is null")
	}

	embeddedLogger := logger.embeddedLogger
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		ErrorP(fmt.Sprintf("unable to set level: %s", level), GetErrorLogParams(err))
		return
	}

	fmt.Printf("setting log level: %s", level)
	embeddedLogger.SetLevel(parsedLevel)
}

func GetErrorLogParams(err error) *LogParams {
	logParams := &LogParams{}
	logParams.Add("error.message", err.Error())

	return logParams
}

func (logParams *LogParams) Add(key string, value interface{}) *LogParams {
	if logParams.params == nil {
		logParams.params = make(map[string]interface{})
	}
	logParams.params[key] = value

	return logParams
}

func InfoP(msg string, logParams *LogParams) {
	getLogger().embeddedLogger.WithFields(getDefaultParams(logParams, "INFO").params).Info(msg)
}

func Info(msg string) {
	InfoP(msg, nil)
}

func ErrorP(msg string, logParams *LogParams) {
	getLogger().embeddedLogger.WithFields(getDefaultParams(logParams, "ERROR").params).Error(msg)

}

func Error(msg string) {
	ErrorP(msg, nil)
}

func WarnP(msg string, logParams *LogParams) {
	getLogger().embeddedLogger.WithFields(getDefaultParams(logParams, "WARN").params).Warn(msg)
}

func Warn(msg string) {
	WarnP(msg, nil)
}

func DebugP(msg string, logParams *LogParams) {
	getLogger().embeddedLogger.WithFields(getDefaultParams(logParams, "DEBUG").params).Debug(msg)
}

func Debug(msg string) {
	DebugP(msg, nil)
}

func getLogger() *Logger {
	if logger == nil {
		defaultLogger := &Logger{
			serviceName:    "default",
			embeddedLogger: logrus.StandardLogger(),
		}
		if !printedWarning {
			defaultLogger.embeddedLogger.Warn("using default logger")
			printedWarning = true
		}

		return defaultLogger
	}

	return logger
}

func (logger *Logger) setFormatter() {
	logger.embeddedLogger.SetFormatter(&logrus.JSONFormatter{})
}

func getDefaultParams(logParams *LogParams, severity string) *LogParams {
	if logParams == nil {
		logParams = &LogParams{}
	}

	if logger != nil {
		logParams.Add("service", logger.serviceName) // add service name to each log
		logParams.Add("severity", severity)          // cloud logs native setting
	}

	return logParams
}
