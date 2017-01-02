package main

import (
	"fmt"
	"strings"
)

func decryptPart(part string, sectorId int) (decrypted string) {
	for _, r := range part {
		prime := int(r) + sectorId%26
		if prime > 122 {
			prime -= 26
		}
		decrypted += fmt.Sprintf("%c", prime)
	}
	return
}

func decryptRoom(room string) (decrypted string, sectorId int) {
	parts := strings.Split(room, "-")

	var lock string
	part := parts[len(parts)-1]
	fmt.Sscanf(strings.Trim(part, "]"), "%d[%s]", &sectorId, &lock)

	for _, part := range parts[:len(parts)-1] {
		decrypted += decryptPart(part, sectorId)
		decrypted += " "
	}
	return
}

func main() {
	var rooms []string
	var r string

	n, _ := fmt.Scan(&r)
	for n > 0 {
		rooms = append(rooms, r)
		n, _ = fmt.Scan(&r)
	}

	for _, room := range rooms {
		fmt.Println(decryptRoom(room))
	}
}
