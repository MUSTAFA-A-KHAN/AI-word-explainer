package main

import (
	"fmt"

	"github.com/gonum/blas/blas64"
	"github.com/gonum/matrix/mat64"
)

func main() {
	VectorA := mat64.NewVector(3, []float64{2, 55, -1.3})
	VectorB := mat64.NewVector(3, []float64{1, 2, 3})

	dotProduct := mat64.Dot(VectorA, VectorB) // VectorA and VectorB must be of type matrics

	// dotProduct := floats.Dot(VectorA, VectorB)
	fmt.Println("Vector Dot product:", dotProduct)

	//Scale multiplies every element in dst by the scalar c.
	VectorA.ScaleVec(1.5, VectorA)
	fmt.Println("VectorA:", VectorA)

	//comput the norm or length of B
	normB := blas64.Nrm2(3, VectorB.RawVector())
	fmt.Println("the norm or length of B", normB)
}
