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
	validPassPhases := do(file)
	fmt.Println(validPassPhases)
}

func usage() {
	fmt.Fprintln(os.Stderr, "\n\tday04a file_name")
}
