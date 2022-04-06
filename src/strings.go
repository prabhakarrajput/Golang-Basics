package main
import "fmt"

func main() {
	s := "this is a string"
	//s[2] = u
	// fmt.Printf("%v %T",s,s)
	
	b := []byte(s)
	fmt.Printf("%v %T",b,b)
}