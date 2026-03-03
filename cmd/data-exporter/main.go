package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/data-exporter/model"
)

func main() {
	fmt.Println("Data Exporter CLI")
	reader:=bufio.NewReader(os.Stdin)
	names:=[]string{"Alice","Bob","Charlie"}
	data :=make([]string,0)
	for {
	choice := util.ReadByte(
			reader,
			"\nEnter 'a' to create an array"+
				"\nEnter 'c' to run CSV Explorer"+
				"\nEnter 'j' to run JSON Explorer"+
				"\nEnter any key to quit\n",
			'q',
		)
	switch choice {
	case 'a':
		data=util.ReadArrayOfStrings(reader,"Create the Array",names)
		fmt.Println("Data Array Added Successfully")
	case 'c':
		csv:=model.NewCSVExporter()
		csv.Export(data)
	case 'j':
		json:=model.NewJSONExporter()
		json.Export(data)
    default :
		fmt.Println("Exiting the CLI")
		return
	}
	}
}
