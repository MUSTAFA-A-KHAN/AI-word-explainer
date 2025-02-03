package main

import (
	"fmt"
	"log"

	"github.com/gonum/matrix/mat64"
)

func main() {
	// Create a flat representaion ofour matrix

	//form our matrix
	a := mat64.NewDense(4, 4, []float64{
		1, 3, 5, 4,
		7, 6, 7, 8,
		9, 0, 12, 3,
		4, 2, 1, 3})
	fa := mat64.Formatted(a, mat64.Prefix("   "))
	fmt.Println("a=", fa)

	// Calculate the determinant
	det := mat64.Det(a)
	fmt.Println("Determinant of a:", det)

	//Determine the Inverse
	aInverse := mat64.NewDense(0, 0, nil)
	err := aInverse.Inverse(a)
	if err != nil {
		fmt.Println("err", err)
	}
	fIa := mat64.Formatted(aInverse, mat64.Prefix("     "))
	fmt.Println("fIa=", fIa)
	log.Fatal(err)

}
