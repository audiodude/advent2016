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
	t := make([]int, 9)

	n, _ := fmt.Scan(&t[0], &t[1], &t[2], &t[3], &t[4], &t[5], &t[6], &t[7],
		&t[8])
	for n == 9 {
		triangles = append(triangles, []int{t[0], t[3], t[6]})
		triangles = append(triangles, []int{t[1], t[4], t[7]})
		triangles = append(triangles, []int{t[2], t[5], t[8]})
		n, _ = fmt.Scan(&t[0], &t[1], &t[2], &t[3], &t[4], &t[5], &t[6], &t[7],
			&t[8])
	}

	count := 0
	for _, triangle := range triangles {
		count += isTriangle(triangle)
	}
	fmt.Println(count)
}
