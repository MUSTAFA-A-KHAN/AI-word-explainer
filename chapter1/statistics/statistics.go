package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
)

func main() {
	irisFile, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)
	fmt.Println(irisDF)

	sepalLength := irisDF.Col("petal_length").Float()

	//Calculate the mean of the sepal length column
	meanValue := stat.Mean(sepalLength, nil)

	//Calculate the mode
	modeVal, count := stat.Mode(sepalLength, nil)

	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Mean:", meanValue)
	fmt.Println("Median:", medianVal)
	fmt.Println("modeVal:", modeVal, "count", count)
}
