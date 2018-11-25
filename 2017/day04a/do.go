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
		if !hasDuplicatedWords(reader) {
			counter++
		}
	}
	if err := lineScanner.Err(); err != nil {
		log.Fatalf("scanning by lines: %v", err)
	}
	return counter
}

// HasDuplicatedWords returs if the line has duplicated words.
func hasDuplicatedWords(line io.Reader) bool {
	scanner := bufio.NewScanner(line)
	scanner.Split(bufio.ScanWords)
	wordsSeen := make(map[string]struct{})
	for scanner.Scan() {
		word := scanner.Text()
		_, seen := wordsSeen[word]
		if seen {
			return true
		}
		wordsSeen[word] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanning by words: %v", err)
	}
	return false
}
