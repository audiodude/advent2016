package main

import (
	"fmt"
	"regexp"
)

func hasAba(seq string) (abas [][]byte) {
	i := 0
	for i <= len(seq)-3 {
		if seq[i] != seq[i+1] && seq[i] == seq[i+2] {
			abas = append(abas, []byte{seq[i], seq[i+1]})
		}
		i++
	}
	return
}

func hasBab(seq string, b byte, a byte) bool {
	i := 0
	for i <= len(seq)-3 {
		if seq[i] == b && seq[i+1] == a && seq[i] == seq[i+2] {
			return true
		}
		i++
	}
	return false
}

func isTls(ip string) int {
	re := regexp.MustCompile("(\\[\\w+\\])")
	pairs := re.FindAllStringSubmatchIndex(ip, -1)
	var abas [][]byte

	if pairs != nil {
		for i := 0; i <= len(pairs); i++ {
			if i == 0 {
				abas = append(abas, hasAba(ip[0:pairs[i][0]])...)

			} else if i == len(pairs) {
				abas = append(abas, hasAba(ip[pairs[i-1][1]:len(ip)])...)

			} else {
				abas = append(abas, hasAba(ip[pairs[i-1][1]:pairs[i][0]])...)
			}
		}
		for _, pair := range pairs {
			for _, aba := range abas {
				if hasBab(ip[pair[0]+1:pair[1]-1], aba[1], aba[0]) {
					return 1
				}
			}
		}
	}
	return 0
}

func main() {
	var lines []string
	var l string

	n, _ := fmt.Scan(&l)
	for n > 0 {
		lines = append(lines, l)
		n, _ = fmt.Scan(&l)
	}

	count := 0

	isTls("abxbpqr[abcxbxcba]jkfkfyz")
	for _, ip := range lines {
		count += isTls(ip)
	}
	fmt.Println(count)
}
