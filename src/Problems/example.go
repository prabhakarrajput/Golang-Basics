package main
import (
	"fmt"
	//"time"
)

func main(){
	outer:
	for i:=0 ; i<3; i++ {
		for j:=1; j<4; j++ {
			fmt.Printf("i = %v , j = %v \n",i,j)
			if i==1{
				break outer
			}
		}
	}
}