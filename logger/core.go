package logger

import (
	"math"
	"text/tabwriter"
)

func NewManager(
	baseNames []string,
	level Level,
	printer *Printer,
) *Manager {
	return &Manager{
		names:   baseNames,
		level:   level,
		printer: printer,
	}
}

func NewLogger(names []string, level Level, printer *Printer) *Logger {
	return &Logger{
		names:   names,
		level:   level,
		printer: printer,
	}
}

func NewTable(size uint, level Level, printer *Printer) *Table {
	var lineSize = len(LINE)
	var tab = int(float64(lineSize) / float64(size))
	var min = math.Min(float64(tab), float64(4))

	return &Table{
		size:    size,
		level:   level,
		printer: printer,
		writer: tabwriter.NewWriter(
			printer.writer,
			int(min),
			tab,
			2,
			' ',
			0,
		),
	}
}

var DefaultManager = NewManager(
	make([]string, 0),
	INFO,
	DefaultPrinter,
)
