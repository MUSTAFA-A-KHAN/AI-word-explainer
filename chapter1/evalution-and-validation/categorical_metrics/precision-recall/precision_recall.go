package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../data/labeled.csv")
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
			log.Printf("error parsin value at line %d: %v.Unexpected type", line, err)
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("error parsin value at line %d: %v.Unexpected type", line, err)
			continue
		}
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++

	}

	classes := []int{0, 1, 2}
	var truePos int
	var falsePos int
	var falseNeg int
	for _, class := range classes {

		for idx, oVal := range observed {
			switch oVal {
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				}
				falseNeg++
			default:
				if predicted[idx] == class {
					falsePos++
				}
			}
		}
		precision := float64(truePos) / float64(truePos+falseNeg)
		recall := float64(truePos) / float64(truePos+falseNeg)

		fmt.Printf("\nRecall:=(class%d)=%0.2f", class, recall)
		fmt.Printf("\nprecsion:=class%d)=%0.2f", class, precision)

	}

}
