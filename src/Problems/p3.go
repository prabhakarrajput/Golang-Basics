// multiple producer - single consumer

package main
import (
	"fmt"
	"sync"
)

var messages = []string{
	"1. Ram",
	"2. Raghu",
	"3. Kavi",
	"4. Ryan",
	// "5. Adarsh",
	// "6. Honey",
	// "7. Gaurav",
	// "8. Lucky",
}

func produce(p int, channel chan<- string, wg *sync.WaitGroup){
	defer wg.Done()
	for _,v := range messages {
		fmt.Printf("Producer %v is sending the message %v \n",p, v)
		channel <- v
	}
}

func consume(channel <-chan string, done chan<- bool) {
	for v := range channel {
		fmt.Printf("Consumed message %v \n", v)
	}
	done <- true
}

func main() {
	channel := make(chan string, 4)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	// lets suppose there are 3 producres
	for i:=1 ; i<=3 ; i++ {
		wg.Add(1)
		go produce(i,channel,&wg)
	}
	go consume(channel,done)
	wg.Wait()
	close(channel)
	<-done
}