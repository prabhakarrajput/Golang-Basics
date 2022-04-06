package main
import "fmt"

// method is just a function which is getting executed in known context

func main(){
	g := greets {
		greeting : "hello",
		name : "go",
	}
	g.greet()
	fmt.Println("the new name is ",g.name)
}

type greets struct {
	greeting string
	name string
}

func (g *greets) greet() {
	fmt.Println(g.greeting, g.name)
	g.name="ram"
}