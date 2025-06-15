package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc (v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// can call value receiver with value
	fmt.Println(v.Abs())

	var p *Vertex
	p = &v
	// can call value receiver with pointer
	fmt.Println(p.Abs())

	// can call func with value argument with value 
	fmt.Println(AbsFunc(v))

	// can call func with value argument with value referenced by p
	fmt.Println(AbsFunc(*p))

	// cannot call func with value argument with pointer
	// fmt.Println(AbsFunc(&v))
}

