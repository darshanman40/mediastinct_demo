package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/darshanman40/mediastinct_demo/data"
	"github.com/darshanman40/mediastinct_demo/http_client"
)

const (
	genderTag = "gender"
	ageTag    = "age"

	postMethodNotAllowed   = "Post method not allowed"
	getMethodNotAllowed    = "Get method not allowed"
	putMethodNotAllowed    = "Put method not allowed"
	deleteMethodNotAllowed = "Delete method not allowed"
)

var mutex *sync.Mutex

type page struct {
	title string
	body  []byte
}

//RespData ...
type RespData struct {
	Adcode string `json:"adcode"`
}

type getMeAdService struct {
	bids       chan *httpclient.RespAd
	clientURLs []data.ClientURL
}

func (g *getMeAdService) DoGet(w http.ResponseWriter, r *http.Request) {
	var gender = r.FormValue(genderTag)
	var age, err = strconv.Atoi(r.FormValue(ageTag))
	if err != nil {
		log.Println("Failed to conver age datatype: " + err.Error())
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	rm := httpclient.NewRequestManager(gender, age, g.bids)
	adCode := process(g.clientURLs, rm, g.bids)

	if adCode == "" {
		log.Println("No adCode found")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	} else {
		rd := RespData{Adcode: adCode}
		if data, err := json.Marshal(rd); err != nil {
			log.Println("Marshal failed for adCode")
			http.Error(w, "500: Internal Server Error", http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	}
}

func (g *getMeAdService) DoPost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, postMethodNotAllowed, http.StatusMethodNotAllowed)
}

func (g *getMeAdService) DoPut(w http.ResponseWriter, r *http.Request) {
	http.Error(w, putMethodNotAllowed, http.StatusMethodNotAllowed)
}

func (g *getMeAdService) DoDelete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, deleteMethodNotAllowed, http.StatusMethodNotAllowed)
}

func process(clientURLs []data.ClientURL, rm httpclient.RequestManager, rspAd chan *httpclient.RespAd) string {
	maxBid := 0.00
	var adCode string

	rm.Works(clientURLs)

	for bid := range rspAd {
		if bid != nil {
			newBid := strconv.FormatFloat(bid.Bid, 'f', 6, 64)
			log.Println("Recieved: AdCode:" + bid.AdCode + " Bid:" + newBid)
			if maxBid < bid.Bid {
				log.Println("Pushing: AdCode:" + bid.AdCode + " Bid:" + newBid)
				maxBid = bid.Bid
				adCode = bid.AdCode
			}
		} else {
			log.Println("Received nil bid")
		}
	}

	return adCode
}

//NewGetMeAdService ...
func NewGetMeAdService(mock bool) Service {
	mutex = &sync.Mutex{}
	var clientURLs []data.ClientURL
	if mock {
		clientURLs = []data.ClientURL{
			data.ClientURL{
				Name: "Nike",
				URL:  "http://localhost:8082/getad",
			},
			data.ClientURL{
				Name: "amazon",
				URL:  "http://localhost:7000/getmead",
			},
			data.ClientURL{
				Name: "ebay",
				URL:  "http://localhost:9000/getmead",
			},
		}
	} else {
		clientURLs = data.ParseAll()
	}

	return &getMeAdService{bids: make(chan *httpclient.RespAd, len(clientURLs)),
		clientURLs: clientURLs,
	}
}
