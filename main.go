package main

import "fmt"

func main() {

	inputs := []float32{1, 2, 3, 2.5}

	weights1 := []float32{0.2, 0.8, -0.5, 1.0}
	weights2 := []float32{0.5, -0.91, 0.26, -0.5}
	weights3 := []float32{-0.26, -0.27, 0.17, 0.87}

	var bias1 float32 = 2.0
	var bias2 float32 = 3.0
	var bias3 float32 = 0.5

	outputs := make([]float32, 3)

	outputs[0] = (inputs[0]*weights1[0] +
		inputs[1]*weights1[1] +
		inputs[2]*weights1[2] +
		inputs[3]*weights1[3] + bias1)

	outputs[1] = (inputs[0]*weights2[0] +
		inputs[1]*weights2[1] +
		inputs[2]*weights2[2] +
		inputs[3]*weights2[3] +
		bias2)

	outputs[2] = (inputs[0]*weights3[0] +
		inputs[1]*weights3[1] +
		inputs[2]*weights3[2] +
		inputs[3]*weights3[3] +
		bias3)
	fmt.Println(outputs)

}
