package main

import (
	"fmt"
	"sync"
	"time"
)

// How wait group works ?
// Lets see with this example

func process(i int, wg *sync.WaitGroup){
	fmt.Println("started goroutine ",i)
	time.Sleep(2*time.Second)
	fmt.Printf("goroutine %v ended \n",i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i:=0; i<no; i++ {
		wg.Add(1)
		go process(i,&wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished executing")
}