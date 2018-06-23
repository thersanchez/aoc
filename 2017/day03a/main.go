package main

import (
	"fmt"
	"os"
)

func main() {
	ret, err := Do(368078)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	fmt.Println(ret)
}
