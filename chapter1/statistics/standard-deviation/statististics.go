package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func main() {
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// builds a DataFrame with the resulting records.
	irisDF := dataframe.ReadCSV(irisFile)

	//  Series with the petal_length column name contained in the DataFrame.
	sepalLength := irisDF.Col("petal_length").Float()

	// Min returns the minimum value in the sepalLength It panics if s is zero length
	minVal := floats.Min(sepalLength)

	// MAX returns the maximum value in the sepalLength It panics if s is zero length
	maxVal := floats.Max(sepalLength)

	rangeVal := maxVal - minVal

	//  the unbiased weighted sample variance:
	varianceVal := stat.Variance(sepalLength, nil)

	// e standard deviation.
	stdDevVal := stat.StdDev(sepalLength, nil)

	// Sort the values
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	fmt.Println("sepalLength\n", sepalLength, "\nvarianceVal", varianceVal, "\nstdDevVal", stdDevVal, "\ninds", inds, "\nrangeVal", rangeVal)

}
