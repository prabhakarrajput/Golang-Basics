package main 
import "fmt" 

func describe(i interface{}){
	fmt.Printf("Type = %T, value = %v \n",i,i)
}

func main(){
	s := "Hello world"
	describe(s)
	i := 50
	describe(i)

	strt := struct{
		name string
	}{
		name : "Naveen R",
	}
	describe(strt)
}