package main
import "fmt"

func main() {
	// var a bool = true;
	// a := 1==2
	// fmt.Printf("%v %T",a,a)

	// var n uint16 = 15
	// fmt.Printf("%v %T",n,n)

	// var x int = 2
	// var y int8 = 4
	// fmt.Println(x+int(y))

	// a := 10   // 1010
	// b := 3    // 0011

	// fmt.Println(a & b)    //0010  -> 2
	// fmt.Println(a | b)	  //1011  -> 11
	// fmt.Println(a ^ b)	  //1001  -> 9
	// fmt.Println(a &^ b)   //0100  ->8

	// ------------Bit shift---------------------
	// a := 25               
	// fmt.Println(a << 5)  // 25 * 2^5                left shift
	// fmt.Println(a >> 5)  // 25 / 2^5                right shift

	//------------complex number -----------------------
	var a complex64 = 2i
	fmt.Printf("%v %T",a,a)

}