package main

import (
	"fmt"
	"regexp"
)

func hasAbba(seq string) bool {
	i := 0
	for i <= len(seq)-4 {
		if seq[i] != seq[i+1] && seq[i] == seq[i+3] && seq[i+1] == seq[i+2] {
			return true
		}
		i++
	}
	return false
}

func isTls(ip string) int {
	re := regexp.MustCompile("(\\[\\w+\\])")
	pairs := re.FindAllStringSubmatchIndex(ip, -1)
	if pairs != nil {
		for _, pair := range pairs {
			if hasAbba(ip[pair[0]+1 : pair[1]-1]) {
				return 0
			}
		}
		for i := 0; i <= len(pairs); i++ {
			if i == 0 {
				if hasAbba(ip[0:pairs[i][0]]) {
					return 1
				}
			} else if i == len(pairs) {
				if hasAbba(ip[pairs[i-1][1]:len(ip)]) {
					return 1
				}
			} else {
				if hasAbba(ip[pairs[i-1][1]:pairs[i][0]]) {
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
	for _, ip := range lines {
		count += isTls(ip)
	}
	fmt.Println(count)
}
