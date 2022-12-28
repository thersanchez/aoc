package main

import (
	"bufio"
	"errors"
	"fmt"
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

	// Note that if a node is unbalanced by a certain amount, all
	// its ancestors will also be unbalanced by the same amount.
	//
	// So, in order to find the program with the wrong weight,
	// we can find all unbalanced nodes and then the one with
	// the lowest weight.
	var lighter *day07b.Node
	{
		unbalanced := []*day07b.Node{}
		for _, node := range nodeSet {
			if !node.IsBalanced() {
				unbalanced = append(unbalanced, node)
			}
		}

		if len(unbalanced) == 0 {
			log.Fatal("no unbalanced nodes")
		}

		lighter = unbalanced[0]
		for _, n := range unbalanced[1:] {
			if lighter.TotalWeight() > n.TotalWeight() {
				lighter = n
			}
		}
	}

	correct, incorrect, err := childrenRepresentatives(lighter)
	if err != nil {
		log.Fatal(err)
	}

	result := correctedWeight(correct, incorrect)

	fmt.Println(result)
}

// Analises an unbalanced parent node to find out the children with the
// incorrect weight. It returns the incorrect child and one of the correct ones.
func childrenRepresentatives(parent *day07b.Node) (
	correct, incorrect *day07b.Node, err error,
) {
	children := parent.Children()

	// weightCounter will keep track of how many times each TotalWeight
	// happens for each children. We don't count the total amount of times
	// the correct weight shows up, just until we have seen enough weights
	// to have the correct and incorrect ones.
	//
	// For example, given children with the following weights:
	//
	//     [45, 45, 45, 45, 68, 45, 45, 45]
	//
	// the weightCounter map will have:
	//
	//     {45: 4, 68: 1}
	//
	// we will skip counting the last 3 elements.
	weightCounter := map[int]int{
		children[0].TotalWeight(): 1,
	}
	for i, c := range children[1:] {
		if _, ok := weightCounter[c.TotalWeight()]; ok {
			weightCounter[c.TotalWeight()] += 1
		} else {
			weightCounter[c.TotalWeight()] = 1
		}
		if i > 2 && len(weightCounter) > 1 {
			break
		}
	}

	// given a weightCounter like the one in the example above:
	//
	//     {45: 4, 68: 1}
	//
	// we want to find the incorrect node (the one with weight 68),
	// and any other node will be the correct one.
	var incorrectValue int
	for k, v := range weightCounter {
		if v == 1 {
			incorrectValue = k
		}
	}

	var incorrectNode, correctNode *day07b.Node
	for _, n := range children {
		if n.TotalWeight() == incorrectValue {
			incorrectNode = n
		} else {
			correctNode = n
		}

		if incorrectNode != nil && correctNode != nil {
			return correctNode, incorrectNode, nil
		}
	}

	return nil, nil, errors.New("balanced parent")
}

// Returns the weight that incorrectNode should have for its parent to be
// balanced.
func correctedWeight(correctNode, incorrectNode *day07b.Node) int {
	unbalanceAmount := correctNode.TotalWeight() - incorrectNode.TotalWeight()
	return incorrectNode.Weight() + unbalanceAmount
}
