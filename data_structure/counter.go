package main

import "fmt"

type Counter map[interface{}]int

func (c *Counter) Add(key interface{}) {
	(*c)[key] += 1
}

func (c *Counter) Find(key interface{}) bool {
	_, ok := (*c)[key]
	return ok
}

func (c *Counter) Get(key interface{}) (int, bool) {
	val, ok := (*c)[key]
	return val, ok
}

func main() {
	mp := make(Counter)
	mp.Add("a")
	mp.Add("b")
	mp.Add("c")

	fmt.Println(mp.Find("a"))
	fmt.Println(mp.Find("b"))

	fmt.Println(mp.Get("c"))
	fmt.Println(mp.Get("d"))

}
