package main

import (
	"fmt"

	t "gorgonia.org/tensor"
)

func main() {
	inputs := t.New(t.WithBacking([]float32{1, 2, 3, 2.5}))
	weights := t.New(t.WithBacking([]float32{0.2, 0.8, -0.5, 1.0}))
	biase := t.New(t.WithBacking([]float32{2.0}))

	dotproduct, _ := t.Dot(weights, inputs)
	output, _ := t.Add(dotproduct, biase)
	fmt.Println(output)

}
