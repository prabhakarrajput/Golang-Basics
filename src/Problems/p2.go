// single producer - multiple consumers

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

func consumer(c int,link <-chan string, done chan<- bool) {
	for b := range link {
		fmt.Printf("Message received by the consumer %v is : %v \n",c,b)
	}
	done <- true
}

func main() {
	link := make(chan string, 4)
	done := make(chan bool)
	go producer(link)
	for i:=1 ; i<=3; i++ {
		go consumer(i,link,done)
	}
	<-done
}