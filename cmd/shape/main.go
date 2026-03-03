package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/shape/model"
)

func main() {
	fmt.Println("Shapes CLI")
	reader:=bufio.NewReader(os.Stdin)
    shapes:=make([]model.Shape,0)
	for {
	choice := util.ReadByte(
			reader,
			"\nEnter 'c' to create Circle"+
				"\nEnter 'r' to create Rectangle"+
				"\nEnter 'a' to list all Shapes"+
				"\nEnter any key to quit\n",
			'q',
		)
	switch choice {
	case 'c':
		radius:=util.ReadFloat64(reader,"Enter radius",1.0)
		circle:=model.NewCircle(radius)
		shapes=append(shapes, circle)
		fmt.Println("Circle Added Successfully")
	case 'r':
		width:=util.ReadFloat64(reader,"Enter Width",1.0)
		height:=util.ReadFloat64(reader,"Enter Height",1.0)
		rectangle:=model.NewRectangle(width,height)
		shapes=append(shapes, rectangle)
		fmt.Println("Rectangle Added Successfully")
	case 'a':
		if len(shapes) == 0 {
			fmt.Println("No Shapes Created yet.")
			continue
		}
		fmt.Println("\nAll Shapes: ")
		for i,shape:= range shapes {
			fmt.Printf("%d.",i+1)
			shape.Describe()
		}
    case 'q':
		fmt.Println("Exiting the CLI")
		return
	}
	}
}