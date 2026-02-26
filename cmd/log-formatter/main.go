package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/log-formatter/model"
)

func main() {
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Log Formatter CLI")
	message:= util.ReadString(reader,"Enter the Message","Server live on port 8080")
	formatter:=util.ReadInt(reader,"For Plain Formatter enter 1\nFor Json Formatter enter 2\nTo Exit enter any other key\n",1)
	switch formatter {
	case 1:
		logger := model.NewLogger(&model.PlainFormatter{})
		logger.Log(message)
	case 2:
		logger := model.NewLogger(&model.JsonFormatter{})
		logger.Log(message)
    default:
		fmt.Println("Exiting the CLI")
		return
	}


}