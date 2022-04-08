package main
import (
	"fmt"
	"sync"
)
var x = 0
func increment(wg *sync.WaitGroup){
	x+=1
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	for i:=0 ; i<1000 ; i++ {
		wg.Add(1)
		go increment(&wg) 
	}
	wg.Wait()
	fmt.Println("Final value of x is ",x)
}