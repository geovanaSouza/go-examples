package main

import "fmt"

func main() {

	// Force error: Creating variable without define type. Golang can't infer the type
	// Error: use of untyped nil
	// a := nil
	// fmt.Println(a)

	var s *string = nil
	fmt.Println(s)

	// Force error: Compare different types of zero-init. Golang can't compare
	// Error: mismatched types float64 and int
	// var f float64
	// var i int
	// fmt.Println(f == i)

	var f float64
	var i int
	fmt.Println(int(f) == i)

}
