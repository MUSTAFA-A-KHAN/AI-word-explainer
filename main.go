package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {
	//A chi-square table is a reference table that lists critical values for chi-square hypothesis tests.
	//Most of the time these will be dynamic
	observed := []float64{20, 18}
	expected := []float64{20, 13}

	// 	ChiSquare computes the chi-square distance between the observed frequencies 'obs' and expected frequencies 'exp' given by:

	// \sum_i (obs_i-exp_i)^2 / exp_i
	// The lengths of obs and exp must be equal.
	chiSquare := stat.ChiSquare(observed, expected)
	fmt.Println("Chisquare:", chiSquare)
}
