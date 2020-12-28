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
)

// StatSample ...
type StatSample struct {
	Count int
	Value float64
}

type valuePairs struct {
	pdf    []float64
	sample []float64
}

var samples = []StatSample{{10, 20}, {3, 10}, {5, 2}}

// SetSeed ...
func SetSeed(seed int64) {
	rand.Seed(seed)
}

var vals valuePairs

func integrate(samples []StatSample, N int, Nsamples int) {
	log.Println("CALL INTEGRATE FOR PDF")

	vals.pdf = make([]float64, N)
	vals.sample = make([]float64, N)

	vals.sample[0] = samples[0].Value
	vals.pdf[0] = float64(samples[0].Count) / float64(Nsamples)

	for i := 1; i < N; i++ {
		vals.pdf[i] = vals.pdf[i-1] + float64(samples[i].Count)/float64(Nsamples)
		vals.sample[i] = samples[i].Value

	}

	log.Println("PDF", vals.pdf)
	log.Println("Values", vals.sample)

}

// CreateStatFromArrays ...
func CreateStatFromArrays(counts, values []string) {
	log.Println("CREATESTARTFROMARRAY")

	var samples []StatSample

	for i, count := range counts {
		log.Printf("count %s value %s", count, values[i])
		countf, err := strconv.Atoi(count)
		ut.Errfunc(err)
		valf, err := strconv.ParseFloat(values[i], 64)
		ut.Errfunc(err)
		ss := StatSample{Count: countf, Value: valf}
		samples = append(samples, ss)
	}

	CreateStat(samples)
}

// CreateStatFromCSV ...
func CreateStatFromCSV(filename string) {
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

	CreateStat(samples)
}

// Pickone ...
func Pickone() float64 {

	randomIndex := rand.Intn(len(samplerarray))
	pick := samplerarray[randomIndex]
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

var samplerarray []float64

// CreateStat ...
func CreateStat(samples []StatSample) {

	log.Println("CREATING STAT")
	var array []float64

	var Nvals int = 0
	var Nsamples int = 0
	for _, v := range samples {
		Nvals++
		Nsamples = Nsamples + v.Count
		fmt.Println(v.Count, v.Value)
		arr := Createarray(v)
		array = append(array, arr...)
	}

	log.Printf("Nvals: %d, Nsamples %d", Nvals, Nsamples)
	log.Println(samples)
	integrate(samples, Nvals, Nsamples)
	samplerarray = array
}

func maindemo() {

	CreateStat(samples)

	for true {
		time.Sleep(10 * time.Second)

		fmt.Println("cicle", Pickone())
	}
}
