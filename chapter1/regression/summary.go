package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {

	adv, err := os.Open("./data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer adv.Close()

	advDF := dataframe.ReadCSV(adv)
	advSummary := advDF.Describe()

	fmt.Println("Advertising Summary\n", advSummary)
}
