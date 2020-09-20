package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/thersanchez/aoc/2017/day06a/mem"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "wrong number of arguments")
		usage()
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	howMany, err := do(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(howMany)
}

func usage() {
	fmt.Fprintln(os.Stderr, "want only one argument, the name of the file with the banks")
}

func do(r io.Reader) (int, error) {
	howMany := 1

	banks, err := parseBanks(r)
	if err != nil {
		return 0, fmt.Errorf("parsing banks: %v", err)
	}

	m, err := mem.NewMem(banks)
	if err != nil {
		return 0, fmt.Errorf("creating memory: %v", err)
	}

	states := mem.NewStates()

	for {
		states.Add(m)
		mc := m.FindMostCrowded()

		if err := m.RedistributeBlocks(mc); err != nil {
			return 0, fmt.Errorf("redistributing blocks: %v", err)
		}

		if states.Has(m) {
			return howMany, nil
		}

		howMany++
	}
}

func parseBanks(r io.Reader) ([]int, error) {
	all, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("reading all: %v", err)
	}

	chunks := strings.Split(string(all), "\t")
	result := make([]int, len(chunks))

	for i, c := range chunks {
		result[i], err = strconv.Atoi(c)
		if err != nil {
			return nil, fmt.Errorf("parsing value #%d: %v", i+1, err)
		}
	}

	return result, nil
}
