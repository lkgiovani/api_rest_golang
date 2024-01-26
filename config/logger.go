package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer	io.Writer
}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	
)

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix,log.Ldate | log.Ltime)

	return &Logger{
		debug: log.New(writer, colorBlue+"DEBUG: ", logger.Flags()), 
		info:  log.New(writer, colorGreen+"INFO: ", logger.Flags()), 
		warning:log.New(writer, colorYellow+"WARNING: ", logger.Flags()), 
		err: log.New(writer, colorRed+"ERR: ", logger.Flags()),
		writer: writer,
	}
}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}){
	l.debug.Println(v...)

}

func (l *Logger) Info(v ...interface{}){
	l.info.Println(v...)

}

func (l *Logger) Warning(v ...interface{}){
	l.warning.Println(v...)

}

func (l *Logger) Err(v ...interface{}){
	l.err.Println(v...)
}

// Create Formatted Logs
func (l *Logger) DebugF(format string,v ...interface{}){
	l.debug.Printf(format, v...)
}

func (l *Logger) InfoF(format string,v ...interface{}){
	l.info.Printf(format, v...)
}

func (l *Logger) WarningF(format string,v ...interface{}){
	l.warning.Printf(format, v...)
}

func (l *Logger) ErrF(format string,v ...interface{}){
	l.err.Printf(format, v...)
}