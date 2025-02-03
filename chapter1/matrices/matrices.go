package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {
	data := []float64{1.3, -5.7, -2.5, .3}
	fmt.Println("Original data: ", data)

	// NewDense creates a new Dense matrix with r rows and c columns.
	// If data == nil, a new slice is allocated for the backing slice.
	// If len(data) == r*c, data is used as the backing slice, and changes
	// to the elements of the returned Dense will be reflected in data.
	// If neither of these is true, NewDense will panic.

	// The data must be arranged in row-major order, i.e. the (i*c + j)-th element in the data slice is the {i, j}-th element in the matrix.
	a := mat64.NewDense(2, 2, data)
	fmt.Println(a)           //will print the unformatted matrices from
	fa := mat64.Formatted(a) //will format the matrices
	fmt.Println(fa)
}
