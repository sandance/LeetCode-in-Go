package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "Sha this is string"
	fmt.Println(s)

	h := sha256.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)
	fmt.Println(bs)
}
