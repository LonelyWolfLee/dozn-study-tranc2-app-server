package logger

import (
	"io/ioutil"
	"log"
	"os"
)

type Logger struct {
	trace *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func (l *Logger) Trace(message string) {
	l.trace.Println(message)
}

func (l *Logger) Info(message string) {
	l.info.Println(message)
}

func (l *Logger) Warn(message string) {
	l.warn.Println(message)
}

func (l *Logger) Error(message string) {
	l.error.Println(message)
}

func get() Logger {
	l := Logger{
		log.New(ioutil.Discard, "[TRACE] ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	return l
}

var LOG = get()
