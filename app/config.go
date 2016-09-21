package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Config map[string]interface{}

func LoadConfig() {
	log.Println("loading app config...")
	file, error := ioutil.ReadFile("config.json")
	if error != nil {
		log.Fatal("config file is missing...")
	}

	var f interface{}
	error = json.Unmarshal(file, &f)
	if error != nil {
		log.Fatal("config file has invalid JSON data...")
	}

	Config = f.(map[string]interface{})
	log.Println("app config is loaded...")
}
