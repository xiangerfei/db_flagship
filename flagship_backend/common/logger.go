package common

import (
	"fmt"
	"log"
)

type Logger struct {
	log.Logger
}

func (logger *Logger) Debug(format string, a ...interface{}){
	logger.SetPrefix("DEBUG:")
	logger.Output(2, fmt.Sprintf(format, a...))
}


func (logger *Logger) Info(format string, a ...interface{}){
	logger.SetPrefix("INFO:")
	logger.Output(2, fmt.Sprintf(format, a...))
}

func (logger *Logger) Warning(format string, a ...interface{}){
	logger.SetPrefix("WARNING:")
	logger.Output(2, fmt.Sprintf(format, a...))
}

func (logger *Logger) Error(format string, a ...interface{}){
	logger.SetPrefix("ERROR:")
	logger.Output(2, fmt.Sprintf(format, a...))
}