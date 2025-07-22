package main

import (
	"fmt"
	"hash/maphash"
)

func main() {
	var h maphash.Hash

	h.WriteString("hello, ")
	fmt.Println(h.Sum64())
	h.WriteString("rashed")
}
