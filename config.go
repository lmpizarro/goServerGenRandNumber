package statproject

import (
	"log"
)

// DEBUG ...
var DEBUG bool = true

// Config ...
func Config() string {
	return "statproject config"
}

// CsvCountValfile ...
var CsvCountValfile = "countVals.csv"

// GetCSVFilename ...
func GetCSVFilename() string {

	log.Println("SETTING CSV FILE")

	if DEBUG {
		return "../../" + CsvCountValfile
	}
	return "./" + CsvCountValfile

}
