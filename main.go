package main

import (
	"fmt"
	"golangunittestexample/calculation"
)


func main() {
	var a, b int
	fmt.Println("Input your first number : ")
    fmt.Scan(&a)
	fmt.Println("Input your second number : ")
	fmt.Scan(&b)
	fmt.Println("Calculate result : ", calculation.Add(a,b))
}