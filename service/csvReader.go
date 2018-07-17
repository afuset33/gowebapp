package service

import (
	"encoding/csv"
	"log"
	"os"
)

/*
CSVファイル全体を読み込み、スライスにして返却します
*/
func ReadAll() (rows [][]string) {
	log.Printf("call ReadAll")
	fr, errOpen := os.Open("../config/dictionary.csv")

	if errOpen != nil {
		log.Fatal("ReadAll file open")
	}

	defer fr.Close()

	r := csv.NewReader(fr)

	rows, errRead := r.ReadAll()
	if errRead != nil {
		log.Fatal("ReadAll csv reading")
	}
	return
}
