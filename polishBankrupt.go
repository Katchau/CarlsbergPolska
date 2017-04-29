package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fxsjy/gonn/gonn"
)

func importDataSet(filepath string) ([][]float64, [][]float64, [][]float64, [][]float64) {
	start := time.Now()

	f, _ := os.Open(filepath)
	defer f.Close()

	content, _ := ioutil.ReadAll(f)
	sContent := string(content)
	lines := strings.Split(sContent, "\n")

	inputs := make([][]float64, 0)
	targets := make([][]float64, 0)

	for _, line := range lines {

		line = strings.TrimRight(line, "\r\n")

		if len(line) == 0 {
			break
		}

		tup := strings.Split(line, ",")

		//tup = remove(tup, 61) // (short-term liabilities *365) / sales
		//tup = remove(tup, 54) // working capital
		//tup = remove(tup, 43) // (receivables * 365) / sales ( X61
		//	tup = remove(tup, 19) //(inventory * 365) / sales (X6
		//	tup = remove(tup, 18)
		//	tup = remove(tup, 17)

		in := tup[:len(tup)-1]
		out := tup[len(tup)-1]
		inValues := make([]float64, 0)

		for _, x := range in {
			fX, _ := strconv.ParseFloat(x, 64)
			inValues = append(inValues, fX)
		}

		inputs = append(inputs, inValues)
		outValue := make([]float64, 1)
		outParse, r := strconv.ParseFloat(out, 64)
		if r != nil {
			fmt.Print(r)
		}

		outValue[0] = outParse
		targets = append(targets, outValue)
	}

	trainInputs := make([][]float64, 0)
	testTargets := make([][]float64, 0)
	resultInputs := make([][]float64, 0)
	resultTargets := make([][]float64, 0)

	rangeValues := minMax(inputs)

	for i, x := range inputs {

		x = normalize(x, rangeValues)
		if i%3 == 0 {
			testTargets = append(testTargets, x)
			resultTargets = append(resultTargets, targets[i])
		} else {
			trainInputs = append(trainInputs, x)
			resultInputs = append(resultInputs, targets[i])
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Load DataSet took %s \n", elapsed)

	return trainInputs, resultInputs, testTargets, resultTargets
}

func normalize(in []float64, minMax [][]float64) []float64 {

	out := make([]float64, len(in))

	for i := range in {
		valueRange := minMax[i]
		value := in[i]

		out[i] = (value - valueRange[0]) / (valueRange[1] - valueRange[0])
		//	fmt.Printf("%.3f \n", out[i])
	}

	return out
}

func minMax(input [][]float64) [][]float64 {
	minMax := make([][]float64, len(input[0]))
	for i := range minMax {
		minMax[i] = make([]float64, 2)
		minMax[i][0] = 9999.0
		minMax[i][1] = -9999.0
	}
	for i := range input {
		line := input[i]

		for t := range line {
			if minMax[t][0] > line[t] {
				minMax[t][0] = line[t]
			}
			if minMax[t][1] < line[t] {
				minMax[t][1] = line[t]
			}
		}
	}
	return minMax
}

//NNBP create,train,test
func NNBP(trainInput [][]float64, trainTargets [][]float64, testInputs [][]float64, testTargets [][]float64) {

	start := time.Now()
	fmt.Printf("Size: %d \n", len(trainInput[0]))
	nn := gonn.NewNetwork(len(trainInput[0]), 500, 1, false, 0.25, 0.1)//TODO ver isto também

	nn.Train(trainInput, trainTargets, 140) //TODO ver isto

	gonn.DumpNN("1.nn", nn)

	nn = nil

	nn = gonn.LoadNN("1.nn")

	errCount := 0.0

	maxError := -1.0
	minError := 1.0

	good := 0.0
	for i := 0; i < len(testInputs); i++ {
		output := nn.Forward(testInputs[i])
		expect := testTargets[i][0]
		error := math.Abs(expect - output[0])

		fmt.Println(output[0], expect)

		if output[0] < 0.1 && expect == 0 {
			good++
		} else if output[0] > 0.8 && expect == 1 {
			good++
		} else if output[0] != expect {
			errCount++
			if maxError < error {
				maxError = error
			}
			if minError > error {
				minError = error
			}
		}

	}
	fmt.Printf("success rate: %.2f %% \n", (good / float64(len(testInputs)) * 100))
	fmt.Printf("error rate: %.2f %% \n", (errCount / float64(len(testInputs)) * 100))
	fmt.Printf("error range [%.4f , %.4f]\n", minError, maxError)
	elapsed := time.Since(start)
	fmt.Printf("Training and test took %s \n ", elapsed)

}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

//Dir  location of datasets
const Dir = "dataSet/"

//FileName default structure of Filename
const FileName = "yearV2.arff"

func main() {
	for i := 2; i <= 2; i++ {
		name := Dir + strconv.Itoa(i) + FileName
		fmt.Printf("\n" + name + "\n")
		t, tr, r, rt := importDataSet(name)
		fmt.Printf("\nGenerated %d %d Training sets and %d %d test sets \n", len(t), len(tr), len(r), len(rt))

		NNBP(t, tr, r, rt)
	}

}
