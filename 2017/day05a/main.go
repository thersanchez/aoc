package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "wrong number of arguments")
		usage()
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	steps, err := do(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(steps)
}

func usage() {
	fmt.Fprintln(os.Stderr, "want only one argument, the name of the file with the program to load")
}
