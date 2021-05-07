package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	level := logrus.InfoLevel
	_, exists := os.LookupEnv("HEIMDALL_DEBUG")
	if exists {
		level = logrus.DebugLevel
	}
	logrus.SetLevel(level)
}
