package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	doorId := []byte("wtnhxymk")

	nFound := 0
	i := 0
	for nFound < 8 {
		hash := md5.Sum(append(doorId, []byte(fmt.Sprintf("%d", i))...))
		hexRep := fmt.Sprintf("%x", hash)
		if hexRep[:5] == "00000" {
			fmt.Printf("%c", hexRep[5])
			nFound++
		}
		i++
	}
	fmt.Println()
}
