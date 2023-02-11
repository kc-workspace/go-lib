# Logger

loggers.DefaultPrinter *Printer
loggers.NewPrinter(writer *os.Writer)
loggers.SetPrinter(printer *Printer)

printer.Write(writer io.Writer, message interface{})
printer.Print(message interface{})

loggers.DefaultManager *Manager
loggers.NewManager(
  baseName []string,
  levels []Level,
  printer *Printer,
) *Manager

manager.SetLevel(level Level) *Manager
manager.New(name []string) *Logger // Always create new logger
manager.NewTable(size uint) *Table // Always create new table

logger.Debug(format string, msg interface{})
logger.Info(format string, msg interface{})
logger.Warn(format string, msg interface{})
logger.Error(err error)
logger.ErrorString(format string, msg interface{})
logger.Panic(err error)
logger.PanicString(format string, msg interface{})
logger.Log(format string, msg interface{})
logger.Line()
logger.NewLine()

logger.New(name []string) *Logger
logger.Extend(name []string) *Logger

table.IsInitial() bool
table.Size(size uint) *Table
table.NewRow(colume []string) *Table
table.End() error
