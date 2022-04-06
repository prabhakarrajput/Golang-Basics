package main
import "fmt"

func main() {
	// var a int = 40
	// var b *int = &a
	// fmt.Println(a,*b)
	// *b = 50
	// fmt.Println(a,*b)

	// a := [3]int{1,2,3}
	// b := &a[0]
	// c := &a[1]
	// fmt.Printf("%v %p %p\n",a,b,c)

	var ms *myStruct
	//ms = &myStruct{foo:40}
	ms = new(myStruct)
	// (*ms).foo = 42
	// fmt.Println((*ms).foo)

	ms.foo = 50
	fmt.Println(ms.foo)

}

type myStruct struct {
	foo int
}