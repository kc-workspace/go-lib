package logger

import "github.com/kc-workspace/go-lib/utils"

var level Level = INFO
var printer *Printer = NewDefaultPrinter()
var storage = make(map[string]*Logger)

func SetLevel[T float32 | float64 | int | Level](l T) {
	level = ToLevel(l)
}

func GetLevel() Level {
	return level
}

func GetPrinter() *Printer {
	return printer
}

func GetTable(size uint) *Table {
	return NewDefaultTable(size).Init()
}

func Get(names ...string) *Logger {
	name := utils.JoinString(":", names...)
	if storage[name] == nil {
		storage[name] = New(name, printer)
	}

	return storage[name]
}
