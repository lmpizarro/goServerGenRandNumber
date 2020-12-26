package randgen

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	ut "github.com/lmpizarro/statproject/utils"
	"github.com/montanaflynn/stats"
)

// StatSample ...
type StatSample struct {
	Count int
	Value float64
}

var samples = []StatSample{{10, 20}, {3, 10}, {5, 2}}

// SetSeed ...
func SetSeed(seed int64) {
	rand.Seed(seed)
}

// CreateStatFromCSV ...
func CreateStatFromCSV(filename string) []float64 {
	fp, err := os.Open(filename)
	ut.Errfunc(err)

	var samples []StatSample

	log.Println("Successfully Opened CSV file")
	defer fp.Close()

	lines, err := csv.NewReader(fp).ReadAll()
	ut.Errfunc(err)

	for _, line := range lines {
		fmt.Println(line)
		count, err := strconv.Atoi(line[0])
		ut.Errfunc(err)

		value, err := strconv.ParseFloat(line[1], 64)
		ut.Errfunc(err)

		ss := StatSample{Count: count, Value: value}
		samples = append(samples, ss)

	}
	return CreateStat(samples)
}

// Pickone ...
func Pickone(in []float64) float64 {
	
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	return pick
}

// Createarray ...
func Createarray(sample StatSample) []float64 {
	myarr := make([]float64, sample.Count)

	for i := range myarr {
		myarr[i] = sample.Value
	}

	return myarr
}

// CreateStat ...
func CreateStat(samples []StatSample) []float64 {

	log.Println("CREATING STAT")
	var array []float64

	for _, v := range samples {
		fmt.Println(v.Count, v.Value)
		arr := Createarray(v)
		array = append(array, arr...)
	}

	return array
}

func maindemo() {

	a := CreateStat(samples)
	fmt.Println(a)

	mean, err := stats.Mean(a)
	ut.Errfunc(err)
	for true {
		time.Sleep(10 * time.Second)

		fmt.Println("cicle", Pickone(a), mean, err)
	}
}
