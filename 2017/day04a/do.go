package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func do(r io.Reader) {
	var counter int
	lineScanner := bufio.NewScanner(r)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		reader := strings.NewReader(line)
		if hasDuplicateWords(reader) {
			counter++
		}
	}
	if err := lineScanner.Err(); err != nil {
		log.Fatalf("scanning by lines: %v", err)
	}
	fmt.Println(counter)
}

func hasDuplicateWords(line io.Reader) bool {
	scanner := bufio.NewScanner(line)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanning by words: %v", err)
	}
	return true
}
