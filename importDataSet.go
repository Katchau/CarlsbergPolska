package iart

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/sbinet/go-arff"
)

//Result : stores all entrances
type Result struct {
	count   int
	Objects [10503]Entrance
}

//Entrance : define a line of the dataset
type Entrance struct {
	Atr1  float64 `arff:"Attr1"`
	Atr2  float64 `arff:"Attr2"`
	Atr3  float64 `arff:"Attr3"`
	Atr4  float64 `arff:"Attr4"`
	Atr5  float64 `arff:"Attr5"`
	Atr6  float64 `arff:"Attr6"`
	Atr7  float64 `arff:"Attr7"`
	Atr8  float64 `arff:"Attr8"`
	Atr9  float64 `arff:"Attr9"`
	Atr10 float64 `arff:"Attr10"`
	Atr11 float64 `arff:"Attr11"`
	Atr12 float64 `arff:"Attr12"`
	Atr13 float64 `arff:"Attr13"`
	Atr14 float64 `arff:"Attr14"`
	Atr15 float64 `arff:"Attr15"`
	Atr16 float64 `arff:"Attr16"`
	Atr17 float64 `arff:"Attr17"`
	Atr18 float64 `arff:"Attr18"`
	Atr19 float64 `arff:"Attr19"`
	Atr20 float64 `arff:"Attr20"`
	Atr21 float64 `arff:"Attr21"`
	Atr22 float64 `arff:"Attr22"`
	Atr23 float64 `arff:"Attr23"`
	Atr24 float64 `arff:"Attr24"`
	Atr25 float64 `arff:"Attr25"`
	Atr26 float64 `arff:"Attr26"`
	Atr27 float64 `arff:"Attr27"`
	Atr28 float64 `arff:"Attr28"`
	Atr29 float64 `arff:"Attr29"`
	Atr30 float64 `arff:"Attr30"`
	Atr31 float64 `arff:"Attr31"`
	Atr32 float64 `arff:"Attr32"`
	Atr33 float64 `arff:"Attr33"`
	Atr34 float64 `arff:"Attr34"`
	Atr35 float64 `arff:"Attr35"`
	Atr36 float64 `arff:"Attr36"`
	Atr37 float64 `arff:"Attr37"`
	Atr38 float64 `arff:"Attr38"`
	Atr39 float64 `arff:"Attr39"`
	Atr40 float64 `arff:"Attr40"`
	Atr41 float64 `arff:"Attr41"`
	Atr42 float64 `arff:"Attr42"`
	Atr43 float64 `arff:"Attr43"`
	Atr44 float64 `arff:"Attr44"`
	Atr45 float64 `arff:"Attr45"`
	Atr46 float64 `arff:"Attr46"`
	Atr47 float64 `arff:"Attr47"`
	Atr48 float64 `arff:"Attr48"`
	Atr49 float64 `arff:"Attr49"`
	Atr50 float64 `arff:"Attr50"`
	Atr51 float64 `arff:"Attr51"`
	Atr52 float64 `arff:"Attr52"`
	Atr53 float64 `arff:"Attr53"`
	Atr54 float64 `arff:"Attr54"`
	Atr55 float64 `arff:"Attr55"`
	Atr56 float64 `arff:"Attr56"`
	Atr57 float64 `arff:"Attr57"`
	Atr58 float64 `arff:"Attr58"`
	Atr59 float64 `arff:"Attr59"`
	Atr60 float64 `arff:"Attr60"`
	Atr61 float64 `arff:"Attr61"`
	Atr62 float64 `arff:"Attr62"`
	Atr63 float64 `arff:"Attr63"`
	Atr64 float64 `arff:"Attr64"`
	Class float64 `arff:"class"`
}

func importDataSet(filepath string) Result {

	start := time.Now()

	f, err := os.Open(filepath)

	dec, err := arff.NewDecoder(f)

	var r Result

	r.count = 0

	for {
		var v Entrance
		err = dec.Decode(&v)
		if err == io.EOF {
			break
		}

		r.Objects[r.count] = v
		r.count++
	}

	f.Close()

	elapsed := time.Since(start)
	fmt.Printf("Load DataSet took %s", elapsed)

	return r
}

func generateSets(filepath string) ([][]float64, [][]float64, [][]float64, [][]float64) {
	train_inputs := make([][]float64, 0)
	train_targets := make([][]float64, 0)

	test_inputs := make([][]float64, 0)
	test_targets := make([][]float64, 0)

	r := importDataSet(filepath)

	testSize := r.count / 3
	trainSize := testSize * 2

	currentTestSize := 0
	currentTrainSize := 0

	for _, element := range r.Objects {
		random := rand.Intn(99999)

		if random%3 == 0 {
			//ADD to test set
		} else {
			//ADD to train set
		}

	}

	return train_inputs, train_targets, test_inputs, test_targets
}
