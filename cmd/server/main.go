package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	sp "github.com/lmpizarro/statproject"
	rg "github.com/lmpizarro/statproject/pkg/randgen"
)

func getsample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	p := rg.Pickone()

	log.Printf("SERVING %f", p)

	resp := fmt.Sprintf(`{"randomnumber": %f" }`, p)

	w.Write([]byte(resp))
}

//
// curl "http://localhost:8080/api/v1/setsampler?count=10&value=5&count=10&value=6"
func setsampler(w http.ResponseWriter, r *http.Request) {
	log.Println("CALL SETSAMPLER WEB")
	query := r.URL.Query()
	counts, presentc := query["count"]
	values, presentv := query["value"]

	if !presentc || !presentv || len(counts) == 0 || len(values) == 0 {
		fmt.Println("counts values not present")
	}

	if len(counts) != len(values) {
		fmt.Println("lengths not equal")
	}

	if len(counts) > 2 {
		fmt.Println("not enough values")
	}



	rg.CreateStatFromArrays(counts, values)	

	w.WriteHeader(200)

	resp := fmt.Sprintf(`hello query`)
	w.Write([]byte(resp))
}

// main func
func main() {
	log.Println("CALL MAIN ", sp.Config())

	filename := sp.GetCSVFilename() // "./countVals.csv"
	rg.CreateStatFromCSV(filename)
	rg.SetSeed(5)

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/sampler", getsample).Methods(http.MethodGet)
	api.HandleFunc("/setsampler", setsampler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
