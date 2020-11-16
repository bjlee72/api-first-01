package logger

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"api-first-01/config"
)

// entry defines a log entry. This format is something that Stackdriver expects for the structured log.
type entry struct {
	Message  string `json:"message"`
	Severity string `json:"severity,omitempty"`
	Trace    string `json:"logging.googleapis.com/trace,omitempty"`

	// Stackdriver Log Viewer allows filtering and display of this as `jsonPayload.component`.
	Component string `json:"component,omitempty"`
}

// String renders an entry structure to the JSON format expected by Stackdriver.
func (e entry) String() string {
	if e.Severity == "" {
		e.Severity = "INFO"
	}
	out, err := json.Marshal(e)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
	}
	return string(out)
}

type logSeverity string

const (
	severityInfo     logSeverity = "INFO"
	severityError                = "ERROR"
	severityCritical             = "CRITICAL"
)

// Logf creates the string to log, logs it, and returns it.
func Logf(ctx context.Context, formatString string, args ...interface{}) string {
	errMessage := fmt.Sprintf(formatString, args...)
	logStructured(ctx, severityInfo, errMessage)
	return errMessage
}

// Log logs the given string, and returns it.
func Log(ctx context.Context, str string) string {
	logStructured(ctx, severityInfo, str)
	return str
}

// Fatalf logs the given string, and close the execution.
func Fatalf(ctx context.Context, formatString string, args ...interface{}) {
	errMessage := fmt.Sprintf(formatString, args...)
	logStructured(ctx, severityCritical, errMessage)
}

// Error creates the string to log, logs, and returns it as an severityError.
func Error(ctx context.Context, err interface{}) error {
	if e, ok := err.(error); ok {
		logStructured(ctx, severityError, e.Error())
		return e
	}
	if msg, ok := err.(string); ok {
		logStructured(ctx, severityError, msg)
		return errors.New(msg)
	}
	if sg, ok := err.(fmt.Stringer); ok {
		logStructured(ctx, severityError, sg.String())
		return errors.New(sg.String())
	}
	return fmt.Errorf("%#v", err)
}

// Errorf creates the string to log, logs, and returns it as an severityError.
// If one of the parameter is an severityError, that will be returned.
func Errorf(ctx context.Context, formatString string, args ...interface{}) error {
	errMessage := fmt.Sprintf(formatString, args...)
	logStructured(ctx, severityError, errMessage)
	err := errors.New(errMessage)
	for _, arg := range args {
		if e, ok := arg.(error); ok {
			err = e
			break
		}
	}
	return err
}

func logStructured(ctx context.Context, severity logSeverity, message string) {
	// if we use logging libraries, the timestamp is prefixed to the result.
	// Thus, it's not picked up as a structured log with JSON format.
	fmt.Fprintln(os.Stderr, entry{
		Message:   message,
		Severity:  string(severity),
		Trace:     getTrace(ctx),
		Component: getComponentString(getCaller()),
	})

	if severity == severityCritical {
		panic(message)
	}
}

func getTrace(ctx context.Context) string {
	return config.Trace(ctx)
}

// the files that have log related implementations.
// these files shouldn't be recorded in the structured log.
var loggerFiles = map[string]bool{
	"errors.go":       true,
	"log.go":          true,
	"response_log.go": true,
}

// return the getCaller information of a logger function.
func getCaller() (string, int, string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])

	var (
		frame runtime.Frame
		ok    bool
	)
	for {
		if frame, ok = frames.Next(); !ok {
			break
		}
		segments := strings.Split(strings.ToLower(frame.File), "/")
		fileName := segments[len(segments)-1]
		if _, ok := loggerFiles[fileName]; !ok {
			break
		}
	}

	return frame.File, frame.Line, frame.Function
}

func getComponentString(file string, line int, _ string) string {
	return fmt.Sprintf("%s:%d", file, line)
}
