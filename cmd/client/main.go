package main

import (
	"log"
	"fmt"
	// "io/ioutil"
	"net/http"
	"encoding/json"
)

var url string = "http://localhost:8080/api/v1/"

func createquerySetsampler(counts []int, vals []float64) string {

	if len(vals) != len(counts){
		panic(2)
	}

	log.Printf("len counts %d len vals %d", len(counts), len(vals))
	var query string = ""
	for i, v := range counts {
		val := vals[i]
		pair := fmt.Sprintf("count=%d&value=%f&", v, val)
		query = query + pair
	}

	le := len(query) - 1
	query = query[:le]
	return query
}


func makeRequestGet(url string)  map[string]interface{} {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

    return result
}


func makeRequestSampler(){
	url = url + "sampler"
}

func main(){
	counts := []int{10, 20, 10, 4}
	vals := []float64{2, 4.4, 3.4, 5.4}

	query := createquerySetsampler(counts, vals)

	urlq := url + "setsampler" + "?" + query
	
	resp := makeRequestGet(urlq)
	log.Println(resp["message"])

	urlq = url + "sampler"
	
	log.Println(url)
	for i := 0; i < 200; i++ {
	    resp = makeRequestGet(urlq)
		fmt.Println(resp["randomnumber"])
	}

}