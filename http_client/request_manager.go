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
	Work([]data.ClientURLs)
}

type requestManger struct {
	client  *http.Client
	rAd     requestAd
	urlStrs []string
	respAd  chan *RespAd
}

//RespAd ...
type RespAd struct {
	Bid    string
	AdCode string
}

type requestAd struct {
	gender string
	age    string
}

func (r *requestManger) Work(clientURLs []data.ClientURLs) {
	var rspAd RespAd
	for _, clientURL := range clientURLs {
		log.Println("Making request to " + clientURL.URL)
		req, _ := http.NewRequest("POST", clientURL.URL, nil)
		query := req.URL.Query()
		query.Set("gender", r.rAd.gender)
		query.Set("age", r.rAd.age)
		go func() {
			resp, err := r.client.Do(req)
			if err != nil {
				log.Println("ERR: " + err.Error())
				r.respAd <- nil
				return
			}
			if resp.StatusCode == 200 {
				buf := bytes.NewBuffer(make([]byte, 0))
				buf.ReadFrom(resp.Body)
				body := buf.Bytes()
				if err := json.Unmarshal(body, &rspAd); err != nil {
					log.Println("JSON Unmarshal ERR: " + err.Error())
					r.respAd <- nil
					return
				}
			} else {
				log.Println("respone status code received: " + strconv.Itoa(resp.StatusCode))
				r.respAd <- nil
				return
			}
			r.respAd <- &rspAd
		}()
	}

}

//NewRequestManager ...
func NewRequestManager(gender string, age int, respAd chan *RespAd) RequestManager {
	var rAd = requestAd{
		gender: gender,
		age:    strconv.Itoa(age),
	}
	client := &http.Client{Timeout: twoSecsDuration}

	return &requestManger{rAd: rAd, client: client, respAd: respAd}
}
