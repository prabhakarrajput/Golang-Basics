package main

import (
	"fmt"
	"strconv"
)

// var i int = 10
func main() {
	// fmt.Printf("%v %T\n", i, i)
	// var i float32 = 20.5
	// fmt.Printf("%v %T", i, i)

	var i int = 20
	fmt.Printf("%v %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v %T\n", j, j)
}
