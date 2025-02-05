package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func main() {
	trainingAdv, err := os.Open("../../data/split/training_adv.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer trainingAdv.Close()
	reader := csv.NewReader(trainingAdv)
	// FieldsPerRecord is the number of expected fields per record.
	reader.FieldsPerRecord = 4
	trainingAdvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var regr regression.Regression
	regr.SetObserved("Sales")
	regr.SetVar(0, "TV")

	for _, record := range trainingAdvData {
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
	fmt.Println("Regressoin formula:\n", regr.Formula)

	testing_adv, err := os.Open("../../data/split/testing_adv.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer testing_adv.Close()
	reader = csv.NewReader(testing_adv)
	// FieldsPerRecord is the number of expected fields per record.
	reader.FieldsPerRecord = 4
	testingAdvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var mAE float64

	for _, record := range testingAdvData {
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		yPredicted, err := regr.Predict([]float64{tvVal})
		if err != nil {
			log.Fatal(err)
		}
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testingAdvData))
	}
	fmt.Println("MAE=", mAE)

}
