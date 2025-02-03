package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func main() {
	VectorA := []float64{2, 55, -1.3}
	VectorB := []float64{1, 2, 3}

	// dotProduct := mat64.Dot(VectorA, VectorB)// VectorA and VectorB must be of type matrics

	dotProduct := floats.Dot(VectorA, VectorB)
	fmt.Println("Vector Dot product:", dotProduct)

	//Scale multiplies every element in dst by the scalar c.
	floats.Scale(1.5, VectorA)
	fmt.Println("VectorA:", VectorA)

	//comput the norm or length of B
	normB := floats.Norm(VectorB, 2)
	fmt.Println("the norm or length of B", normB)
}
