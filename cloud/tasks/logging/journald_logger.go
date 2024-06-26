package logging

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/coreos/go-systemd/journal"
	"github.com/ydb-platform/nbs/library/go/core/log"
)

////////////////////////////////////////////////////////////////////////////////

// callerSkipOffset skips the following stacktrace:
// -> journaldLogger.Trace
// -> journaldLogger.structuredLog
// -> journaldLogger.structuredLogWithCaller
const callerSkipOffset = 4

////////////////////////////////////////////////////////////////////////////////

func getPriority(level Level) journal.Priority {
	switch level {
	case TraceLevel:
		return journal.PriDebug
	case DebugLevel:
		return journal.PriDebug
	case InfoLevel:
		return journal.PriInfo
	case WarnLevel:
		return journal.PriWarning
	case ErrorLevel:
		return journal.PriErr
	case FatalLevel:
		return journal.PriCrit
	}

	panic(fmt.Errorf("unknown log level %v", level))
}

////////////////////////////////////////////////////////////////////////////////

func getFields(fields ...log.Field) map[string]string {
	if len(fields) == 0 {
		return nil
	}

	vars := make(map[string]string)
	for _, f := range fields {
		// The variable name must be in uppercase and consist only of characters,
		// numbers and underscores, and may not begin with an underscore:
		// https://www.freedesktop.org/software/systemd/man/sd_journal_print.html
		key := strings.ToUpper(f.Key())
		vars[key] = fmt.Sprintf("%v", f.Any())
	}

	return vars
}

// Returns the caller of the log methods.
func captureStacktrace(skip int) (string, int, string, bool) {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(skip+1, rpc)
	if n < 1 {
		return "", 0, "", false
	}
	frame, _ := runtime.CallersFrames(rpc).Next()
	return frame.File, frame.Line, frame.Function, true
}

////////////////////////////////////////////////////////////////////////////////

type journaldLogger struct {
	level      Level
	name       string
	callerSkip int
}

// Sends log entity to journal, formatting the message as follows:
//
//	(*) caller function in format <filename>:<line_no> is prepended to the message and added as a `CALLER` var.
//	(*) if a logger has a name, name is prepended to the message.
//	(*) if there is an error field, error contents are appended to the message.
func (l *journaldLogger) structuredLogWithCaller(
	level Level,
	msg string,
	fields ...log.Field,
) {

	if len(l.name) > 0 {
		msg = l.name + " => " + msg
	}

	if fileName, lineNumber, function, ok := captureStacktrace(l.callerSkip + callerSkipOffset); ok {
		// See: https://www.freedesktop.org/software/systemd/man/systemd.journal-fields.html#CODE_FILE=
		fields = append(fields,
			[]log.Field{
				log.String("code_file", fileName),
				log.Int("code_line", lineNumber),
				log.String("code_func", function),
			}...,
		)

		// Short link to the file and line within the file.
		caller := fmt.Sprintf("%s:%d", path.Base(fileName), lineNumber)
		msg = caller + " " + msg
	}

	for _, f := range fields {
		if f.Key() == log.DefaultErrorFieldName && f.Type() == log.FieldTypeError {
			msg += " => " + f.Error().Error()
		}
	}

	err := journal.Send(msg, getPriority(level), getFields(fields...))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to journald %v\n", err)
	}
}

// Writes log with the fields.
// This method must always be called directly by a method in the Logger interface (e.g. Trace, Debug, Fatal).
// This skips structuredLog and Trace/Debug/Fatal methods that called it.
func (l *journaldLogger) structuredLog(
	level Level,
	msg string,
	fields ...log.Field,
) {

	if level < l.level {
		return
	}
	l.structuredLogWithCaller(level, msg, fields...)
}

func (l *journaldLogger) Trace(msg string, fields ...log.Field) {
	l.structuredLog(TraceLevel, msg, fields...)
}

func (l *journaldLogger) Debug(msg string, fields ...log.Field) {
	l.structuredLog(DebugLevel, msg, fields...)
}

func (l *journaldLogger) Info(msg string, fields ...log.Field) {
	l.structuredLog(InfoLevel, msg, fields...)
}

func (l *journaldLogger) Warn(msg string, fields ...log.Field) {
	l.structuredLog(WarnLevel, msg, fields...)
}

func (l *journaldLogger) Error(msg string, fields ...log.Field) {
	l.structuredLog(ErrorLevel, msg, fields...)
}

func (l *journaldLogger) Fatal(msg string, fields ...log.Field) {
	l.structuredLog(FatalLevel, msg, fields...)
}

// Formats log message.
// This method must always be called directly by a formatting method in the Logger interface (e.g. Tracef, Debugf, Fatalf).
// This skips fmtLog and the Trace/Debug/Fatal methods that called it.
func (l *journaldLogger) fmtLog(
	level Level,
	format string,
	args ...interface{},
) {

	if level < l.level {
		return
	}
	l.structuredLogWithCaller(level, fmt.Sprintf(format, args...))
}

func (l *journaldLogger) Tracef(format string, args ...interface{}) {
	l.fmtLog(TraceLevel, format, args...)
}

func (l *journaldLogger) Debugf(format string, args ...interface{}) {
	l.fmtLog(DebugLevel, format, args...)
}

func (l *journaldLogger) Infof(format string, args ...interface{}) {
	l.fmtLog(InfoLevel, format, args...)
}

func (l *journaldLogger) Warnf(format string, args ...interface{}) {
	l.fmtLog(WarnLevel, format, args...)
}

func (l *journaldLogger) Errorf(format string, args ...interface{}) {
	l.fmtLog(ErrorLevel, format, args...)
}

func (l *journaldLogger) Fatalf(format string, args ...interface{}) {
	l.fmtLog(FatalLevel, format, args...)
}

func (l *journaldLogger) Logger() log.Logger {
	return l
}

func (l *journaldLogger) Structured() log.Structured {
	return l
}

func (l *journaldLogger) Fmt() log.Fmt {
	return l
}

func (l *journaldLogger) WithName(name string) log.Logger {
	if len(l.name) > 0 {
		name = l.name + "." + name
	}

	return &journaldLogger{
		level: l.level,
		name:  name,
	}
}

func (l *journaldLogger) AddCallerSkip(skip int) log.Logger {
	return &journaldLogger{
		level:      l.level,
		name:       l.name,
		callerSkip: l.callerSkip + skip,
	}
}

////////////////////////////////////////////////////////////////////////////////

func NewJournaldLogger(level Level) Logger {
	if !journal.Enabled() {
		panic(fmt.Errorf("cannot initialize journald logger"))
	}

	return &journaldLogger{
		level: level,
	}
}
