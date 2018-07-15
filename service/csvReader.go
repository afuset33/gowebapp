package service

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadAll() (rows [][]string) {
	log.Printf("call ReadLine")
	fr, errOpen := os.Open("../config/dictionary.csv")

	if errOpen != nil {
		log.Fatal("ReadLine file open")
	}

	defer fr.Close()

	r := csv.NewReader(fr)

	rows, errRead := r.ReadAll()
	if errRead != nil {
		log.Fatal("ReadLine csv reading")
	}
	return
}
