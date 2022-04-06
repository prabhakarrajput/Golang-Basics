package main
import "fmt"

func producer(chnl chan int){
	for i:=1; i<=9; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)

	// for {
	// 	v,ok := <-ch
	// 	if ok==false{
	// 		break
	// 	}
	// 	fmt.Println("Reveived : ",v,ok)
	// }

	for v := range ch {
		fmt.Println("Received ",v)
	}
}