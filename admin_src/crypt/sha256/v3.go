package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "Hello world\n"
	sum := sha256.Sum256([]byte(s))
	fmt.Printf("%x", sum)
}
