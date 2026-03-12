package main



import (
	"bufio"
	"fmt"
	"os"
	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/shape-hierarchy/model"
)

func main() {

	fmt.Println("Multiple Shape CLI")

	reader := bufio.NewReader(os.Stdin)

	shapes := map[int]model.Shape{}
	i := 1

	for {

		choice := util.ReadByte(
			reader,
			"\nEnter 'c' to create a shape"+
				"\nEnter 'l' to list shapes"+
				"\nEnter any other key to exit\n",
			'q',
		)

		switch choice {

		case 'c':

			shape, err := CreateShape(reader)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			if shape != nil {
				shapes[i] = shape
				fmt.Println("Shape created successfully with id:", i)
				i++
			}

		case 'l':

			if len(shapes) == 0 {
				fmt.Println("No shapes found")
				continue
			}

			fmt.Println("\nShapes:")

			for id, s := range shapes {
				fmt.Printf("ID: %d ", id)
				model.Describe(s)
				fmt.Println()
			}

		default:
			fmt.Println("Exiting CLI")
			return
		}
	}
}

func CreateShape(reader *bufio.Reader) (model.Shape, error) {

	choice := util.ReadByte(
		reader,
		"\nEnter 'c' for Circle"+
			"\nEnter 'r' for Rectangle"+
			"\nAny other key to cancel\n",
		'q',
	)

	switch choice {

	case 'c':

		radius := util.ReadFloat64(reader, "Enter radius", 1)

		if radius <= 0 {
			return nil, fmt.Errorf("radius must be positive")
		}

		return model.NewCircle(radius), nil

	case 'r':

		width := util.ReadFloat64(reader, "Enter width", 1)
		height := util.ReadFloat64(reader, "Enter height", 1)

		if width <= 0 || height <= 0 {
			return nil, fmt.Errorf("width/height must be positive")
		}

		return model.NewRectangle(width, height), nil

	default:
		return nil, nil
	}
}