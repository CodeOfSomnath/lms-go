package main

import (
	"bufio"
	"fmt"
	lib "github.com/lms/lms_parser"
	"os"
	"runtime"
)

func input(prompt string, inp *string) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*inp = scanner.Text()

}
func main() {
	var inp string
	var library = lib.NewLibrary()
	fmt.Println(" 1.Add\n 2.Show\n 3.Get\n 4.Exit\n 5.Option")
	for {
		input("Enter option: ", &inp)
		if inp == "1" {
			// adding book
			var name, author string
			//fmt.Scanln(&name)
			//fmt.Scanln(&author)
			input("Enter book name: ", &name)
			input("Enter author name: ", &author)
			library.Add(lib.NewBook(name, author))
			runtime.GC()
		} else if inp == "2" {
			// printing book
			fmt.Println(library)
		} else if inp == "3" {
			// giving book
			var name, author string

			input("Enter book name: ", &name)
			input("Enter author name: ", &author)
			var book = library.Get(name, author)
			if book.Compare("None", "None") {
				fmt.Printf("'%v' by '%v', book not found.\n", name, author)
			} else {
				fmt.Printf("'%v' by '%v', book giving you.\n", book.Name, book.Author)
			}
			runtime.GC()
		} else if inp == "4" {
			// for exit loop
			break
		} else if inp == "5" {
			// printing option
			fmt.Println(" 1.Add\n 2.Show\n 3.Get\n 4.Exit\n 5.Option")
		} else {
			// invalid option
			fmt.Println("Please enter correct option")
		}

	}

}
