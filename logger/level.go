package logger

type Level int8

const (
	SILENT Level = iota
	ERROR
	WARN
	INFO
	DEBUG
)

func ToLevel[T float32 | float64 | int | Level](lvl T) Level {
	var level = float64(lvl)
	if level < float64(SILENT) {
		return SILENT
	} else if level > float64(DEBUG) {
		return DEBUG
	} else {
		return Level(level)
	}
}
