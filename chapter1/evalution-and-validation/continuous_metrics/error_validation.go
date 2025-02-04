package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../data/continuous_data.csv")
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
	// error is the diference between the observed value and the predicted value
	// the average of the absolute value of all the errors
	var mAE float64

	// // the average of squares all the errors
	var mSE float64

	for idx, obsVal := range observed {
		mAE += math.Abs(obsVal-predicted[idx]) / float64(len(observed))
		mSE += math.Pow(obsVal-predicted[idx], 2) / float64(len(observed))

	}
	fmt.Println("Mean Absolute Error:", mAE)
	fmt.Println("Mean Squared Error:", mSE)
}
