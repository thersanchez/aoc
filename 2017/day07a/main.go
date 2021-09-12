package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	nodeSet := map[string]struct{}{}
	childrenSet := map[string]struct{}{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		name, _, children, err := ParseLine(scanner.Text())
		if err != nil {
			log.Fatalf("error parsing line: %v", err)
		}

		nodeSet[name] = struct{}{}
		for _, c := range children {
			childrenSet[c] = struct{}{}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failure scanning input file: %v", err)
	}

	roots := setSubtract(nodeSet, childrenSet)

	if len(roots) == 0 {
		log.Fatalf("malformed tree: didn't found any root")
	}
	if len(roots) > 1 {
		log.Fatalf("malformed tree: it is a forest")
	}

	// print the only element in roots
	for name, _ := range roots {
		fmt.Println(name)
		break
	}
}
