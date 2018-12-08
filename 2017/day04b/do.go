package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// Do count invalid passphrases in r (one passphrase per line).
// Invalid passphrases are the ones with anagrams.
func do(r io.Reader) int {
	var counter int
	lineScanner := bufio.NewScanner(r)
	for lineScanner.Scan() {
		line := lineScanner.Text()
		reader := strings.NewReader(line)
		if isValidPassphrase(reader) {
			counter++
		}
	}
	if err := lineScanner.Err(); err != nil {
		log.Fatalf("scanning by lines: %v", err)
	}
	return counter
}

// isValidPassphrase returs if the line has no anagrams.
// Also, an empty line is invalid.
func isValidPassphrase(line io.Reader) bool {
	scanner := bufio.NewScanner(line)
	scanner.Split(bufio.ScanWords)

	wordsSeen := make(map[string]struct{})
	for scanner.Scan() {
		word := scanner.Text()
		for seen := range wordsSeen {
			if areAnagrams(word, seen) {
				return false
			}
		}
		wordsSeen[word] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanning by words: %v", err)
	}
	return true
}

func areAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	return false
}
func letterFrequencies(w string) map[rune]int {
	return nil
}

func equalFrequencies(a, b map[rune]int) bool {
	return true
}
