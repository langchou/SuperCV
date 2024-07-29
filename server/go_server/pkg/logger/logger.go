package logger

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	infoLog := &lumberjack.Logger{
		Filename:   "./logs/superCV.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}

	errorLog := &lumberjack.Logger{
		Filename:   "./logs/superCV.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}

	InfoLogger = log.New(io.MultiWriter(os.Stdout, infoLog), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(io.MultiWriter(os.Stderr, errorLog), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
