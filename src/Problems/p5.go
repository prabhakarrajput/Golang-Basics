// Consumer will consume for 2 sec and then sleep for 5 sec, producer is producing continuously

package main
import (
	"fmt"
	"time"
)

func producer(job chan<- int,done chan bool){
	i := 1
	for {
		fmt.Println("Produced ",i)
		job <- i
		i++
		if(i>1000000){
			break
		}
	}
	done <- true
}

func consumer(job <-chan int){
	for {
		loop :
		for timeout := time.After(2*time.Second);;{
			select {
			case <-timeout :
				fmt.Println("Consumer is in sleep mode for 5 seconds")
				time.Sleep(5*time.Second)
				break loop
			default :
			}
			fmt.Println("Consumed ",<-job)
		}
	}
}

func main(){
	job := make(chan int)
	done := make(chan bool)

	go producer(job,done)
	go consumer(job)
	msg := <-done
	fmt.Println(msg)

}