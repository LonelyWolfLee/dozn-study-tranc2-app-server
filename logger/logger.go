package logger

import (
	"io/ioutil"
	"log"
	"os"
)

type Logger struct {
	Trace *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
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
