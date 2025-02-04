package main

import (
	"fmt"

	"github.com/gonum/stat"
	"gonum.org/v1/gonum/integrate"
)

func main() {
	scores := []float64{0.1, 0.35, 0.4, 0.8}
	classes := []bool{true, false, true, false}

	tpr, fpr := stat.ROC(0, scores, classes, nil)

	auc := integrate.Trapezoidal(fpr, tpr)

	fmt.Println("tpr :", tpr)
	fmt.Println("fpr :", fpr)
	fmt.Println("auc :", auc)
}
