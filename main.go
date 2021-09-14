package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// logger := get()
	// logger.trace("Starting the application...")
	// logger.info("Something noteworthy happened...")
	// logger.warn("There is something you should know about...")
	// logger.error("Something went wrong...")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type Logger struct {
	Trace *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

func (l *Logger) trace(message string) {
	l.Trace.Println(message)
}

func (l *Logger) info(message string) {
	l.Info.Println(message)
}

func (l *Logger) warn(message string) {
	l.Warn.Println(message)
}

func (l *Logger) error(message string) {
	l.Error.Println(message)
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
