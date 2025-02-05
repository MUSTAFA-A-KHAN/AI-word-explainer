package main

import (
	"bufio"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	adv, err := os.Open("../data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}

	advDF := dataframe.ReadCSV(adv)

	trainingNum := (4 * advDF.Nrow()) / 5
	testingNum := (advDF.Nrow()) / 5

	if trainingNum+testingNum < advDF.Nrow() {
		trainingNum++
	}

	trainingIdx := make([]int, trainingNum)
	testingIdx := make([]int, testingNum)

	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}
	for i := 0; i < testingNum; i++ {
		testingIdx[i] = i
	}

	trainingDF := advDF.Subset(trainingIdx)
	testingDF := advDF.Subset(testingIdx)

	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testingDF,
	}

	for idx, setName := range []string{"../data/split/training_adv.csv", "../data/split/testing_adv.csv"} {
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}
		w := bufio.NewWriter(f)

		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}

}
