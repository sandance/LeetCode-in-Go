package main

import (
	"bytes"
	"fmt"
)

func main() {

	var b bytes.Buffer
	b.WriteString("Hello World")
	b.WriteString("ABC")

	fmt.Fprintf(&b, "A number %d", 123)
	b.WriteString("[DONE]")
	fmt.Println(b.String())
}
