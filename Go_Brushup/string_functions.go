package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	p("Contains", strings.Contains("test", "es"))
	p("HasPrefix", strings.HasPrefix("test", "es"))
	p("Count", strings.Count("test", "es"))
	p("Index", strings.Index("test", "st"))
	p("Join", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat", strings.Repeat("a", 5))
	p("Split", strings.Split("a-b-c-d-e", "-"))
	p("toUpper", strings.ToUpper("test"))
	p("toLower", strings.ToLower("test"))

}
