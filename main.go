package main

import (
	"fmt"
	"go-calculator-exercice/utils"
)


func main() {
	
	fmt.Println("La somme est ", utils.Sum("4 ","10"))
	fmt.Println("Le produit est ", utils.Multiplication("10", "5"))
}