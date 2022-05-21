package main

import (
	"bufio"
	"log"
	"os"

	"github.com/thersanchez/aoc/2017/day07a"
	"github.com/thersanchez/aoc/2017/day07b"
)

func main() {
	nodeSet := map[string]*day07b.Node{}
	nodeNamesSet := map[string]struct{}{}
	childrenNamesSet := map[string]struct{}{}

	// first parse of the file:
	// - add all node names to nodeNamesSet
	// - add all children names to childrenNamesSet
	// - create all nodes and add them to nodeSet
	{
		r, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			id, weight, children, err := day07a.ParseLine(scanner.Text())
			if err != nil {
				log.Fatalf("error parsing line: %v", err)
			}

			nodeNamesSet[id] = struct{}{}
			for _, c := range children {
				childrenNamesSet[c] = struct{}{}
			}

			n, err := day07b.NewNode(id, weight)
			if err != nil {
				log.Fatalf("error creating node: %v", err)
			}

			nodeSet[id] = n
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("failure scanning input file: %v", err)
		}
	}

	// second file parse:
	// - build the tree from the nodes in the nodeSet
	{
		r, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			id, _, children, err := day07a.ParseLine(scanner.Text())
			if err != nil {
				log.Fatalf("error parsing line: %v", err)
			}

			parent, ok := nodeSet[id]
			if !ok {
				log.Fatalf("node %s not found in nodeSet", id)
			}

			for _, childID := range children {
				c, ok := nodeSet[childID]
				if !ok {
					log.Fatalf("child %s not found in nodeSet", childID)
				}

				parent.AddChildren(c)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("failure scanning input file: %v", err)
		}
	}

	rootNameSet := day07a.SetSubtract(nodeNamesSet, childrenNamesSet)

	if len(rootNameSet) == 0 {
		log.Fatalf("malformed tree: didn't found any root")
	}
	if len(rootNameSet) > 1 {
		log.Fatalf("malformed tree: it is a forest")
	}

	var rootName string
	for k := range rootNameSet {
		rootName = k
		break
	}

	root, ok := nodeSet[rootName]
	if !ok {
		log.Fatalf("can't find root node %s in nodeSet", rootName)
	}

	root.TotalWeight()
}
