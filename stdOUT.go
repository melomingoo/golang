package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)

}
