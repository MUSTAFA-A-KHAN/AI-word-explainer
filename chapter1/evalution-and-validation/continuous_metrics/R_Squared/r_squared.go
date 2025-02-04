package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

// Rsqaured hold the results of a regression analysis
// R-squared measures the proportion of the variance in the observed values of the variance that we capture in the predicted values.
// R-squared measures the proportion of the variance in the dependent variable that is predictable from the independent variable(s), relative to the total variance.
func main() {
	f, err := os.Open("../../data/continuous_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	var observed []float64
	var predicted []float64

	line := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		//skip th first headers
		if line == 1 {
			line++
			continue
		}

		observedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed. Unexpected type", line)
			continue
		}
		predictedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed. Unexpected type", line)
			continue
		}
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++

	}
	rSquared := stat.RSquaredFrom(observed, predicted, nil)
	fmt.Printf("rSquared(R^2)=%0.2f", rSquared)
}
