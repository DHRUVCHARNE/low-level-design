package main

import (
	"bufio"
	"fmt"
	"github.com/DHRUVCHARNE/low-level-design/cmd"
	"github.com/DHRUVCHARNE/low-level-design/internal/book/model"
	"os"
)

func main() {
	fmt.Println("Bank Account CLI")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---Creating a New Bank Account---")
	author := util.ReadString(reader, "Enter the Book Title", "Shekhar: A Life")
	title := util.ReadString(reader, "Enter the Name of the Author", "Agyeya")
	isbn := util.ReadString(reader, "Enter the ISBN", "978-0-143-42635-6")
	book := model.NewBook(author,title,isbn)
	fmt.Println("---New Book Added---")
	book.DisplayInfo()
	for {
		fmt.Println("Type 'd' for display, 'r' for returning the book and 'b' for borrowing the book and 'q' for exiting the cli")
		choice:=util.ReadByte(reader,"Enter the choice character",'q')
		switch choice {
		case 'd':
			book.DisplayInfo()
        case 'r':
			book.ReturnBook()
		case 'b':
			book.BorrowBook()
		case 'q':
			fmt.Println("---Exiting the Bank Account CLI---")
			return
		default:
			continue
		}
	}
}
