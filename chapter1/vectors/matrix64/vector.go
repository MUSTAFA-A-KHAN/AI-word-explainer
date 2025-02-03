package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {
	myVector := mat64.NewVector(2, []float64{11.0, 6.2})
	fmt.Println("myVector", myVector)
}
