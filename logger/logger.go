package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/kc-workspace/go-lib/utils"
)

type Logger struct {
	names   []string
	level   Level
	printer *Printer
}

func (l *Logger) format(lvl, format string, msg ...interface{}) string {
	// format syntax datetime: `02-01-2006 15:04:05`
	var datetime = time.Now().Format("15:04:05")

	var arr = make([]interface{}, 3)
	arr[0] = datetime
	arr[1] = GetName(l.names)
	arr[2] = strings.ToUpper(lvl)
	arr = append(arr, msg...)

	return fmt.Sprintf("%s %-20s [%-5s] | "+format, arr...)
}

func (l *Logger) valid(lvl Level) bool {
	return ShouldPrint(l.level, lvl)
}

func (l *Logger) SetLevel(lvl Level) *Logger {
	l.level = lvl
	return l
}

func (l *Logger) IsDebug() bool {
	return l.valid(DEBUG)
}

func (l *Logger) Debug(format string, msg ...interface{}) {
	if l.IsDebug() {
		l.printer.Print(l.format("debug", format, msg...))
	}
}

func (l *Logger) IsInfo() bool {
	return l.valid(INFO)
}

func (l *Logger) Info(format string, msg ...interface{}) {
	if l.IsInfo() {
		l.printer.Print(l.format("info", format, msg...))
	}
}

func (l *Logger) IsWarn() bool {
	return l.valid(WARN)
}

func (l *Logger) Warnf(format string, msg ...interface{}) {
	if l.IsWarn() {
		l.printer.Print(l.format("warn", format, msg...))
	}
}

func (l *Logger) Warn(err error) {
	if err != nil {
		l.Warnf(err.Error())
	}
}

func (l *Logger) IsError() bool {
	return l.valid(ERROR)
}

func (l *Logger) Errorf(format string, msg ...interface{}) {
	if l.IsError() {
		l.printer.Print(l.format("warn", format, msg...))
	}
}

func (l *Logger) Error(err error) {
	if err != nil {
		l.Errorf(err.Error())
	}
}

func (l *Logger) Panicf(format string, msg ...interface{}) {
	panic(l.format("panic", format, msg...))
}

func (l *Logger) Panic(err error) {
	if err != nil {
		l.Panicf(err.Error())
	}
}

func (l *Logger) Log(format string, msg ...interface{}) {
	if l.level != SILENT {
		l.printer.Print(fmt.Sprintf(format, msg...))
	}
}

func (l *Logger) Line() {
	l.Log(LINE)
}

func (l *Logger) NewLine() {
	l.Log(EMPTY)
}

func (l *Logger) New(names ...string) *Logger {
	return &Logger{
		names:   names,
		level:   l.level,
		printer: l.printer,
	}
}

func (l *Logger) Extend(names ...string) *Logger {
	return l.New(utils.CloneArray(l.names, names...)...)
}

func (l *Logger) ToTable(size uint) *Table {
	return NewTable(size, l.level, l.printer)
}
