package main

import "fmt"

func main() {
	var myVector []float64
	myVector = append(myVector, 1.0, 2.0, 3.9)
	myVector = append(myVector, 9.0, 8.0)
	fmt.Println("myVector", myVector)
}
