// Producer will publish for 2 sec and then sleep for 5 sec, consumer is consuming  continuously.

package main
import (
	"fmt"
	"time"
)

func producer(job chan<- int,done chan bool){
	i := 1
	for {
		loop :
		for timeout := time.After(2*time.Second);;{
			select {
			case <-timeout :
				time.Sleep(5*time.Second)
				break loop
			default :
			}
			i++
			job<-i
			fmt.Println("Produced ",i)
			if(i>=1000000){
				break
			}
		}
		if(i>=1000000){
			break
		}
	}
	done <- true
}

func consumer(job <-chan int){
	for {
		fmt.Println("Consumed ",<-job)
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