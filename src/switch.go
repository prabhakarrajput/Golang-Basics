package main
import "fmt"

func main(){
	// i := 3
	// switch i {
	// case 1 :
	// 	fmt.Println("One")
	// case 2 :
	// 	fmt.Println("Two")
	// default : 
	// 	fmt.Println("not one or two")
	// }

	var i interface{} = 10.2
	switch i.(type){
	case int :
		fmt.Println("integer")
	case float64 :
		fmt.Println("float64")
	default :
		fmt.Println("another type")
	}
}