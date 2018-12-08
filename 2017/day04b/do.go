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
// Therefore, an empty line is valid.
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
	fa := letterFrequencies(a)
	fb := letterFrequencies(b)
	return equalFrequencies(fa, fb)
}

func letterFrequencies(w string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range w {
		m[r] = m[r] + 1
	}
	return m
}

func equalFrequencies(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for r, fa := range a {
		fb, ok := b[r]
		if !ok {
			return false
		}
		if fa != fb {
			return false
		}
	}
	return true
}
