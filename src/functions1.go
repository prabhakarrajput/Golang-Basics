package main
import "fmt"

func main(){
	// greeting := "Hello"
	// name := "Ram"
	// sayMessage(&greeting,&name)
	// fmt.Println(name)

	// s := sum(1,2,3,4,5)
	// fmt.Println("The sum of values is : ",*s)

	x, err := divide(5.0,0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(x) 

}

func sayMessage(greeting, name *string){
	fmt.Println(*greeting,*name)
	*name = "Shyam"
}

//*************variatic functiion*******************************
func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _,v := range values {
		result += v
	}
	//fmt.Println("The sum of values is : ",result)
	return &result                 // return this variable as a pointer is the safe operation in go
								  //  as go handles the memory sharing of this local memory
}

func divide(a,b float64) (float64,error){
	if b == 0.0 {
		//panic("cannot divide by zero")
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil;
}