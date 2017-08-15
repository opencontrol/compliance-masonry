package logger

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

// internal type for our enums
type levels int

// emulate enum
const (
	FATAL   levels = -1
	ERROR   levels = 0
	WARNING levels = 1
	INFO    levels = 2
	DEBUG   levels = 3
	TRACE   levels = 4
	UNSUPP  levels = 99
)

// operator overload - cast enum to string
func (l levels) String() string {
	switch l {
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	default:
		return "UNSUPP"
	}
}

// Level - interface to wrap our enum - it cannot be inherited / subclassed
type Level interface {
	Levels() levels
}

// operator overload - access the underlying enum value
func (l levels) Levels() levels {
	return l
}

// LocalLogger - our custom logger structure
type LocalLogger struct {
	filename string
	*log.Logger
}

// instance singletons
var theLogger *LocalLogger
var once sync.Once

// GetInstance - start logging via singleton
func GetInstance() *LocalLogger {
	once.Do(func() {
		theDest := os.Getenv("DEBUG_LOG")
		if len(theDest) == 0 {
			theDest = "mylogger.log"
		}
		theLogger = createLogger(theDest)
	})
	return theLogger
}

// internal function to access logger (eck)
func createLogger(fname string) *LocalLogger {
	var file *os.File
	if fname != "-" {
		var err error
		file, err = os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			log.Fatalf("Unable to open %s for logging", fname)
		}
	} else {
		fname = "stderr"
		file = os.Stderr
	}

	return &LocalLogger{
		filename: fname,
		Logger:   log.New(file, "compliance-masonry: ", log.Lshortfile),
	}
}

// log and print with given stack and format
func internalPrint(stackLevels int, desired Level, format string, args ...interface{}) bool {
	if Check(desired) {
		var formattedStr = fmt.Sprintf(format, args...)
		theInstance := GetInstance()
		theInstance.Output(stackLevels, fmt.Sprintf("%s %s\n", desired.Levels(), formattedStr))
		return true
	}
	return false
}

// Check - should logging occur?
func Check(desired Level) bool {
	var envDebugStr = os.Getenv("DEBUG")
	envDebugVal, err := strconv.Atoi(envDebugStr)
	if err != nil {
		return false
	}
	var actual = UNSUPP
	switch envDebugVal {
	case int(FATAL):
		actual = FATAL
	case int(ERROR):
		actual = ERROR
	case int(WARNING):
		actual = WARNING
	case int(INFO):
		actual = INFO
	case int(DEBUG):
		actual = DEBUG
	case int(TRACE):
		actual = TRACE
	default:
		actual = UNSUPP
	}
	if actual.Levels() >= UNSUPP {
		return false
	}
	return (actual.Levels() >= desired.Levels())
}

// Print - log and print
func Print(desired Level, msg string) bool {
	return internalPrint(2, desired, "%s", msg)
}

// Trace - log and print
func Trace(msg string) bool {
	return internalPrint(3, TRACE, "%s", msg)
}

// Debug - log and print
func Debug(msg string) bool {
	return internalPrint(3, DEBUG, "%s", msg)
}

// Info - log and print
func Info(msg string) bool {
	return internalPrint(3, DEBUG, "%s", msg)
}

// Warning - log and print
func Warning(msg string) bool {
	return internalPrint(3, WARNING, "%s", msg)
}

// Error - log and print
func Error(msg string) bool {
	return internalPrint(3, ERROR, "%s", msg)
}

// Fatal - log and print
func Fatal(msg string) bool {
	return internalPrint(3, FATAL, "%s", msg)
}

// Printf - log and print
func Printf(desired Level, format string, args ...interface{}) bool {
	return internalPrint(2, desired, format, args...)
}

// Tracef - log and print
func Tracef(format string, args ...interface{}) bool {
	return internalPrint(3, TRACE, format, args...)
}

// Debugf - log and print
func Debugf(format string, args ...interface{}) bool {
	return internalPrint(3, DEBUG, format, args...)
}

// Infof - log and print
func Infof(format string, args ...interface{}) bool {
	return internalPrint(3, DEBUG, format, args...)
}

// Warningf - log and print
func Warningf(format string, args ...interface{}) bool {
	return internalPrint(3, WARNING, format, args...)
}

// Errorf - log and print
func Errorf(format string, args ...interface{}) bool {
	return internalPrint(3, ERROR, format, args...)
}

// Fatalf - log and print
func Fatalf(format string, args ...interface{}) bool {
	return internalPrint(3, FATAL, format, args...)
}
