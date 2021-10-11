package helpers

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(getLogLevel("LOG_LEVEL"))
}

func getLogLevel(envKey string) logrus.Level {
	switch level := os.Getenv(envKey); level {
	case "ERROR":
		Log.Info("Using log level: ERROR")
		return logrus.ErrorLevel
	case "WARN":
		Log.Info("Using log level: WARN")
		return logrus.WarnLevel
	case "INFO":
		Log.Info("Using log level: INFO")
		return logrus.InfoLevel
	case "DEBUG":
		Log.Info("Using log level: DEBUG")
		return logrus.DebugLevel
	default:
		Log.Info("The env variable (" + envKey + ") for the log level was not set. Using default.")
		Log.Info("Using log level: DEBUG")
		return logrus.DebugLevel
	}
}
