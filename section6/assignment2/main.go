package main

import "fmt"

type triangle struct {
    height float64
    base   float64
}

type square struct {
    sideLength float64
}

type shape interface {
    getArea() float64
}

func main() {
    fmt.Println("Assignment 2")
    t := triangle{10, 10}
    s := square{10}

    printArea(s)
    printArea(t)
}

func (s square) getArea() float64 {
    return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
    return t.base * t.height / 2
}

func printArea(s shape) {
    fmt.Println(s.getArea())
}
