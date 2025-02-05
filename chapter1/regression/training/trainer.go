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
	trainig_adv, err := os.Open("../data/split/training_adv.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a csv reader and read the csv file
	reader := csv.NewReader(trainig_adv)

	trainig_advData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var regr regression.Regression
	regr.SetObserved("Sales")
	regr.SetVar(0, "TV")

	for _, record := range trainig_advData[1:] {
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		regr.Train(regression.DataPoint(yVal, []float64{tvVal}))

	}
	regr.Run()
	fmt.Printf("\n Regression Formula:\n%v\n\n", regr.Formula)
}
