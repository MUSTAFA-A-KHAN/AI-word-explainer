package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func main() {

	adv, err := os.Open("../../data/Advertising.csv")
	if err != nil {
		log.Fatal(adv)
	}

	reader := csv.NewReader(adv)
	reader.FieldsPerRecord = 4

	trainingData, err := reader.ReadAll() // Read all records from the CSV file
	if err != nil {
		log.Fatal(err)
	}

	var regr regression.Regression

	regr.SetObserved("Sales")
	regr.SetVar(0, "TV")
	regr.SetVar(1, "Radio")

	for _, record := range trainingData[1:] {
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		regr.Train(regression.DataPoint(yVal, []float64{tvVal, radioVal}))
	}
	regr.Run()
	fmt.Println("Regression Formula:\n", regr.Formula)

}
