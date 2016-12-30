package main

import (
	"fmt"
)

type Digit struct {
	Value int
	Up    *Digit
	Down  *Digit
	Left  *Digit
	Right *Digit
}

type Keypad []Digit

func makeKeypad() Keypad {
	keypad := make(Keypad, 10)

	for i := 0; i < len(keypad); i++ {
		keypad[i].Value = i
	}

	keypad[1].Right = &keypad[2]
	keypad[2].Right = &keypad[3]
	keypad[4].Right = &keypad[5]
	keypad[5].Right = &keypad[6]
	keypad[7].Right = &keypad[8]
	keypad[8].Right = &keypad[9]

	keypad[1].Down = &keypad[4]
	keypad[2].Down = &keypad[5]
	keypad[3].Down = &keypad[6]
	keypad[4].Down = &keypad[7]
	keypad[5].Down = &keypad[8]
	keypad[6].Down = &keypad[9]

	keypad[4].Up = &keypad[1]
	keypad[5].Up = &keypad[2]
	keypad[6].Up = &keypad[3]
	keypad[7].Up = &keypad[4]
	keypad[8].Up = &keypad[5]
	keypad[9].Up = &keypad[6]

	keypad[2].Left = &keypad[1]
	keypad[3].Left = &keypad[2]
	keypad[5].Left = &keypad[4]
	keypad[6].Left = &keypad[5]
	keypad[8].Left = &keypad[7]
	keypad[9].Left = &keypad[8]

	return keypad
}

func (keypad Keypad) nextNumber(code int, input string) int {
	place := code
	for _, c := range input {
		switch c {
		case 'U':
			if keypad[place].Up != nil {
				place = keypad[place].Up.Value
			}
		case 'D':
			if keypad[place].Down != nil {
				place = keypad[place].Down.Value
			}
		case 'L':
			if keypad[place].Left != nil {
				place = keypad[place].Left.Value
			}
		case 'R':
			if keypad[place].Right != nil {
				place = keypad[place].Right.Value
			}
		}
	}
	return place
}

func main() {
	var lines []string
	var l string

	n, _ := fmt.Scanln(&l)
	for n > 0 {
		lines = append(lines, l)
		n, _ = fmt.Scanln(&l)
	}

	keypad := makeKeypad()
	code := 5
	for _, in := range lines {
		code = keypad.nextNumber(code, in)
		fmt.Print(code)
	}
	fmt.Println()
}
