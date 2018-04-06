package main

import "fmt"

func do(input string) int {
	return doubleSum(halfDups(toIntSlice(check(input))))
}

func check(s string) string {
	for i := 0; i < len(s); i++ {
		if !isDigit(s[i]) {
			msg := fmt.Sprintf("byte at position %d is not a digit", i+1)
			panic(msg)
		}
	}
	if s == "" {
		panic("empty string")
	}
	if len(s)%2 != 0 {
		panic("odd number")
	}
	return s
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func toIntSlice(s string) []int {
	ret := []int{}
	for _, r := range s {
		ret = append(ret, int(r-'0'))
	}
	return ret
}

// Assumes that len(nn) is an even number.
func halfDups(nn []int) []int {
	ret := []int{}
	if len(nn) < 2 {
		return ret
	}

	jump := len(nn) / 2
	for i := 0; i < len(nn)/2; i++ {
		if nn[i] == nn[i+jump] {
			ret = append(ret, nn[i])
		}
	}
	return ret
}

func doubleSum(ns []int) int {
	ret := 0
	for _, n := range ns {
		ret += n
	}
	return 2 * ret
}
