package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	sp "github.com/lmpizarro/statproject"
	rg "github.com/lmpizarro/statproject/pkg/randgen"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	p := rg.Pickone()

	log.Printf("SERVING %f", p)

	resp := fmt.Sprintf(`{"randomnumber": %f" }`, p)

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

	api.HandleFunc("/sampler", get).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
