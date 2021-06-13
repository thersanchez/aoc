package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		name, weight, children, err := ParseLine(scanner.Text())
		if err != nil {
			log.Fatalf("error parsing line: %v", err)
		}

		_, _, _ = name, weight, children
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failure scanning input file: %v", err)
	}
}
