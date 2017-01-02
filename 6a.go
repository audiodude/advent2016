package main

import (
	"fmt"
)

func mostCommonLetter(count map[rune]int) (ans rune) {
	max := 0
	for r, n := range count {
		if n > max {
			ans = r
			max = n
		}
	}
	return
}

func main() {
	var line string

	n, _ := fmt.Scan(&line)
	counts := make([]map[rune]int, len(line))
	for i := 0; i < len(line); i++ {
		counts[i] = make(map[rune]int)
	}
	fmt.Println(counts)
	for n > 0 {
		for i, r := range line {
			counts[i][r]++
		}
		n, _ = fmt.Scan(&line)
	}

	var msg []rune
	for i := 0; i < len(counts); i++ {
		msg = append(msg, mostCommonLetter(counts[i]))
	}
	fmt.Printf("%c\n", msg)
}
