// Single producer - single consumer

package main
import (
	"fmt"
	"time"
)

var messages = []string{
	"1. Ram",
	"2. Raghu",
	"3. Kavi",
	"4. Ryan",
	"5. Adarsh",
	"6. Honey",
	"7. Gaurav",
	"8. Lucky",
}

func producer(link chan<- string) {
	i := 1
	for _,m := range messages{
		fmt.Println("Message no ",i," sent")
		link <- m
		i++
		time.Sleep(2*time.Second)
	}
	close(link)
}

func consumer(link <-chan string, done chan<- bool) {
	for b := range link {
		fmt.Println("Message Received ",b)
	}
	done <- true
}

func main() {
	link := make(chan string, 4)
	done := make(chan bool)
	go producer(link)
	go consumer(link,done)
	<-done
}



