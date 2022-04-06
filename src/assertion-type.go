package main 
import "fmt"

func assert(i interface{}) {
	s,ok := i.(string) //get the underlying int value from i
	fmt.Println(s,ok)
}
func main(){
	var s interface{} = "Ram"
	assert(s)
	var i interface{} = 50
	assert(i)
}