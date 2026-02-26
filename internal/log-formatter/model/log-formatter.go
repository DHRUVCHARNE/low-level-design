package model

import "fmt"

type Formatter interface {
	Format(message string) string
}

type PlainFormatter struct {}
type JsonFormatter struct {}

func (p *PlainFormatter) Format(message string) string {
	return message
}

func (j *JsonFormatter) Format(message string) string {
	return fmt.Sprintf(`{"log":"%s"}`,message)
}

type Logger struct {
	formatter Formatter
}

func NewLogger(formatter Formatter) *Logger {
	return &Logger{formatter: formatter}
}

func  (logger *Logger) Log(message string) {
	res:=logger.formatter.Format(message)
	fmt.Println(res)
}

