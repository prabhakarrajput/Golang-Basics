package main
import "fmt"

type Employee struct {
	name string
	phone int
	companions []string
}
func main() {
	emp1 := Employee{
		name : "Ram",
		phone : 9458315531,
		companions: []string{
			"Varun", "Riya",
		},
	}
	fmt.Printf("%v %T\n",emp1,emp1)
	fmt.Printf("%v %T\n",emp1.name,emp1.name)
}