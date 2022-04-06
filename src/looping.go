package main
import "fmt"

func main() {
	// for i :=0; i<5 ; i++ {
	// 	fmt.Println(i)
	// }

	s := []int{1,2,3}
	for k,v := range s{
		fmt.Println(k,v)
	}
}