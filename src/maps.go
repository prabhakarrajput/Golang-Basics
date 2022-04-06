package main
import "fmt"

func main() {
	statePopulations := map[string]int{
		"MP" : 122011000,
		"UP" : 15454992356,
		"AP" : 64662599,
	}
	fmt.Println(statePopulations)
	statePopulations["JK"] = 1246523
	fmt.Println(statePopulations)

	delete(statePopulations,"AP")
	fmt.Println(statePopulations)


}