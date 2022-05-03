package common

import (
	"fmt"
	"log"
)

type Logger struct {
	log.Logger
}

func (logger *Logger) Info(format string, a ...interface{}){

	logger.Output(2, "INFO:" + fmt.Sprintf(format, a...))
}

func (logger *Logger) Warning(format string, a ...interface{}){
	logger.Output(2, "Warning:" + fmt.Sprintf(format, a...))
}

func (logger *Logger) Error(format string, a ...interface{}){
	logger.Output(2, "Error:" + fmt.Sprintf(format, a...))
}