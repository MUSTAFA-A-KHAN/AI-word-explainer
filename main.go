package main

import (
	"fmt"

	t "gorgonia.org/tensor"
)

func main() {

	rawWeights := []float32{
		0.2, 0.8, -0.5, 1.0,
		0.5, -0.91, 0.26, -0.5,
		-0.26, -0.27, 0.17, 0.87,
	}

	inputs := t.New(t.WithBacking([]float32{1, 2, 3, 2.5}))
	weights := t.New(t.WithShape(3, 4), t.WithBacking(rawWeights))
	biase := t.New(t.WithBacking([]float32{2.0}))

	dotproduct, _ := t.Dot(weights, inputs)
	output, _ := t.Add(dotproduct, biase)
	fmt.Println(output)

}
