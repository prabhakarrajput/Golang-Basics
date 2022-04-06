package main
import "fmt"

func main(){
	//var r rune = 'a'
	// r := 'b'
	// fmt.Printf("%v %T",r,r)

	s := "Hello World"
	r := []rune(s)
	fmt.Printf("%v %T",r,r)
}