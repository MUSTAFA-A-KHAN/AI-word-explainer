package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {
	// Create a flat representaion ofour matrix
	matrix := []float64{1, 3, 5, 4, 7, 6}

	//form our matrix
	a := mat64.NewDense(3, 2, matrix)
	fa := mat64.Formatted(a)
	fmt.Println(fa)
	val := a.At(2, 1)
	fmt.Println("val At[1,1]", val)

	// get the values in a specific column
	col := mat64.Col(nil, 0, a)
	fmt.Println("column", col)

	//get the values in a specific Row
	row := mat64.Row(nil, 0, a)
	fmt.Println("row", row)

	//Modify element
	a.Set(0, 0, 2.2)

	//Modify an entire Column
	a.SetCol(0, []float64{3, 5, 4})
	fmt.Println(" ")
	fa = mat64.Formatted(a)
	fmt.Println(fa)

	//Modify an entire Column
	a.SetRow(0, []float64{3, 5})
	fmt.Println(" ")
	fa = mat64.Formatted(a)
	fmt.Println(fa)

}
