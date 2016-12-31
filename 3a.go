package main

import (
	"fmt"
	"sort"
)

func isTriangle(triangle []int) int {
	sort.Ints(triangle)
	if triangle[0]+triangle[1] <= triangle[2] {
		return 0
	} else {
		return 1
	}
}

func main() {
	var triangles [][]int
	var a, b, c int

	n, _ := fmt.Scan(&a, &b, &c)
	for n == 3 {
		triangles = append(triangles, []int{a, b, c})
		n, _ = fmt.Scan(&a, &b, &c)
	}

	count := 0
	for _, triangle := range triangles {
		count += isTriangle(triangle)
	}
	fmt.Println(count)
}
