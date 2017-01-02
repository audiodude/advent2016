package main

import (
	"fmt"
	"sort"
	"strings"
)

type LetterCount struct {
	Letter rune
	Count  int
}

func (lc *LetterCount) String() string {
	return fmt.Sprintf("%c: %d", lc.Letter, lc.Count)
}

type ByCount []*LetterCount

func (a ByCount) Len() int      { return len(a) }
func (a ByCount) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool {
	if a[i].Count == a[j].Count {
		return a[i].Letter < a[j].Letter
	} else {
		return a[i].Count > a[j].Count
	}
}

func roomCodeIfReal(room string) int {
	count := make(map[rune]int)
	parts := strings.Split(room, "-")

	var sectorId int
	var lock string
	for i, part := range parts {
		if i == len(parts)-1 {
			fmt.Sscanf(strings.Trim(part, "]"), "%d[%s]", &sectorId, &lock)
			continue
		}
		for _, r := range part {
			count[r]++
		}
	}
	var letters []*LetterCount
	for r, n := range count {
		letters = append(letters, &LetterCount{r, n})
	}
	sort.Sort(ByCount(letters))
	var key string
	for _, lc := range letters[:5] {
		key += fmt.Sprintf("%c", lc.Letter)
	}
	if key == lock {
		return sectorId
	} else {
		return 0
	}
}

func main() {
	var rooms []string
	var r string

	n, _ := fmt.Scan(&r)
	for n > 0 {
		rooms = append(rooms, r)
		n, _ = fmt.Scan(&r)
	}

	sum := 0
	for _, room := range rooms {
		sum += roomCodeIfReal(room)
	}
	fmt.Println(sum)
}
