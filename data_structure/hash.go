package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Apple"] = 40
	m["Orange"] = 50
	m["Banana"] = 100
	m["Pear"] = 150

	for key, val := range m {
		fmt.Printf("[ ", key, " -> ", val, " ]\n")
	}
	fmt.Println()

	v, ok := m["Apple"]
	if ok {
		fmt.Println("Apple available at price", v)
	} else {
		fmt.Println("Apple not available at price")
	}

	delete(m, "Apple")
	v, ok = m["Apple"]
	if ok {
		fmt.Println("Apple available at price", v)
	} else {
		fmt.Println("Apple not available at price")
	}
}
