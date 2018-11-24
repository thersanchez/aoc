package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// Do count invalid passphrases in r (one passphrase per line).
// Invalid passphrases are the ones with duplicated words.
func do(r io.Reader) int {
	var counter int
	lineScanner := bufio.NewScanner(r)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		reader := strings.NewReader(line)
		if !hasDuplicateWords(reader) {
			counter++
		}
	}
	if err := lineScanner.Err(); err != nil {
		log.Fatalf("scanning by lines: %v", err)
	}
	return counter
}

func hasDuplicateWords(line io.Reader) bool {
	scanner := bufio.NewScanner(line)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanning by words: %v", err)
	}
	// todo
	return true
}
