package jsonReader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type PasswordCheckerConfig struct {
	Length            int `json:"length"`
	NumComboChar      int `json:"numComboChar"`
	NumContinuousChar int `json:"numContinuousChar"`
}

func ReadAll() (conf PasswordCheckerConfig) {
	raw, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(raw, &conf)
	return
}
