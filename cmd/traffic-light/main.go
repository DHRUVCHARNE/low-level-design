package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/traffic-light/model"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Traffic Light CLI")
	fmt.Println("------------------")

	// initial state
	light := model.RED

	for {
		model.Display(light)

		fmt.Println("\nOptions:")
		fmt.Println("n - Next signal")
		fmt.Println("q - Quit")

		choice := util.ReadString(reader, "Enter choice", "n")
		choice = strings.ToLower(choice)

		switch choice {
		case "n":
			light.Next()
		case "q":
			fmt.Println("Exiting Traffic Light CLI")
			return
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
	}
}