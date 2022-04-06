package main
import "fmt"

func calSquare(no int, squareOp chan int){
	sum :=0
	for no != 0 {
		digit := no%10
		sum += digit*digit
		no /= 10
	}
	squareOp <- sum
}

func calCube(no int, cubeOp chan int){
	sum :=0
	for no != 0 {
		digit := no%10
		sum += digit*digit*digit
		no /= 10
	}
	cubeOp <- sum
}

func main() {
	squareCh := make(chan int)
	cubeCh := make(chan int)

	number := 589

	go calSquare(number,squareCh)
	go calCube(number,cubeCh)

	squares, cubes := <-squareCh, <-cubeCh
	fmt.Println("Final output : ", squares+cubes )

}
