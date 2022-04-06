// // producer is publishing 3 items per sec

package main

import (
	"fmt"
	"time"
)
func produce(pipe chan<- int, done chan bool){
	i:=1
	ticker := time.NewTicker(time.Second)
	for timer := range ticker.C{
		j:=i
		fmt.Println("Produced at ", timer)
		for ; j<=i+2 ; j++ {
			fmt.Println("Writes ",j)
			pipe <- j
		}
		i=j
		if(i>30){
			break
		}
	}
	time.Sleep(200*time.Millisecond)
	done <- true
}

func consume(pipe <-chan int){
	for {
		fmt.Println("Consumed ",<-pipe)
	}
}


func main() {
	pipe := make(chan int)
	done := make(chan bool)
	go produce(pipe,done)
	go consume(pipe)
	<-done
}