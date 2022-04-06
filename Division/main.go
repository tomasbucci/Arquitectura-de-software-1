package div

package main


import (
	"fmt"
	"github.com/tomasbucci/Arquitectura-de-software-1/tree/ejercicio-http/Division/div"
)

func main() {
	defer func() {
		fmt.Println("Bye!")
	}()

	a, b := float64(20), float64(10)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	resultDiv, err := div.Division(a, b)
	if err != nil {
		fmt.Println("Error in division: ", err.Error())
		return
	}
	fmt.Println("Div: ", resultDiv)
	
}