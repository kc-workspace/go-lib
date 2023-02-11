package logger

import (
	"math"
	"strings"
	"text/tabwriter"
)

type Table struct {
	level   Level
	size    uint
	printer *Printer
	writer  *tabwriter.Writer
}

func (t *Table) ToMsg(msg ...string) string {
	var str strings.Builder
	for i := 0; i < int(math.Min(float64(len(msg)), float64(t.size))); i++ {
		if i > 0 {
			str.WriteRune(TAB)
		}

		str.WriteString(msg[i])
	}

	return str.String()
}

func (t *Table) Row(msg ...string) *Table {
	t.printer.Write(t.writer, t.ToMsg(msg...))

	return t
}

func (t *Table) End() error {
	return t.writer.Flush()
}

func (l *Table) ToLogger(names ...string) *Logger {
	return NewLogger(names, l.level, l.printer)
}
