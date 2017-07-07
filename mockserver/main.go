package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"
)

var port string

var profile string

//RespData ...
type RespData struct {
	Bid    float64 `json:"bid"`
	Adcode string  `json:"adcode"`
}

var data []byte

const postMethod = "POST"

/*
	This mock server created to test functionality
*/
func main() {
	initData()
	http.ListenAndServe(port, nil)
}

func initData() {
	flag.StringVar(&profile, "profile", "mocknikead", "provide mock server profile [mocknikead|mockamazonad|mockebayad] ")
	flag.Parse()
	log.Println("Profile selected: " + profile)
	switch profile {
	case "mockamazonad":
		data, _ = json.Marshal(
			RespData{
				Bid:    1.00,
				Adcode: "http://www.amazon.com/in/en_gb/?ref=https%253A%252F%252Fwww.google.co.in%252F",
			},
		)
		port = ":7000"
		http.HandleFunc("/getmead", amazonHandler)

	case "mocknikead":
		data, _ = json.Marshal(
			RespData{
				Bid:    2.01,
				Adcode: "http://www.nike.com/in/en_gb/?ref=https%253A%252F%252Fwww.google.co.in%252F",
			},
		)
		port = ":8082"
		http.HandleFunc("/getad", nikeHandler)

	case "mockebayad":
		fallthrough
	default:
		data, _ = json.Marshal(RespData{
			Bid:    3.01,
			Adcode: "http://www.ebay.com/en_gb/?ref=https%253A%252F%252Fwww.google.co.in%252F",
		})
		port = ":9000"
		http.HandleFunc("/getmead", ebayHandler)
	}
}

func nikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == postMethod {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		log.Println(string(data[:len(data)]))

	} else {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}

}

func amazonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == postMethod {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		log.Println(string(data[:len(data)]))
	} else {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}

}
func ebayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == postMethod {
		//To test timeout
		time.Sleep(3 * time.Second)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		log.Println(string(data[:len(data)]))

	} else {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}
