package main

import (
	"fmt"
	"math"
)

type Direction struct {
	Heading string
	Amount  int
}

type Coordinate struct {
	North, East int
}

type Location struct {
	North    int
	East     int
	NorthMod int
	EastMod  int
	Visited  map[Coordinate]bool
}

func (l *Location) process(d Direction) int {
	if d.Heading == "R" {
		turned := false
		switch l.NorthMod {
		case -1:
			l.EastMod = -1
			l.NorthMod = 0
			turned = true
		case 0:
			// Do nothing
		case 1:
			l.EastMod = 1
			l.NorthMod = 0
			turned = true
		}

		if !turned {
			switch l.EastMod {
			case -1:
				l.NorthMod = 1
				l.EastMod = 0
			case 0:
				panic("Location was pointing neither East nor North")
			case 1:
				l.NorthMod = -1
				l.EastMod = 0
			}
		}
	} else if d.Heading == "L" {
		turned := false
		switch l.NorthMod {
		case -1:
			l.EastMod = 1
			l.NorthMod = 0
			turned = true
		case 0:
			// Do nothing
		case 1:
			l.EastMod = -1
			l.NorthMod = 0
			turned = true
		}

		if !turned {
			switch l.EastMod {
			case -1:
				l.NorthMod = -1
				l.EastMod = 0
			case 0:
				panic("Location was pointing neither East nor North")
			case 1:
				l.NorthMod = 1
				l.EastMod = 0
			}
		}
	} else {
		panic("Direction wasn't L or R")
	}

	if l.NorthMod != 0 && l.EastMod != 0 {
		panic("Location facing direction is wrong!")
	}
	for i := 0; i < d.Amount; i++ {
		l.North += l.NorthMod
		l.East += l.EastMod
		curCoord := Coordinate{l.North, l.East}

		if _, revisited := l.Visited[curCoord]; revisited {
			fmt.Println(curCoord)
			return distance(*l)
		} else {
			l.Visited[curCoord] = true
		}
	}
	return 0
}

func distance(l Location) int {
	return int(math.Abs(float64(l.North)) + math.Abs(float64(l.East)))
}

func main() {
	var (
		d    string
		i    int
		n    int
		dirs []Direction
	)

	n, _ = fmt.Scanf("%1s%d", &d, &i)
	for n > 0 {
		dirs = append(dirs, Direction{d, i})
		n, _ = fmt.Scanf("%1s%d", &d, &i)
	}

	visited := make(map[Coordinate]bool)
	visited[Coordinate{0, 0}] = true
	result := Location{0, 0, 1, 0, visited}
	for _, d := range dirs {
		if dist := result.process(d); dist > 0 {
			fmt.Println(dist)
			return
		}
	}
	panic("Couldn't find a double visited location!")
}
