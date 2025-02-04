package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
	// golang.org/x/exp
	// This module is necessary because golang.org/x/exp/rand is imported in:

	// - github.com/MUSTAFA-A-KHAN/AI-word-explainer/chapter1/evalution-and-validation
	// -- gonum.org/v1/gonum/stat/distuv
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	observed := []float64{
		260.0,
		135.0,
		105.0,
	}

	totalObserved := 500.0

	expected := []float64{
		totalObserved * 0.60,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	chisquare := stat.ChiSquare(observed, expected)
	fmt.Println("\n Chisquare:", chisquare)

	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}
	pvalue := chiDist.Prob(chisquare)
	fmt.Println("\n P-value:", pvalue)
}
