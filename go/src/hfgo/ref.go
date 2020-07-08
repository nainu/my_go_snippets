package main

import "fmt"

func myFloatPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func main() {
	myFloatPointer := myFloatPointer()
	fmt.Println(*myFloatPointer)
}
