package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Grid [][]bool

func (g Grid) String() (ans string) {
	for y := 0; y < len(g[0]); y++ {
		for x := 0; x < len(g); x++ {
			if g[x][y] {
				ans += "#"
			} else {
				ans += "."
			}
		}
		ans += "\n"
	}
	return
}

func (g Grid) rect(rows int, cols int) {
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			g[x][y] = true
		}
	}
}

func (g Grid) rotateRow(y int, amt int) {
	for ; amt > 0; amt-- {
		swap := g[len(g)-1][y]
		for i := len(g) - 1; i >= 0; i-- {
			if i == 0 {
				g[i][y] = swap
			} else {
				g[i][y] = g[i-1][y]
			}
		}
	}
}

func (g Grid) rotateCol(x int, amt int) {
	for ; amt > 0; amt-- {
		swap := g[x][len(g[x])-1]
		for i := len(g[x]) - 1; i >= 0; i-- {
			if i == 0 {
				g[x][i] = swap
			} else {
				g[x][i] = g[x][i-1]
			}
		}
	}
}

func (g Grid) process(cmd string) {
	reRect := regexp.MustCompile("rect (\\d+)x(\\d+)")
	matches := reRect.FindAllStringSubmatch(cmd, -1)
	if matches != nil {
		x, _ := strconv.Atoi(matches[0][1])
		y, _ := strconv.Atoi(matches[0][2])
		(&g).rect(x, y)
		return
	}

	reRotateRow := regexp.MustCompile("rotate row y=(\\d+) by (\\d+)")
	matches = reRotateRow.FindAllStringSubmatch(cmd, -1)
	if matches != nil {
		y, _ := strconv.Atoi(matches[0][1])
		amt, _ := strconv.Atoi(matches[0][2])
		(&g).rotateRow(y, amt)
		return
	}

	reRotateCol := regexp.MustCompile("rotate column x=(\\d+) by (\\d+)")
	matches = reRotateCol.FindAllStringSubmatch(cmd, -1)
	if matches != nil {
		x, _ := strconv.Atoi(matches[0][1])
		amt, _ := strconv.Atoi(matches[0][2])
		(&g).rotateCol(x, amt)
		return
	}

	panic("Could not process command: " + cmd)
}

func main() {
	var grid Grid
	for x := 0; x < 50; x++ {
		grid = append(grid, make([]bool, 6))
	}

	in := bufio.NewReader(os.Stdin)
	l, _ := in.ReadString('\n')
	for l != "" {
		(&grid).process(l)
		l, _ = in.ReadString('\n')
	}

	fmt.Print(grid)

	count := 0
	for _, col := range grid {
		for _, val := range col {
			if val {
				count++
			}
		}
	}
	fmt.Println(count)
}
