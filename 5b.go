package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	doorId := []byte("wtnhxymk")
	password := make([]byte, 8)
	spotTaken := make(map[int]bool)

	nFound := 0
	i := 0
	for nFound < 8 {
		hash := md5.Sum(append(doorId, []byte(fmt.Sprintf("%d", i))...))
		i++

		hexRep := fmt.Sprintf("%x", hash)
		if hexRep[:5] == "00000" {
			fmt.Println(hexRep, fmt.Sprintf("%c", hexRep[5]))
			pos, err := strconv.Atoi(fmt.Sprintf("%c", hexRep[5]))
			if err != nil {
				continue
			}
			if pos < 8 {
				_, present := spotTaken[pos]
				if !present {
					password[pos] = hexRep[6]
					fmt.Println(string(password))
					nFound++
					spotTaken[pos] = true
				}
			}
		}
	}
	fmt.Println(password)
}
