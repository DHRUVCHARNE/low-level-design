package util

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadString(reader *bufio.Reader, prompt, defaultVal string) string {
	fmt.Printf("%s [%s] : ", prompt, defaultVal)
	input, _:= reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultVal
	}
	return input
}

func ReadInt(reader *bufio.Reader, prompt string, defaultVal int) int {
	fmt.Printf("%s [%d] : ", prompt, defaultVal)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == ""  {
		return defaultVal
	}
	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid Number, Using Default")
		return defaultVal
	}
	return val
}

func ReadFloat64(reader *bufio.Reader, prompt string, defaultVal float64) float64 {
	fmt.Printf("%s [%.2f] : ", prompt, defaultVal)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == ""  {
		return defaultVal
	}
	val, err := strconv.ParseFloat(input,64)
	if err != nil {
		fmt.Println("Invalid Number, Using Default")
		return defaultVal
	}
	return val
}
