package utils

import "log"

// Errfunc ...
func Errfunc(err error){
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}