package main

import "fmt"

func main() {
	var b bool
	b = true
	b = false
	//b = true || false //or 
	b = true && false // and
	fmt.Println(b)
	//output = "true"
	x := true
	y := false
	z := x || y
	fmt.Println(z)
}
