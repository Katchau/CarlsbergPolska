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

//Z-SCORE = 1.2X1 + 1.4X2 + 3.3X3 + 0.6X4 + 1.0X5
//X1 = WorkingCapital/TotalAssets = X3 do outro
//X2 = RetainedEarning/TotalAssets = x6
//X3 = EBIT/TotalAssets = X7
//X4 = MarketValuesofEquity/BookValueofTotalLiabilities = x8 ?
//x5 = Sales / TotalAssets = x9
// Z < 1.1 Bancarrota :)
//afinal vamos usar este tb Z-SCORE = 0.717X1 +  0.847X2 + 3.107X3 + 0.420X4 + 0.998X5
//para este z > 2.9 safe zone, 1.23 < Z < 2.9 Grey zone, z < 1.23 distress zone

func getZscore(tup []string) []float64 {
	in := [5]string{tup[2], tup[5], tup[6], tup[7], tup[8]}
	inValues := make([]float64, 0)

	for _, x := range in {
		fX, _ := strconv.ParseFloat(x, 64)
		inValues = append(inValues, fX)
	}
	return inValues
}

//nao sabia que chamar a isto
func getInputAndOutput(tup []string) (bool, []float64, []float64) {
	in := tup[:len(tup)-1]
	out := tup[len(tup)-1]
	inValues := make([]float64, 0)
	addInput := true

	for _, x := range in {
		if x == "?" {
			addInput = false
			break
		}
		fX, _ := strconv.ParseFloat(x, 64)
		inValues = append(inValues, fX)
	}

	outValue := make([]float64, 1)
	outParse, r := strconv.ParseFloat(out, 64)
	if r != nil {
		fmt.Print(r)
	}

	outValue[0] = outParse
	return addInput, inValues, outValue
}

//apeteceu-me
//tup = remove(tup, 61) // (short-term liabilities *365) / sales
//tup = remove(tup, 54) // working capital
//tup = remove(tup, 43) // (receivables * 365) / sales ( X61
//	tup = remove(tup, 19) //(inventory * 365) / sales (X6
//	tup = remove(tup, 18)
//	tup = remove(tup, 17)

//esta funcao ta enorme pessoal
func importDataSet(filepath string, isZscore bool) ([][]float64, [][]float64, [][]float64, [][]float64) {
	start := time.Now()

	f, _ := os.Open(filepath)
	defer f.Close()

	content, _ := ioutil.ReadAll(f)
	sContent := string(content)
	lines := strings.Split(sContent, "\n")

	inputs := make([][]float64, 0)
	targets := make([][]float64, 0)

	var train_length int

	for index, line := range lines {

		if index == 0 {
			fmt.Printf("%s\n", line)
			train_length, _ = strconv.Atoi(line)
			continue
		}

		line = strings.TrimRight(line, "\r\n")

		if len(line) == 0 {
			break
		}
		tup := strings.Split(line, ",")

		if !isZscore {
			addInput, inValues, outValue := getInputAndOutput(tup)
			if addInput {
				inputs = append(inputs, inValues)
				targets = append(targets, outValue)
			}
		} else {
			inValues := getZscore(tup)
			out := tup[len(tup)-1]
			outValue := make([]float64, 1)
			outParse, _ := strconv.ParseFloat(out, 64)
			outValue[0] = outParse
			inputs = append(inputs, inValues)
			targets = append(targets, outValue)
		}

	}

	trainInputs := make([][]float64, 0)
	testTargets := make([][]float64, 0)
	resultInputs := make([][]float64, 0)
	resultTargets := make([][]float64, 0)

	rangeValues := minMax(inputs)

	for i, x := range inputs {
		x = normalize(x, rangeValues)

		if i < train_length {
			trainInputs = append(trainInputs, x)
			resultInputs = append(resultInputs, targets[i])
		} else {
			testTargets = append(testTargets, x)
			resultTargets = append(resultTargets, targets[i])
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
	nn := gonn.NewNetwork(len(trainInput[0]), 200, 1, false, 0.25, 0.1) //TODO ver isto também

	nn.Train(trainInput, trainTargets, 200) //TODO ver isto

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

func trainIndividualYear(year int, zscore bool) {
	name := Dir + strconv.Itoa(year) + FileName
	fmt.Printf("\n" + name + "\n")
	t, tr, r, rt := importDataSet(name, zscore)
	fmt.Printf("\nGenerated %d Training sets and %d test sets \n", len(t), len(r))

	NNBP(t, tr, r, rt)
}

func trainAllYearsIndividually(zscore bool) {
	for i := 1; i <= 5; i++ {
		trainIndividualYear(i, zscore)
	}
}

func appendArray(input [][][]float64) [][]float64 {
	i := make([][]float64, 0)
	for _, y := range input {
		for _, x := range y {
			i = append(i, x)
		}
	}
	return i
}

func trainAllYears(zscore bool) {
	input1 := make([][][]float64, 0)
	input2 := make([][][]float64, 0)
	input3 := make([][][]float64, 0)
	input4 := make([][][]float64, 0)
	for i := 1; i <= 5; i++ {
		name := Dir + strconv.Itoa(i) + FileName
		fmt.Printf("\n" + name + "\n")
		t, tr, r, rt := importDataSet(name, zscore)
		input1 = append(input1, t)
		input2 = append(input2, tr)
		input3 = append(input3, r)
		input4 = append(input4, rt)
		fmt.Printf("\nGenerated %d Training sets and %d test sets \n", len(t), len(r))
	}
	i1 := appendArray(input1)
	i2 := appendArray(input2)
	i3 := appendArray(input3)
	i4 := appendArray(input4)

	NNBP(i1, i2, i3, i4)
}

func main() {
	zscore := false
	//trainAllYears(zscore)
	trainIndividualYear(1, zscore)
	//trainAllYearsIndividually(zscore)
}
