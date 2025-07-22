package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) string() {
	fmt.Printf("%v, %v\n", v.X, v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(5)
	fmt.Println(v)
}
