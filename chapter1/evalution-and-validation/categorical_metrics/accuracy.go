package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Accuracy: the percentage of data that were right
func main() {

	f, err := os.Open("./data/labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	var observed []int
	var predicted []int

	line := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if line == 1 {
			line++
			continue
		}

		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Error parsing line %d failed. Unexpected type", line)
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Error parsing line %d failed. Unexpected type", line)
			continue
		}

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)

	}
	var truePosNeg int

	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}
	accuracy := float64(float64(truePosNeg) / float64(len(observed)))

	fmt.Printf("Accuracy of the dataSet is %0.2f", accuracy)
}
