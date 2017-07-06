package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const port = ":8082"

//RespData ...
type RespData struct {
	Bid    string `json:"bid"`
	Adcode string `json:"adcode"`
}

var nikeData []byte

var amazonData []byte

var ebayData []byte

const postMethod = "POST"

/*
	This mock server created to test functionality
*/
func main() {
	initData()

	http.HandleFunc("/nike/getad", nikeHandler)
	http.HandleFunc("/amazon/getmead", amazonHandler)
	http.HandleFunc("/ebay/getmead", ebayHandler)
	http.ListenAndServe(port, nil)
}

func initData() {
	amazonData, _ = json.Marshal(
		RespData{
			Bid:    "1.00",
			Adcode: "http://www.amazon.com/in/en_gb/?ref=https%253A%252F%252Fwww.google.co.in%252F",
		},
	)
	nikeData, _ = json.Marshal(
		RespData{
			Bid:    "2.00",
			Adcode: "http://www.nike.com/in/en_gb/?ref=https%253A%252F%252Fwww.google.co.in%252F",
		},
	)
	ebayData, _ = json.Marshal(RespData{})
}

func nikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == postMethod {
		w.Header().Set("Content-Type", "application/json")
		w.Write(nikeData)
		log.Println(string(nikeData[:len(nikeData)]))

	} else {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}

}

func amazonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == postMethod {
		w.Header().Set("Content-Type", "application/json")
		w.Write(amazonData)
		log.Println(string(amazonData[:len(amazonData)]))
	} else {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}

}
func ebayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == postMethod {
		time.Sleep(3 * time.Second)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		w.Write(ebayData)
		log.Println(string(ebayData[:len(ebayData)]))

	} else {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}
