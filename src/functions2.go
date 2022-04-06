package main
import "fmt"

//anonymous functions
func main(){
	// func() {
	// 	fmt.Println("Hello its anonymous")
	// }()

	for i:=1; i<5; i++ {
		func(j int){
			fmt.Println(j)
		}(i)
	}
}