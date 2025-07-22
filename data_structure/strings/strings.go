package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string = "This is an example of string"
	fmt.Printf("\n%t\n\n", strings.HasPrefix(str, "This is "))

	fmt.Printf("%t\n\n", strings.HasSuffix(str, "ting"))
	fmt.Printf("%v", strings.Index(str, "example"))
	fmt.Printf("%v", strings.LastIndex(str, "example"))
	fmt.Printf("%v", strings.Count(str, "T"))

	fmt.Printf("The size of ints in: %d\n", strconv.IntSize)
	an, _ := strconv.Atoi("20")
	fmt.Printf("The size of ints in: %d\n", an)

	an = an + 5
	newS := strconv.Itoa(an)
	fmt.Printf("The size of ints in: %s\n", newS)

}
