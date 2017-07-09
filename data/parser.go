package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//ClientURL ...
type ClientURL struct {
	Name string
	URL  string
}

//ParseAll ...
func ParseAll() []ClientURL {
	raw, err := ioutil.ReadFile("data/client_ads.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []ClientURL
	err = json.Unmarshal(raw, &c)
	if err != nil {
		log.Println("ERR at parse.go: " + err.Error())
	}

	return c
}
