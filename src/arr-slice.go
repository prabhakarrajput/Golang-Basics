package main
import "fmt"

func main() {
	//var students [3]string
	//fmt.Printf("Students : %v\n", students)
	// students[0] = "Ram"
	// students[1] = "Shyam"
	// fmt.Printf("Students : %v\n", students)

	// var identityMatrix [3][3]int  = [3][3]int{ [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,0} }
	// fmt.Printf("%v\n",identityMatrix)

	// a := []int{1,2,3,4,5,6,7,8,9}
	// b := a[2:6]
	// fmt.Println(b)

	//-----------------------slice and make function -----------------------------------------------------
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length : %v\n",len(a))
	fmt.Printf("Capacity : %v\n",cap(a))

	a = append(a,1)
	fmt.Printf("Length : %v\n",len(a))
	fmt.Printf("Capacity : %v\n",cap(a))

	a = append(a, []int{2,3,4,5}... )
	fmt.Printf("Length : %v\n",len(a))
	fmt.Printf("Capacity : %v\n",cap(a))

}