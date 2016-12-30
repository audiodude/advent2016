package main

import (
	"fmt"
	"math"
)

type Direction struct {
	Heading string
	Amount  int
}

type Location struct {
	North    int
	East     int
	NorthMod int
	EastMod  int
}

func (l *Location) process(d Direction) {
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
	l.North += l.NorthMod * d.Amount
	l.East += l.EastMod * d.Amount
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

	result := Location{0, 0, 1, 0}
	for _, d := range dirs {
		fmt.Println(result)
		result.process(d)
	}
	fmt.Println(result)
	fmt.Println(int(
		math.Abs(float64(result.North)) + math.Abs(float64(result.East))))
}
