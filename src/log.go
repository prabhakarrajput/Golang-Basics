package main
import (
	"log"
)

func main(){
	// log.Fatal("Exception occured")  //similar to print statement but it has exit status 1

	// s:="formatted"
	// log.Fatalf("A %s string for logging",s)

	// log.Fatalln("A logged string with new line")

	log.Panicln("A panic with new line")
}