package main

import (
	"cmp"
	"fmt"
	"slices"
)

/*
Sometimes we’ll want to sort a collection by something other than its natural order.
For example, suppose we wanted to sort strings by their length instead of alphabetically.
Here’s an example of custom sorts in Go
*/

func main() {
	fruits := []string{"apple", "banana", "kiwi"}
	slices.Sort(fruits)
	fmt.Println(fruits)

	lenComp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenComp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println(people)

}
