package logrus

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	logger *log.Logger
	once   sync.Once
)

// Init initializes the global logrus logger with custom config
func Init() {
	formatter := new(prefixed.TextFormatter)

	formatter.ForceColors = true
	formatter.FullTimestamp = true
	// formatter.SpacePadding = 10
	formatter.TimestampFormat = "2006-01-02 15:04:05"

	// Set color scheme
	formatter.SetColorScheme(&prefixed.ColorScheme{
		TimestampStyle:  "white+bh",
		DebugLevelStyle: "cyan+bh",
		InfoLevelStyle:  "magenta+bh",
		WarnLevelStyle:  "yellow+bh",
		ErrorLevelStyle: "red+buh",
		FatalLevelStyle: "red+Bh",
		PanicLevelStyle: "red+Bh",
	})

	once.Do(func() {
		logger = log.New()
		logger.SetFormatter(formatter)
		logger.SetOutput(os.Stdout)
		logger.SetLevel(log.TraceLevel)
	})
}

// GetLogger returns the global logrus logger
func GetLogger() *log.Logger {
	if logger == nil {
		Init()
	}
	return logger
}
