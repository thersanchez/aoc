package main

func do(input string) int {
	return sum(dups(toIntSlice(clean(input))))
}

func clean(s string) string {
	b := []byte{}
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			b = append(b, s[i])
		}
	}
	return string(b)
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

func dups(nn []int) []int {
	ret := []int{}
	if len(nn) < 2 {
		return ret
	}
	for i := range nn {
		next := i + 1
		if i == len(nn)-1 {
			next = 0
		}
		if nn[i] == nn[next] {
			ret = append(ret, nn[i])
		}
	}
	return ret
}

func sum(ns []int) int {
	ret := 0
	for _, n := range ns {
		ret += n
	}
	return ret
}
