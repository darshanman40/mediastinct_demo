package httpclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/darshanman40/mediastinct_demo/data"
)

const twoSecsDuration = 2 * time.Second

//RequestManager ...
type RequestManager interface {
	Works([]data.ClientURL)
}

type requestManger struct {
	client             HTTPClient
	rAd                requestAd
	urlStrs            []string
	respAd, respWorker chan *RespAd
}

//RespAd ...
type RespAd struct {
	Bid    float64 `json:"bid"`
	AdCode string  `json:"adcode"`
}

type requestAd struct {
	gender string
	age    string
}

func (r *requestManger) Works(clientURLs []data.ClientURL) {
	r.respWorker = make(chan *RespAd)

	for _, clientURL := range clientURLs {
		log.Println("Making request to " + clientURL.URL)
		req, _ := http.NewRequest("POST", clientURL.URL, nil)
		query := req.URL.Query()
		query.Set("gender", r.rAd.gender)
		query.Set("age", r.rAd.age)
		go work(r, req, clientURL)
	}

	ticker := time.NewTicker(twoSecsDuration)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case rw, isOpen := <-r.respWorker:
			if !isOpen {
				log.Println("Channel is closed")
			} else {
				r.respAd <- rw
			}

		case <-ticker.C:
			log.Println("Closing channels")
			close(r.respAd)
			return

		}

	}
}

func work(r *requestManger, req *http.Request, clientURL data.ClientURL) {
	var rspAd RespAd
	resp, err := r.client.Do(req)
	if err != nil {
		log.Println("ERR: " + err.Error())
		r.respWorker <- nil
		return
	}

	if resp.StatusCode == 200 {
		buf := bytes.NewBuffer(make([]byte, 0))
		buf.ReadFrom(resp.Body)
		body := buf.Bytes()
		if err := json.Unmarshal(body, &rspAd); err != nil {
			log.Println("JSON Unmarshal ERR: " + err.Error())
			r.respWorker <- nil
			return
		}
	} else {
		log.Println("respone status code received: " + strconv.Itoa(resp.StatusCode))
		r.respWorker <- nil
		return
	}
	log.Println("Recieved from client: " + rspAd.AdCode)
	r.respWorker <- &rspAd
}

//NewRequestManager ...
func NewRequestManager(gender string, age int, respAd chan *RespAd) RequestManager {
	var rAd = requestAd{
		gender: gender,
		age:    strconv.Itoa(age),
	}
	client := NewHTTPClient()
	return &requestManger{rAd: rAd, client: client, respAd: respAd}
}
